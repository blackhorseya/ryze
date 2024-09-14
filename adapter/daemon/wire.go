//go:build wireinject

//go:generate wire

package daemon

import (
	"fmt"

	"github.com/blackhorseya/ryze/internal/domain/account"
	"github.com/blackhorseya/ryze/internal/domain/block"
	"github.com/blackhorseya/ryze/internal/domain/network"
	"github.com/blackhorseya/ryze/internal/domain/transaction"
	"github.com/blackhorseya/ryze/internal/infra/configx"
	"github.com/blackhorseya/ryze/internal/infra/otelx"
	"github.com/blackhorseya/ryze/internal/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/internal/infra/storage/pgx"
	"github.com/blackhorseya/ryze/internal/infra/tonx"
	"github.com/blackhorseya/ryze/internal/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/eventx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

const serviceName = "daemon"

// InitApplication is a function to initialize application.
func InitApplication(config *configx.Configuration) (*configx.Application, error) {
	return config.GetService(serviceName)
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
		NewServer,
		wire.Struct(new(Injector), "*"),
		configx.NewConfiguration,
		InitApplication,
		otelx.SetupSDK,
		grpcx.NewServer,
		NewInitServersFn,

		// infra clients
		InitTonClient,
		mongodbx.NewClientWithClean,
		pgx.NewClient,
		eventx.NewEventBus,

		account.ProviderSet,
		block.ProviderSet,
		network.ProviderSet,
		transaction.ProviderSet,
	))
}
