package grpc

import (
	"fmt"

	"github.com/blackhorseya/ryze/adapter/platform/wirex"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
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

func (i *impl) Start(ctx contextx.Contextx) error {
	err := i.server.Start(ctx)
	if err != nil {
		ctx.Error(
			"Failed to start grpc server",
			zap.Error(err),
			zap.String("addr", i.injector.A.GRPC.GetAddr()),
		)
		return err
	}

	ctx.Info("start grpc server")

	return nil
}

func (i *impl) AwaitSignal(ctx contextx.Contextx) error {
	ctx.Info("receive signal to stop server")

	if err := i.server.Stop(ctx); err != nil {
		ctx.Error("Failed to stop server", zap.Error(err))
		return fmt.Errorf("failed to stop server: %w", err)
	}

	return nil
}

// NewInitServersFn creates a new impl server init function.
func NewInitServersFn() grpcx.InitServers {
	return func(s *grpc.Server) {
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_SERVING)

		reflection.Register(s)
	}
}
