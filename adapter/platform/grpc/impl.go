package grpc

import (
	"context"
	"fmt"

	"github.com/blackhorseya/ryze/adapter/platform/wirex"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	accountB "github.com/blackhorseya/ryze/entity/domain/account/biz"
	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	netB "github.com/blackhorseya/ryze/entity/domain/network/biz"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
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
func NewGRPC(injector *wirex.Injector, server *grpcx.Server) adapterx.Server {
	return &impl{
		injector: injector,
		server:   server,
	}
}

func (i *impl) Start(c context.Context) error {
	ctx := contextx.WithContext(c)
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

func (i *impl) Shutdown(c context.Context) error {
	ctx := contextx.WithContext(c)
	ctx.Info("receive signal to stop server")

	if err := i.server.Stop(ctx); err != nil {
		ctx.Error("Failed to stop server", zap.Error(err))
		return fmt.Errorf("failed to stop server: %w", err)
	}

	return nil
}

// NewInitServersFn creates a new impl server init function.
func NewInitServersFn(
	blockServer blockB.BlockServiceServer,
	networkServer netB.NetworkServiceServer,
	txServer txB.TransactionServiceServer,
	accountServer accountB.AccountServiceServer,
) grpcx.InitServers {
	return func(s *grpc.Server) {
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_SERVING)

		blockB.RegisterBlockServiceServer(s, blockServer)
		netB.RegisterNetworkServiceServer(s, networkServer)
		txB.RegisterTransactionServiceServer(s, txServer)
		accountB.RegisterAccountServiceServer(s, accountServer)

		reflection.Register(s)
	}
}
