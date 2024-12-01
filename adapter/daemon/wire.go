//go:build wireinject

//go:generate wire

package daemon

import (
	"fmt"

	block2 "github.com/blackhorseya/ryze/internal/app/domain/block"
	transaction2 "github.com/blackhorseya/ryze/internal/app/domain/transaction"
	configx2 "github.com/blackhorseya/ryze/internal/app/infra/configx"
	"github.com/blackhorseya/ryze/internal/app/infra/otelx"
	"github.com/blackhorseya/ryze/internal/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/internal/app/infra/storage/pgx"
	"github.com/blackhorseya/ryze/internal/app/infra/tonx"
	grpcx2 "github.com/blackhorseya/ryze/internal/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/eventx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

const serviceName = "daemon"

// InitApplication is a function to initialize application.
func InitApplication(config *configx2.Configuration) (*configx2.Application, error) {
	return config.GetService(serviceName)
}

// InitTonClient is used to initialize the ton client.
func InitTonClient(config *configx2.Configuration) (*tonx.Client, error) {
	settings, ok := config.Networks["ton"]
	if !ok {
		return nil, fmt.Errorf("network [ton] not found")
	}

	n := "mainnet"
	if settings.Testnet {
		n = "testnet"
	}

	return tonx.NewClient(tonx.Options{Network: n})
}

func New(v *viper.Viper) (adapterx.Server, func(), error) {
	panic(wire.Build(
		NewServer,
		wire.Struct(new(Injector), "*"),

		// config
		configx2.NewConfiguration,
		InitApplication,
		otelx.SetupSDK,

		// event
		eventx.NewInMemoryEventBus,

		// storage
		mongodbx.NewClientWithClean,
		pgx.NewClient,

		// transports
		grpcx2.NewServer,
		grpcx2.NewClient,
		InitTonClient,

		// app layer
		NewInitServersFn,

		// domain layer
		block2.ProviderSet,
		block2.NewBlockServiceClient,
		transaction2.ProviderSet,
		transaction2.NewTransactionServiceClient,
	))
}
