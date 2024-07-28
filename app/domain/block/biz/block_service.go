package biz

import (
	"context"

	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
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

func (i *impl) ScanBlock(request *model.ScanBlockRequest, server model.BlockService_ScanBlockServer) error {
	// TODO: 2024/7/28|sean|implement me
	panic("implement me")
}
