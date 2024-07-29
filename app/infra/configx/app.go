package configx

import (
	"fmt"

	"github.com/blackhorseya/ryze/app/infra/transports/httpx"
	"github.com/blackhorseya/ryze/pkg/netx"
)

// Application is the application configuration.
type Application struct {
	Name string `json:"name" yaml:"name"`

	HTTP httpx.Options `json:"http" yaml:"http"`
	GRPC GRPC          `json:"grpc" yaml:"grpc"`

	OTel struct {
		Target string `json:"target" yaml:"target"`
	} `json:"otel" yaml:"otel"`
}

// GRPC is the gRPC configuration.
type GRPC struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

// GetAddr is used to get the gRPC address.
func (x *GRPC) GetAddr() string {
	if x.Host == "" {
		x.Host = "0.0.0.0"
	}

	if x.Port == 0 {
		x.Port = netx.GetAvailablePort()
	}

	return fmt.Sprintf("%s:%d", x.Host, x.Port)
}
