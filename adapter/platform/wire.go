//go:build wireinject

//go:generate wire

package platform

import (
	"fmt"

	"github.com/blackhorseya/ryze/app/domain/account"
	"github.com/blackhorseya/ryze/app/domain/block"
	"github.com/blackhorseya/ryze/app/domain/network"
	"github.com/blackhorseya/ryze/app/domain/transaction"
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

var serviceName = "platform"

func initApplication(config *configx.Configuration) (*configx.Application, error) {
	app, err := config.GetService(serviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s config: %w", serviceName, err)
	}

	return app, nil
}

func initTonx() (*tonx.Client, error) {
	return tonx.NewClient(tonx.Options{Network: "mainnet"})
}

func New(v *viper.Viper) (adapterx.Server, func(), error) {
	panic(wire.Build(
		wire.Struct(new(Injector), "*"),

		NewServer,
		configx.NewConfiguration,
		initApplication,
		grpcx.NewServer,
		NewInitServersFn,
		otelx.NewSDK,
		eventx.NewEventBus,
		grpcx.NewClient,

		account.ProviderSet,
		block.ProviderSet,
		network.ProviderSet,
		transaction.ProviderSet,
		transaction.NewTransactionServiceClient,

		initTonx,
		mongodbx.NewClient,
	))
}