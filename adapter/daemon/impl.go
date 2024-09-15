package daemon

import (
	"context"

	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/app/usecase/event"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/eventx"
	"go.uber.org/zap"
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

	if i.grpcserver != nil {
		if err := i.grpcserver.Start(ctx); err != nil {
			ctx.Error("start grpc server", zap.Error(err))
			return err
		}
	}

	err := i.bus.Subscribe(event.NewFoundBlockHandler())
	if err != nil {
		ctx.Error("subscribe found block handler", zap.Error(err))
		return err
	}
	ctx.Info("subscribed to block events")

	return nil
}

func (i *impl) Shutdown(c context.Context) error {
	ctx := contextx.WithContext(c)
	ctx.Info("server shutdown")

	if i.grpcserver != nil {
		if err := i.grpcserver.Stop(ctx); err != nil {
			ctx.Error("shutdown grpc server", zap.Error(err))
		}
	}

	// TODO: 2024/9/15|sean|unsubscribe found block handler

	return nil
}
