package daemon

import (
	"context"
	"errors"
	"io"

	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/eventx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type impl struct {
	injector   *Injector
	grpcserver *grpcx.Server
	bus        eventx.EventBus
}

// NewServer is a function to create a new server.
func NewServer(injector *Injector, grpcserver *grpcx.Server, bus eventx.EventBus) (adapterx.Server, func(), error) {
	return &impl{
		injector:   injector,
		grpcserver: grpcserver,
		bus:        bus,
	}, func() {}, nil
}

func (i *impl) Start(c context.Context) error {
	ctx := contextx.WithContext(c)
	ctx.Info("server start")

	// start service server
	if i.grpcserver != nil {
		if err := i.grpcserver.Start(ctx); err != nil {
			ctx.Error("start grpc server", zap.Error(err))
			return err
		}
	}

	// start scanning for blocks
	blockScanner, err := i.injector.blockClient.ScanBlock(ctx, &blockB.ScanBlockRequest{})
	if err != nil {
		ctx.Error("failed to scan block", zap.Error(err))
		return err
	}

	// listen for new block events and publish them via the EventBus
	go i.listenForBlockEvents(ctx, blockScanner)

	// subscribe found block handler
	// err = i.bus.Subscribe(event.NewFoundBlockHandler())
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

	if i.grpcserver != nil {
		if err := i.grpcserver.Stop(ctx); err != nil {
			ctx.Error("stop grpc server", zap.Error(err))
		}
	}

	// TODO: 2024/9/15|sean|unsubscribe found block handler

	return nil
}

func (i *impl) listenForBlockEvents(ctx contextx.Contextx, stream grpc.ServerStreamingClient[model.Block]) {
	ctx.Info("start to receive block")

	for {
		newBlock, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			ctx.Error("failed to receive block", zap.Error(err))
			continue
		}

		ctx.Info("received block", zap.String("block_id", newBlock.Id))

		_ = i.bus.Publish(newBlock.Born())
	}
}
