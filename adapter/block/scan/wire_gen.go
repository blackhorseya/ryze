// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package scan

import (
	"errors"
	"fmt"
	"github.com/blackhorseya/ryze/app/domain/block/biz"
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/spf13/viper"
)

import (
	_ "github.com/blackhorseya/ryze/api/block/scan"
)

// Injectors from wire.go:

func New(v *viper.Viper) (adapterx.Service, error) {
	configuration, err := configx.NewConfiguration(v)
	if err != nil {
		return nil, err
	}
	application, err := initApplication(configuration)
	if err != nil {
		return nil, err
	}
	client, err := grpcx.NewClient(configuration)
	if err != nil {
		return nil, err
	}
	blockServiceClient, err := biz.NewBlockServiceClient(client)
	if err != nil {
		return nil, err
	}
	service := NewService(application, blockServiceClient)
	return service, nil
}

// wire.go:

func initApplication(config *configx.Configuration) (*configx.Application, error) {
	app, ok := config.Services["block-scan"]
	if !ok {
		return nil, errors.New("[block-scan] service not found")
	}

	err := otelx.SetupOTelSDK(contextx.Background(), app)
	if err != nil {
		return nil, fmt.Errorf("failed to setup OpenTelemetry SDK: %w", err)
	}

	return app, nil
}
