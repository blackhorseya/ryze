package biz

import (
	"context"

	"github.com/blackhorseya/ryze/entity/domain/block/model"
)

type impl struct {
}

// NewBlockService is used to create a new model.BlockServiceServer
func NewBlockService() model.BlockServiceServer {
	return &impl{}
}

func (i *impl) GetBlock(ctx context.Context, request *model.GetBlockRequest) (*model.Block, error) {
	// TODO: 2024/7/27|sean|implement me
	panic("implement me")
}

func (i *impl) GetBlocks(request *model.GetBlocksRequest, server model.BlockService_GetBlocksServer) error {
	// TODO: 2024/7/27|sean|implement me
	panic("implement me")
}
