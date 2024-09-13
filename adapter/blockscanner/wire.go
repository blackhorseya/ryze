//go:build wireinject

//go:generate wire

package blockscanner

import (
	"fmt"

	"github.com/blackhorseya/ryze/app/domain/block"
	"github.com/blackhorseya/ryze/app/domain/transaction"
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/app/infra/storage/pgx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/eventx"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const serviceName = "block-scanner"

// NewInitServersFn creates a new grpc server initializer.
func NewInitServersFn(
	blockServer biz.BlockServiceServer,
	txServer txB.TransactionServiceServer,
) grpcx.InitServers {
	return func(s *grpc.Server) {
		// register health check service
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_SERVING)

		// register reflection service
		reflection.Register(s)

		// register servers
		biz.RegisterBlockServiceServer(s, blockServer)
		txB.RegisterTransactionServiceServer(s, txServer)
	}
}

// InitApplication is used to initialize the application.
func InitApplication(config *configx.Configuration) (*configx.Application, error) {
	app, err := config.GetService(serviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get service %s: %w", serviceName, err)
	}

	return app, nil
}

// InitTonClient is used to initialize the ton client.
func InitTonClient(config *configx.Configuration) (*tonx.Client, error) {
	settings, ok := config.Networks["ton"]
	if !ok {
		return nil, fmt.Errorf("network [ton] not found")
	}

	network := "mainnet"
	if settings.Testnet {
		network = "testnet"
	}

	return tonx.NewClient(tonx.Options{Network: network})
}

func New(v *viper.Viper) (adapterx.Server, func(), error) {
	panic(wire.Build(
		NewServer,
		wire.Struct(new(Injector), "*"),
		configx.NewConfiguration,
		InitApplication,
		grpcx.NewServer,
		grpcx.NewClient,
		NewInitServersFn,
		otelx.NewSDK,
		eventx.NewEventBus,
		pgx.NewClient,

		mongodbx.NewClient,
		InitTonClient,

		block.ProviderSet,
		block.NewBlockServiceClient,

		transaction.ProviderSet,
		transaction.NewTransactionServiceClient,
	))
}
