package daemon

import (
	"context"

	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.uber.org/zap"
)

type impl struct {
	injector   *Injector
	grpcServer *grpcx.Server
}

// NewServer is a function to create a new server.
func NewServer(injector *Injector, grpcServer *grpcx.Server) (adapterx.Server, func(), error) {
	return &impl{
		injector:   injector,
		grpcServer: grpcServer,
	}, func() {}, nil
}

func (i *impl) Start(c context.Context) error {
	ctx := contextx.WithContext(c)
	ctx.Info("server start")

	if i.grpcServer != nil {
		if err := i.grpcServer.Start(ctx); err != nil {
			ctx.Error("grpc server start", zap.Error(err))
			return err
		}
	}

	return nil
}

func (i *impl) Shutdown(c context.Context) error {
	ctx := contextx.WithContext(c)
	ctx.Info("server shutdown")

	if i.grpcServer != nil {
		if err := i.grpcServer.Stop(ctx); err != nil {
			ctx.Error("grpc server shutdown", zap.Error(err))
		}
	}

	return nil
}
