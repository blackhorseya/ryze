package daemon

import (
	"context"

	"github.com/blackhorseya/ryze/internal/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/eventx"
	"go.uber.org/zap"
)

type impl struct {
	injector *Injector
	bus      eventx.EventBus
}

// NewServer is a function to create a new server.
func NewServer(injector *Injector, bus eventx.EventBus) (adapterx.Server, func(), error) {
	return &impl{
		injector: injector,
		bus:      bus,
	}, func() {}, nil
}

func (i *impl) Start(c context.Context) error {
	ctx := contextx.WithContext(c)
	ctx.Info("server start")

	// start scanning for blocks using local service
	blocks := make(chan *model.Block)
	go func() {
		_ = i.injector.BlockSvc.ScanBlock(ctx, blocks)
		close(blocks)
	}()

	// listen for new block events and publish them via the EventBus
	go i.listenForBlockEvents(ctx, blocks)

	// subscribe found block handler
	// err := i.bus.Subscribe(event.NewFoundBlockHandlerV2(i.injector.blockClient, i.injector.txClient))
	// if err != nil {
	// 	ctx.Error("subscribe found block handler", zap.Error(err))
	// 	return err
	// }
	// ctx.Info("subscribed to block events")

	return nil
}

func (i *impl) Shutdown(c context.Context) error {
	ctx := contextx.WithContext(c)
	ctx.Info("server shutdown")

	// TODO: 2024/9/15|sean|unsubscribe found block handler

	return nil
}

func (i *impl) listenForBlockEvents(ctx contextx.Contextx, blocks <-chan *model.Block) {
	ctx.Info("start to receive block")

	for {
		select {
		case <-ctx.Done():
			ctx.Info("context done, stopping block event listener")
			return
		case newBlock, ok := <-blocks:
			if !ok {
				ctx.Info("block channel closed")
				return
			}
			ctx.Info("received block", zap.String("block_id", newBlock.ID))
			// _ = i.bus.Publish(newBlock.Born())
		}
	}
}
