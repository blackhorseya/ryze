package block

import (
	"context"
	"strconv"
	"time"

	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/entity/domain/block/repo"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/eventx"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type impl struct {
	tonClient *tonx.Client
	bus       *eventx.EventBus

	txClient txB.TransactionServiceClient

	blocks repo.IBlockRepo
}

// NewBlockService is used to create a new block service
func NewBlockService(
	tonClient *tonx.Client,
	blocks repo.IBlockRepo,
	bus *eventx.EventBus,
) biz.BlockServiceServer {
	return &impl{
		tonClient: tonClient,
		bus:       bus,
		blocks:    blocks,
	}
}

//nolint:gocognit // ignore cognitive complexity
func (i *impl) ScanBlock(req *biz.ScanBlockRequest, stream biz.BlockService_ScanBlockServer) error {
	c := stream.Context()
	next, span := otelx.Tracer.Start(c, "block.biz.ScanBlock")
	defer span.End()

	// 初始化 TON API 客戶端
	api := ton.NewAPIClient(i.tonClient, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(i.tonClient.Config)

	ctx := contextx.WithContext(c)

	// 獲取主鏈資訊
	master, err := api.GetMasterchainInfo(next)
	if err != nil {
		ctx.Error("failed to get master-chain info", zap.Error(err))
		return err
	}
	ctx.Info("master proofs chain successfully verified, all data is now safe and trusted!")

	// 綁定單一伺服器的上下文以保持一致性
	stickyContext := api.Client().StickyContext(next)

	// 儲存分片的最後序列號，防止重複處理
	shardLastSeqno := map[string]uint32{}

	// 從主鏈獲取所有的分片資訊
	firstShards, err := api.GetBlockShardsInfo(stickyContext, master)
	if err != nil {
		ctx.Error("failed to get block shards info", zap.Error(err))
		return err
	}

	// 初始化分片序號的記錄
	for _, shard := range firstShards {
		shardLastSeqno[tonx.GetShardID(shard)] = shard.SeqNo
	}

	// 持續監聽所有分片上的新區塊
	for {
		// 獲取每個 workchain 和 shard 上的新區塊
		currentShards, err2 := api.GetBlockShardsInfo(stickyContext, master)
		if err2 != nil {
			ctx.Error("failed to get block shards info", zap.Error(err2))
			return err2
		}

		for _, shard := range currentShards {
			// 只監聽指定的 workchain
			if req.Workchain != nil && shard.Workchain != *req.Workchain {
				continue
			}

			// 檢查是否有新的區塊
			if lastSeqno, ok := shardLastSeqno[tonx.GetShardID(shard)]; ok && shard.SeqNo <= lastSeqno {
				continue
			}

			// 更新分片序號
			shardLastSeqno[tonx.GetShardID(shard)] = shard.SeqNo

			// 創建一個新的區塊事件並發送
			newBlock, err3 := model.NewBlock(shard.Workchain, shard.Shard, shard.SeqNo)
			if err3 != nil {
				ctx.Error("failed to create block", zap.Error(err3))
				return err3
			}

			err3 = stream.Send(newBlock)
			if err3 != nil {
				ctx.Error("failed to send block event", zap.Error(err3))
				return err3
			}

			ctx.Info("block event sent", zap.Any("block", &newBlock))
		}

		// 更新主鏈區塊以繼續監控新地分片區塊
		nextSeqNo := master.SeqNo + 1
		master, err2 = api.WaitForBlock(nextSeqNo).LookupBlock(stickyContext, master.Workchain, master.Shard, nextSeqNo)
		if err2 != nil {
			ctx.Error("failed to lookup next block", zap.Uint32("seq_no", nextSeqNo), zap.Error(err2))
			return err2
		}
	}
}

func (i *impl) FoundNewBlock(c context.Context, req *biz.FoundNewBlockRequest) (*model.Block, error) {
	next, span := otelx.Tracer.Start(c, "block.biz.FoundNewBlock")
	defer span.End()

	ctx := contextx.WithContext(c)

	api := ton.NewAPIClient(i.tonClient).WithRetry()
	blockID, err := api.LookupBlock(ctx, req.Workchain, req.Shard, req.SeqNo)
	if err != nil {
		ctx.Error("failed to lookup block", zap.Error(err), zap.Any("req", &req))
		return nil, err
	}

	blockData, err := api.GetBlockData(ctx, blockID)
	if err != nil {
		ctx.Error("failed to get block data", zap.Error(err))
		return nil, err
	}

	block, err := model.NewBlock(blockID.Workchain, blockID.Shard, blockID.SeqNo)
	if err != nil {
		ctx.Error("failed to create block", zap.Error(err))
		return nil, err
	}
	block.Timestamp = timestamppb.New(time.Unix(int64(blockData.BlockInfo.GenUtime), 0))

	event := block.Born()

	err = i.blocks.Create(next, block)
	if err != nil {
		ctx.Error("failed to create block", zap.Error(err))
		return nil, err
	}

	i.bus.Publish(event)

	return block, nil
}

func (i *impl) GetBlock(c context.Context, req *biz.GetBlockRequest) (*model.Block, error) {
	next, span := otelx.Tracer.Start(c, "block.biz.GetBlock")
	defer span.End()

	ctx := contextx.WithContext(c)

	block, err := i.blocks.GetByID(next, req.BlockId)
	if err != nil {
		ctx.Error("failed to get block", zap.Error(err))
		return nil, err
	}

	return block, nil
}

func (i *impl) ListBlocks(req *biz.ListBlocksRequest, stream grpc.ServerStreamingServer[model.Block]) error {
	c := stream.Context()
	next, span := otelx.Tracer.Start(c, "block.biz.ListBlocks")
	defer span.End()

	ctx := contextx.WithContext(c)

	items, total, err := i.blocks.List(next, repo.ListCondition{
		Limit: req.PageSize,
		Skip:  (req.Page - 1) * req.PageSize,
	})
	if err != nil {
		ctx.Error("failed to list blocks", zap.Error(err))
		return err
	}

	for _, item := range items {
		err = stream.Send(item)
		if err != nil {
			ctx.Error("failed to send block", zap.Error(err))
			return err
		}
	}
	stream.SetTrailer(metadata.New(map[string]string{"total": strconv.Itoa(total)}))

	ctx.Debug("list blocks", zap.Any("items", items), zap.Int("total", total))

	return nil
}
