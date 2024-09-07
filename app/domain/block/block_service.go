package block

import (
	"context"
	"time"

	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/entity/domain/block/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type impl struct {
	client *tonx.Client
	blocks repo.IBlockRepo
}

// NewBlockService is used to create a new model.BlockServiceServer
func NewBlockService(client *tonx.Client, blocks repo.IBlockRepo) biz.BlockServiceServer {
	return &impl{
		client: client,
		blocks: blocks,
	}
}

func (i *impl) GetBlock(c context.Context, req *biz.GetBlockRequest) (*model.Block, error) {
	ctx := contextx.WithContext(c)

	api := ton.NewAPIClient(i.client).WithRetry()
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

	ret, err := model.NewBlock(blockID.Workchain, blockID.Shard, blockID.SeqNo)
	if err != nil {
		ctx.Error("failed to create block", zap.Error(err))
		return nil, err
	}
	ret.Timestamp = timestamppb.New(time.Unix(int64(blockData.BlockInfo.GenUtime), 0))

	return ret, nil
}

func (i *impl) GetBlocks(req *biz.GetBlocksRequest, stream biz.BlockService_GetBlocksServer) error {
	c := stream.Context()

	next, span := otelx.Tracer.Start(c, "block.biz.GetBlocks")
	defer span.End()

	ctx := contextx.WithContext(c)

	items, _, err := i.blocks.List(next, repo.ListCondition{
		Limit: 0,
		Skip:  0,
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

	return nil
}

func (i *impl) ScanBlock(req *biz.ScanBlockRequest, stream biz.BlockService_ScanBlockServer) error {
	api := ton.NewAPIClient(i.client, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(i.client.Config)

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
		ctx.Info("block sent", zap.String("block_id", newBlock.Id))

		next := master.SeqNo + 1
		master, err2 = api.WaitForBlock(next).LookupBlock(stickyContext, master.Workchain, master.Shard, next)
		if err2 != nil {
			ctx.Error("failed to lookup block", zap.Uint32("seq_no", next), zap.Error(err2))
			return err2
		}
	}
}
