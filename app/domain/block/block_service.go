package block

import (
	"context"
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

func (i *impl) ScanBlock(req *biz.ScanBlockRequest, stream biz.BlockService_ScanBlockServer) error {
	api := ton.NewAPIClient(i.tonClient, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(i.tonClient.Config)

	ctx := contextx.WithContext(stream.Context())
	master, err := api.GetMasterchainInfo(ctx)
	if err != nil {
		ctx.Error("failed to get masterchain info", zap.Error(err))
		return err
	}
	ctx.Info("master proofs chain successfully verified, all data is now safe and trusted!")

	stickyContext := api.Client().StickyContext(ctx)
	shardLastSeqno := map[string]uint32{}
	firstShards, err := api.GetBlockShardsInfo(stickyContext, master)
	if err != nil {
		ctx.Error("failed to get block shards info", zap.Error(err))
		return err
	}

	for _, shard := range firstShards {
		shardLastSeqno[tonx.GetShardID(shard)] = shard.SeqNo
	}

	for {
		newBlock, err2 := model.NewBlock(master.Workchain, master.Shard, master.SeqNo)
		if err2 != nil {
			ctx.Error("failed to create block", zap.Error(err2))
			return err2
		}

		err2 = stream.Send(newBlock)
		if err2 != nil {
			ctx.Error("failed to send block", zap.Uint32("seq_no", master.SeqNo), zap.Error(err2))
			return err2
		}
		ctx.Info("block sent", zap.Any("block", &newBlock))

		next := master.SeqNo + 1
		master, err2 = api.WaitForBlock(next).LookupBlock(stickyContext, master.Workchain, master.Shard, next)
		if err2 != nil {
			ctx.Error("failed to lookup block", zap.Uint32("seq_no", next), zap.Error(err2))
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
	ctx.Debug("get block data from ton", zap.Any("block_data", &blockData))

	block, err := model.NewBlock(blockID.Workchain, blockID.Shard, blockID.SeqNo)
	if err != nil {
		ctx.Error("failed to create block", zap.Error(err))
		return nil, err
	}
	block.Timestamp = timestamppb.New(time.Unix(int64(blockData.BlockInfo.GenUtime), 0))

	ctx.Debug("get block", zap.Any("block", &block))
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
	// TODO: 2024/9/13|sean|implement me
	panic("implement me")
}

func (i *impl) ListBlocks(req *biz.ListBlocksRequest, stream grpc.ServerStreamingServer[model.Block]) error {
	// TODO: 2024/9/13|sean|implement me
	panic("implement me")
}
