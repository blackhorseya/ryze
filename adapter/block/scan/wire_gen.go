// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package scan

import (
	"fmt"
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Injectors from wire.go:

func New(v *viper.Viper) (adapterx.Server, func(), error) {
	configuration, err := configx.NewConfiguration(v)
	if err != nil {
		return nil, nil, err
	}
	application, err := InitApplication(configuration)
	if err != nil {
		return nil, nil, err
	}
	injector := &Injector{
		C: configuration,
		A: application,
	}
	initServers := NewInitServersFn()
	server, err := grpcx.NewServer(application, initServers)
	if err != nil {
		return nil, nil, err
	}
	adapterxServer := NewServer(injector, server)
	return adapterxServer, func() {
	}, nil
}

// wire.go:

const serviceName = "block-scanner"

// NewInitServersFn creates a new grpc server initializer.
func NewInitServersFn() grpcx.InitServers {
	return func(s *grpc.Server) {
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_SERVING)
		reflection.Register(s)
	}
}

// InitApplication is used to initialize the application.
func InitApplication(config *configx.Configuration) (*configx.Application, error) {
	app, err := config.GetService(serviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get service %s: %w", serviceName, err)
	}

	return app, nil
}
