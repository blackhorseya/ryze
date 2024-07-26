package configx

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Configuration is the application configuration
type Configuration struct {
	Networks map[string]*Network `json:"networks" yaml:"networks"`
}

// NewConfiguration creates a new configuration.
func NewConfiguration(v *viper.Viper) (*Configuration, error) {
	configFile := viper.GetString("config")
	if configFile == "" {
		home, _ := os.UserHomeDir()
		if home == "" {
			home = "/root"
		}
		configFile = home + "/.config/ryze/.ryze.yaml"
	}

	v.SetConfigFile(configFile)

	err := v.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := new(Configuration)
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return config, nil
}
