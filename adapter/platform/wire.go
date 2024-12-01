//go:build wireinject

//go:generate wire

package platform

import (
	"fmt"

	"github.com/blackhorseya/ryze/internal/app/domain/account"
	"github.com/blackhorseya/ryze/internal/app/domain/block"
	"github.com/blackhorseya/ryze/internal/app/domain/network"
	transaction2 "github.com/blackhorseya/ryze/internal/app/domain/transaction"
	"github.com/blackhorseya/ryze/internal/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/internal/app/infra/storage/pgx"
	grpcx2 "github.com/blackhorseya/ryze/internal/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/internal/shared/configx"
	"github.com/blackhorseya/ryze/internal/shared/otelx"
	"github.com/blackhorseya/ryze/internal/shared/tonx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

const serviceName = "platform"

func initApplication(config *configx.Configuration) (*configx.Application, error) {
	app, err := config.GetService(serviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s config: %w", serviceName, err)
	}

	return app, nil
}

// InitTonClient is used to initialize the ton client.
func InitTonClient(config *configx.Configuration) (*tonx.Client, error) {
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
		wire.Struct(new(Injector), "*"),

		NewServer,
		configx.NewConfiguration,
		initApplication,
		grpcx2.NewServer,
		NewInitServersFn,
		otelx.SetupSDK,
		grpcx2.NewClient,
		pgx.NewClient,

		account.ProviderSet,
		block.ProviderSet,
		network.ProviderSet,
		transaction2.ProviderSet,
		transaction2.NewTransactionServiceClient,

		InitTonClient,
		mongodbx.NewClient,
	))
}
