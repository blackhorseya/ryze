package grpc

import (
	"fmt"

	"github.com/blackhorseya/ryze/adapter/block/wirex"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type impl struct {
	injector *wirex.Injector
	server   *grpcx.Server
}

// NewGRPC creates a new impl service.
func NewGRPC(injector *wirex.Injector, server *grpcx.Server) adapterx.Service {
	return &impl{
		injector: injector,
		server:   server,
	}
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	err := i.server.Start(ctx)
	if err != nil {
		ctx.Error("Failed to start grpc server", zap.Error(err))
		return err
	}

	ctx.Info("start grpc server")

	return nil
}

func (i *impl) AwaitSignal() error {
	ctx := contextx.Background()
	ctx.Info("receive signal to stop server")

	if err := i.server.Stop(ctx); err != nil {
		ctx.Error("Failed to stop server", zap.Error(err))
		return fmt.Errorf("failed to stop server: %w", err)
	}

	return nil
}

// NewInitServersFn creates a new impl server init function.
func NewInitServersFn(injector *wirex.Injector) grpcx.InitServers {
	return func(s *grpc.Server) {
		model.RegisterBlockServiceServer(s, injector.BlockService)
	}
}
