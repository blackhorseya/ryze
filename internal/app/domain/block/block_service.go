package block

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/internal/app/repo"
	"github.com/blackhorseya/ryze/internal/shared/tonx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type impl struct {
	tonClient *tonx.Client

	blocks repo.IBlockRepo
}

// NewBlockService is used to create a new block service
func NewBlockService(
	tonClient *tonx.Client,
	blocks repo.IBlockRepo,
) biz.BlockServiceServer {
	return &impl{
		tonClient: tonClient,
		blocks:    blocks,
	}
}

//nolint:gocognit // ignore cognitive complexity
func (i *impl) ScanBlock(req *biz.ScanBlockRequest, stream biz.BlockService_ScanBlockServer) error {
	ctx, span := contextx.StartSpan(stream.Context(), "block.biz.ScanBlock")
	defer span.End()

	// 初始化 TON API 客戶端
	api := ton.NewAPIClient(i.tonClient, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(i.tonClient.Config)

	// 獲取主鏈資訊
	master, err := api.GetMasterchainInfo(ctx)
	if err != nil {
		ctx.Error("failed to get master-chain info", zap.Error(err))
		return err
	}
	ctx.Info("master proofs chain successfully verified, all data is now safe and trusted!")

	// 綁定單一伺服器的上下文以保持一致性
	stickyContext := api.Client().StickyContext(ctx)

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
		currentShards, err2 := api.GetBlockShardsInfo(ctx, master)
		if errors.Is(err2, context.Canceled) {
			ctx.Info("scan block canceled")
			return nil
		}
		if err2 != nil {
			ctx.Error("failed to get block shards info", zap.Error(err2))
			return err2
		}

		for _, shard := range currentShards {
			// 只監聽指定的 workchain
			if req.Workchain != nil && shard.Workchain != req.GetWorkchain() {
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
		master, err2 = api.WaitForBlock(nextSeqNo).LookupBlock(ctx, master.Workchain, master.Shard, nextSeqNo)
		if errors.Is(err2, context.Canceled) {
			ctx.Info("scan block canceled")
			return nil
		}
		if err2 != nil {
			ctx.Error("failed to lookup next block", zap.Uint32("seq_no", nextSeqNo), zap.Error(err2))
			return err2
		}
	}
}

func (i *impl) FoundNewBlock(stream grpc.BidiStreamingServer[model.Block, model.Block]) error {
	ctx, span := contextx.StartSpan(stream.Context(), "block.biz.FoundNewBlock")
	defer span.End()

	for {
		newBlock, err := stream.Recv()
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.Canceled {
			ctx.Info("found new block canceled")
			return nil
		}
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			ctx.Error("failed to receive new block", zap.Error(err))
			continue
		}

		err = i.fetchBlockInfo(ctx, newBlock)
		if errors.Is(err, context.Canceled) {
			ctx.Info("found new block canceled")
			return nil
		}
		if err != nil {
			ctx.Error("failed to fetch block info", zap.Error(err))
			continue
		}

		err = i.blocks.Create(ctx, newBlock)
		if err != nil {
			ctx.Error("failed to create block", zap.Error(err))
			continue
		}

		err = stream.Send(newBlock)
		if err != nil {
			ctx.Error("failed to send new block", zap.Error(err))
			continue
		}
	}
}

func (i *impl) GetBlock(c context.Context, req *biz.GetBlockRequest) (*model.Block, error) {
	ctx, span := contextx.StartSpan(c, "block.biz.GetBlock")
	defer span.End()

	block, err := i.blocks.GetByID(ctx, req.BlockId)
	if err != nil {
		ctx.Error("failed to get block", zap.Error(err))
		return nil, err
	}

	return block, nil
}

func (i *impl) ListBlocks(req *biz.ListBlocksRequest, stream grpc.ServerStreamingServer[model.Block]) error {
	ctx, span := contextx.StartSpan(stream.Context(), "block.biz.ListBlocks")
	defer span.End()

	items, total, err := i.blocks.List(ctx, repo.ListCondition{
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

func (i *impl) FoundNewBlockNonStream(c context.Context, req *biz.FoundNewBlockRequest) (*model.Block, error) {
	ctx, span := contextx.StartSpan(c, "block.biz.FoundNewBlockNonStream")
	defer span.End()

	block, err := model.NewBlock(req.Workchain, req.Shard, req.SeqNo)
	if err != nil {
		ctx.Error("failed to create block", zap.Error(err), zap.Any("request", &req))
		return nil, err
	}

	err = i.fetchBlockInfo(ctx, block)
	if errors.Is(err, context.Canceled) {
		ctx.Info("found new block canceled")
		return block, nil
	}
	if err != nil {
		ctx.Error("failed to fetch block info", zap.Error(err), zap.Any("block", &block))
		return nil, fmt.Errorf("failed to fetch block info: %w", err)
	}

	err = i.blocks.Create(ctx, block)
	if err != nil {
		ctx.Error("failed to create block", zap.Error(err), zap.Any("block", &block))
		return nil, fmt.Errorf("failed to create block: %w", err)
	}

	return block, nil
}

// fetchBlockInfo is used to fetch block info
func (i *impl) fetchBlockInfo(c context.Context, block *model.Block) (err error) {
	ctx, span := contextx.StartSpan(c, "block.biz.fetchBlockInfo")
	defer span.End()

	// 初始化 TON API 客戶端
	api := ton.NewAPIClient(i.tonClient).WithRetry()

	// 查找區塊
	blockID, err := api.LookupBlock(ctx, block.Workchain, block.Shard, block.SeqNo)
	if errors.Is(err, context.Canceled) {
		ctx.Info("fetch block info canceled")
		return err
	}
	if err != nil {
		ctx.Error("failed to lookup block", zap.Error(err), zap.Any("block", block))
		return err
	}

	// 獲取區塊資訊
	blockData, err := api.GetBlockData(ctx, blockID)
	if errors.Is(err, context.Canceled) {
		ctx.Info("get block data canceled")
		return err
	}
	if err != nil {
		ctx.Error("failed to get block data", zap.Error(err))
		return err
	}

	// fill block info
	block.Timestamp = timestamppb.New(time.Unix(int64(blockData.BlockInfo.GenUtime), 0))

	return nil
}
