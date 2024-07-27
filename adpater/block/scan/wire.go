//go:build wireinject

//go:generate wire

package scan

import (
	"github.com/blackhorseya/ryze/adpater/block/wirex"
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func New(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(
		wire.Struct(new(wirex.Injector), "*"),
		configx.NewConfiguration,

		NewRestful,
	))
}
