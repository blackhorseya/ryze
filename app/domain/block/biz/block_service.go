package biz

import (
	"context"
	"time"

	"github.com/blackhorseya/ryze/app/infra/tonx"
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
func NewBlockService(client *tonx.Client, blocks repo.IBlockRepo) model.BlockServiceServer {
	return &impl{
		client: client,
		blocks: blocks,
	}
}

func (i *impl) GetBlock(c context.Context, request *model.GetBlockRequest) (*model.Block, error) {
	ctx := contextx.WithContext(c)

	api := ton.NewAPIClient(i.client)
	blockID, err := api.LookupBlock(ctx, request.Workchain, request.Shard, request.SeqNo)
	if err != nil {
		ctx.Error("failed to lookup block", zap.Error(err))
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

func (i *impl) GetBlocks(request *model.GetBlocksRequest, server model.BlockService_GetBlocksServer) error {
	// TODO: 2024/7/27|sean|implement me
	panic("implement me")
}

func (i *impl) ScanBlock(request *model.ScanBlockRequest, stream model.BlockService_ScanBlockServer) error {
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
		ctx.Info("block sent", zap.Uint32("seq_no", master.SeqNo))

		next := master.SeqNo + 1
		master, err2 = api.WaitForBlock(next).LookupBlock(ctx, master.Workchain, master.Shard, next)
		if err2 != nil {
			ctx.Error("failed to lookup block", zap.Uint32("seq_no", next), zap.Error(err2))
			return err2
		}
	}
}

func (i *impl) FetchAndStoreBlock(
	c context.Context,
	request *model.FetchAndStoreBlockRequest,
) (*model.FetchAndStoreBlockResponse, error) {
	ctx := contextx.WithContext(c)

	block, err := i.GetBlock(ctx, &model.GetBlockRequest{
		Workchain: request.Workchain,
		Shard:     request.Shard,
		SeqNo:     request.SeqNo,
	})
	if err != nil {
		return nil, err
	}

	err = i.blocks.Create(ctx, block)
	if err != nil {
		ctx.Error("failed to create block", zap.Error(err))
		return nil, err
	}

	return &model.FetchAndStoreBlockResponse{
		Block: block,
	}, nil
}
