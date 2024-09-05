package configx

import (
	"errors"
	"fmt"
	"os"

	"github.com/blackhorseya/ryze/pkg/logging"
	"github.com/spf13/viper"
)

// Configuration is the application configuration
type Configuration struct {
	Log logging.Options `json:"log" yaml:"log"`

	Networks map[string]*Network     `json:"networks" yaml:"networks"`
	Services map[string]*Application `json:"services" yaml:"services"`
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

	err = logging.Init(config.Log)
	if err != nil {
		return nil, fmt.Errorf("failed to init logging: %w", err)
	}

	return config, nil
}

// GetService is used to get the service by name.
func (x *Configuration) GetService(name string) (*Application, error) {
	service, ok := x.Services[name]
	if !ok {
		return nil, fmt.Errorf("service %s not found", name)
	}

	return service, nil
}
