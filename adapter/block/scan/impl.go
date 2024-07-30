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
	stream, err := i.blockClient.ScanBlock(ctx, &model.ScanBlockRequest{})
	if err != nil {
		return err
	}
	go func() {
		for {
			block, err2 := stream.Recv()
			if err2 != nil {
				ctx.Error("receive block error", zap.Error(err2))
				return
			}

			ctx.Info("receive block", zap.Any("block", block))
		}
	}()

	return nil
}

func (i *scan) AwaitSignal(ctx contextx.Contextx) error {
	ctx.Info("await signal")
	return nil
}
