package biz

import (
	"context"

	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type impl struct {
	client *tonx.Client
}

// NewBlockService is used to create a new model.BlockServiceServer
func NewBlockService(client *tonx.Client) model.BlockServiceServer {
	return &impl{
		client: client,
	}
}

func (i *impl) GetBlock(ctx context.Context, request *model.GetBlockRequest) (*model.Block, error) {
	// TODO: 2024/7/27|sean|implement me
	return &model.Block{
		Id:             request.Id,
		Height:         0,
		Timestamp:      timestamppb.Now(),
		TransactionIds: nil,
	}, nil
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
		err = stream.Send(&model.Block{
			Id:             nil,
			Height:         master.SeqNo,
			Timestamp:      nil,
			TransactionIds: nil,
		})
		if err != nil {
			ctx.Error("failed to send block", zap.Uint32("seq_no", master.SeqNo), zap.Error(err))
			return err
		}
		ctx.Info("block sent", zap.Uint32("seq_no", master.SeqNo))

		next := master.SeqNo + 1
		master, err = api.WaitForBlock(next).LookupBlock(ctx, master.Workchain, master.Shard, next)
		if err != nil {
			ctx.Error("failed to lookup block", zap.Uint32("seq_no", next), zap.Error(err))
			return err
		}
	}
}