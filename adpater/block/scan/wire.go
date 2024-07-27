//go:build wireinject

//go:generate wire

package scan

import (
	"github.com/blackhorseya/ryze/adpater/block/wirex"
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/transports/httpx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func initApplication(config *configx.Configuration) (*configx.Application, error) {
	return config.Services["block-scan"], nil
}

func initServer(app *configx.Application) (*httpx.Server, error) {
	return httpx.NewServer(app.HTTP)
}

func New(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(
		wire.Struct(new(wirex.Injector), "*"),
		configx.NewConfiguration,
		initApplication,

		NewRestful,
		initServer,
	))
}
