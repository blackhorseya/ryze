// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package grpc

import (
	"fmt"

	"github.com/blackhorseya/ryze/adapter/platform/wirex"
	biz4 "github.com/blackhorseya/ryze/app/domain/account"
	block2 "github.com/blackhorseya/ryze/app/domain/block"
	"github.com/blackhorseya/ryze/app/domain/block/repo/block"
	biz2 "github.com/blackhorseya/ryze/app/domain/network"
	biz3 "github.com/blackhorseya/ryze/app/domain/transaction"
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
	injector := &wirex.Injector{
		C: configuration,
		A: application,
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
	blockServiceServer := block2.NewBlockService(client, iBlockRepo)
	networkServiceServer := biz2.NewNetworkService(client)
	transactionServiceServer := biz3.NewTransactionService(client)
	accountServiceServer := biz4.NewAccountService(client)
	initServers := NewInitServersFn(blockServiceServer, networkServiceServer, transactionServiceServer, accountServiceServer)
	server, err := grpcx.NewServer(application, initServers)
	if err != nil {
		return nil, err
	}
	service := NewGRPC(injector, server)
	return service, nil
}

// wire.go:

var serviceName = "platform-grpc"

func initApplication(config *configx.Configuration) (*configx.Application, error) {
	app, err := config.GetService(serviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s config: %w", serviceName, err)
	}

	err = otelx.SetupOTelSDK(contextx.Background(), app)
	if err != nil {
		return nil, fmt.Errorf("failed to setup OpenTelemetry SDK: %w", err)
	}

	return app, nil
}

func initTonx() (*tonx.Client, error) {
	return tonx.NewClient(tonx.Options{Network: "mainnet"})
}
