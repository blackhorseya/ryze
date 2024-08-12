//go:build wireinject

//go:generate wire

package grpc

import (
	"fmt"

	"github.com/blackhorseya/ryze/adapter/platform/wirex"
	blockB "github.com/blackhorseya/ryze/app/domain/block/biz"
	netB "github.com/blackhorseya/ryze/app/domain/network/biz"
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var serviceName = "platform-grpc"

func initApplication(config *configx.Configuration) (*configx.Application, error) {
	app, err := config.GetService(serviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s config: %w", serviceName, err)
	}

	err = otelx.SetupOTelSDK(contextx.Background(), app)
	if err != nil {
		return nil, fmt.Errorf("failed to setup OpenTelemetry SDK: %w", err)
	}

	return app, nil
}

func initTonx() (*tonx.Client, error) {
	return tonx.NewClient(tonx.Options{Network: "mainnet"})
}

func New(v *viper.Viper) (adapterx.Service, error) {
	panic(wire.Build(
		wire.Struct(new(wirex.Injector), "*"),

		NewGRPC,
		configx.NewConfiguration,
		initApplication,
		grpcx.NewServer,
		NewInitServersFn,

		blockB.ProviderBlockServiceSet,
		netB.NewNetworkService,

		initTonx,
		mongodbx.NewClient,
	))
}
