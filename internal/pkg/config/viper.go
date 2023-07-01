package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// NewViper serve caller to get a new viper.Viper instance
func NewViper(path string) (*viper.Viper, error) {
	v := viper.New()

	if path != "" {
		v.SetConfigFile(path)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, errors.Wrap(err, "get user home dir error")
		}

		v.AddConfigPath(home)
		v.SetConfigType("yaml")
		v.SetConfigName(".ryze")
	}

	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "read config file error")
	}

	fmt.Printf("read config file success, path: %s\n", v.ConfigFileUsed())

	return v, nil
}
