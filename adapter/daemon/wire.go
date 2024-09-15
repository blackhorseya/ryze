//go:build wireinject

//go:generate wire

package daemon

import (
	"fmt"

	"github.com/blackhorseya/ryze/app/domain/block"
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
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

		// config
		configx.NewConfiguration,
		InitApplication,
		otelx.SetupSDK,

		// event
		eventx.NewInMemoryEventBus,

		// storage
		mongodbx.NewClientWithClean,

		// transports
		grpcx.NewServer,
		grpcx.NewClient,
		InitTonClient,

		// app layer
		NewInitServersFn,

		// domain layer
		block.ProviderSet,
		block.NewBlockServiceClient,
	))
}
