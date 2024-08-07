// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package grpc

import (
	"errors"
	"fmt"
	"github.com/blackhorseya/ryze/adapter/block/wirex"
	"github.com/blackhorseya/ryze/app/domain/block/biz"
	"github.com/blackhorseya/ryze/app/domain/block/repo/block"
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/spf13/viper"
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
	client, err := initTonx()
	if err != nil {
		return nil, err
	}
	mongoClient, err := mongodbx.NewClient(application)
	if err != nil {
		return nil, err
	}
	iBlockRepo := block.NewMongoDB(mongoClient)
	blockServiceServer := biz.NewBlockService(client, iBlockRepo)
	injector := &wirex.Injector{
		C:            configuration,
		A:            application,
		BlockService: blockServiceServer,
	}
	initServers := NewInitServersFn(injector)
	server, err := grpcx.NewServer(application, initServers)
	if err != nil {
		return nil, err
	}
	service := NewGRPC(injector, server)
	return service, nil
}

// wire.go:

func initApplication(config *configx.Configuration) (*configx.Application, error) {
	app, ok := config.Services["block-grpc"]
	if !ok {
		return nil, errors.New("[block-grpc] service not found")
	}

	err := otelx.SetupOTelSDK(contextx.Background(), app)
	if err != nil {
		return nil, fmt.Errorf("failed to setup OpenTelemetry SDK: %w", err)
	}

	return app, nil
}

func initTonx() (*tonx.Client, error) {
	return tonx.NewClient(tonx.Options{Network: "mainnet"})
}
