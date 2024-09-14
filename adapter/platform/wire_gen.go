// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package platform

import (
	"fmt"
	"github.com/blackhorseya/ryze/app/domain/account"
	"github.com/blackhorseya/ryze/app/domain/block"
	"github.com/blackhorseya/ryze/app/domain/network"
	"github.com/blackhorseya/ryze/app/domain/transaction"
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/app/infra/storage/pgx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/eventx"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func New(v *viper.Viper) (adapterx.Server, func(), error) {
	configuration, err := configx.NewConfiguration(v)
	if err != nil {
		return nil, nil, err
	}
	application, err := initApplication(configuration)
	if err != nil {
		return nil, nil, err
	}
	sdk, cleanup, err := otelx.SetupSDK(application)
	if err != nil {
		return nil, nil, err
	}
	client, err := grpcx.NewClient(configuration)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	transactionServiceClient, err := transaction.NewTransactionServiceClient(client)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	injector := &Injector{
		C:        configuration,
		A:        application,
		OTelx:    sdk,
		txClient: transactionServiceClient,
	}
	tonxClient, err := initTonx()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mongoClient, err := mongodbx.NewClient(application)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	iBlockRepo, err := mongodbx.NewBlockRepo(mongoClient)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	eventBus := eventx.NewEventBus()
	blockServiceServer := block.NewBlockService(tonxClient, iBlockRepo, eventBus)
	networkServiceServer := network.NewNetworkService(tonxClient)
	db, err := pgx.NewClient(application)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	iTransactionRepo, err := pgx.NewTransactionRepo(db)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	transactionServiceServer := transaction.NewTransactionService(tonxClient, iTransactionRepo)
	accountServiceServer := account.NewAccountService(tonxClient)
	initServers := NewInitServersFn(blockServiceServer, networkServiceServer, transactionServiceServer, accountServiceServer)
	server, err := grpcx.NewServer(application, initServers)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	adapterxServer := NewServer(injector, server)
	return adapterxServer, func() {
		cleanup()
	}, nil
}

// wire.go:

var serviceName = "platform"

func initApplication(config *configx.Configuration) (*configx.Application, error) {
	app, err := config.GetService(serviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s config: %w", serviceName, err)
	}

	return app, nil
}

func initTonx() (*tonx.Client, error) {
	return tonx.NewClient(tonx.Options{Network: "mainnet"})
}
