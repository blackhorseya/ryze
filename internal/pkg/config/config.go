package config

import (
	"fmt"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// NewConfig serve caller to create a viper.Viper
func NewConfig(path string) (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
	)

	v.AddConfigPath(".")
	v.SetConfigFile(path)

	if err = v.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "read config file error")
	}

	fmt.Printf("read config file success, path: %s\n", v.ConfigFileUsed())

	return v, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewConfig)
