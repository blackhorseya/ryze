//go:build wireinject

//go:generate wire

package scan

import (
	"fmt"

	"github.com/blackhorseya/ryze/app/domain/block/biz"
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func initApplication(config *configx.Configuration) (*configx.Application, error) {
	app, err := config.GetService("block-scan")
	if err != nil {
		return nil, fmt.Errorf("failed to get block-scan config: %w", err)
	}

	err = otelx.SetupOTelSDK(contextx.Background(), app)
	if err != nil {
		return nil, fmt.Errorf("failed to setup OpenTelemetry SDK: %w", err)
	}

	return app, nil
}

func New(v *viper.Viper) (adapterx.Service, error) {
	panic(wire.Build(
		configx.NewConfiguration,
		initApplication,

		NewService,
		biz.NewBlockServiceClient,
		grpcx.NewClient,
	))
}
