//go:build wireinject

//go:generate wire

package scan

import (
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func New(v *viper.Viper) (adapterx.Server, func(), error) {
	panic(wire.Build(
		NewServer,
	))
}
