package scan

import (
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.uber.org/zap"
)

type scan struct {
	app         *configx.Application
	blockClient model.BlockServiceClient
}

func NewService(app *configx.Application, blockClient model.BlockServiceClient) adapterx.Service {
	return &scan{
		app:         app,
		blockClient: blockClient,
	}
}

func (i *scan) Start(ctx contextx.Contextx) error {
	// TODO: 2024/7/28|sean|add block scan logic here
	// i.injector.BlockService.ScanBlock(&model.ScanBlockRequest{}, stream)
	block, err := i.blockClient.GetBlock(ctx, &model.GetBlockRequest{
		Workchain: -1,
		Shard:     8000000000000000,
		SeqNo:     39346131,
	})
	if err != nil {
		return err
	}
	ctx.Debug("get block", zap.Any("block", &block))

	return nil
}

func (i *scan) AwaitSignal(ctx contextx.Contextx) error {
	ctx.Info("await signal")
	// TODO: 2024/7/29|sean|add block scan await signal logic here
	return nil
}
