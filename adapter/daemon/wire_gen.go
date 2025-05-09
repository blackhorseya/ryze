// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package daemon

import (
	"fmt"

	"github.com/blackhorseya/ryze/internal/domain/block"
	"github.com/blackhorseya/ryze/internal/domain/transaction"
	"github.com/blackhorseya/ryze/internal/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/internal/infra/storage/pgx"
	"github.com/blackhorseya/ryze/internal/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/internal/shared/configx"
	"github.com/blackhorseya/ryze/internal/shared/messaging"
	"github.com/blackhorseya/ryze/internal/shared/otelx"
	"github.com/blackhorseya/ryze/internal/shared/tonx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/spf13/viper"
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
	sdk, cleanup, err := otelx.SetupSDK(application)
	if err != nil {
		return nil, nil, err
	}
	client, err := grpcx.NewClient(configuration)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	blockServiceClient, err := block.NewBlockServiceClient(client)
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
		C:           configuration,
		A:           application,
		OTelx:       sdk,
		blockClient: blockServiceClient,
		txClient:    transactionServiceClient,
	}
	tonxClient, err := InitTonClient(configuration)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mongoClient, cleanup2, err := mongodbx.NewClientWithClean(application)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	iBlockRepo, err := mongodbx.NewBlockRepo(mongoClient)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	blockServiceServer := block.NewBlockService(tonxClient, iBlockRepo)
	db, err := pgx.NewClient(application)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	iTransactionRepo, err := pgx.NewTransactionRepo(db)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	transactionServiceServer := transaction.NewTransactionService(tonxClient, iTransactionRepo)
	initServers := NewInitServersFn(blockServiceServer, transactionServiceServer)
	server, err := grpcx.NewServer(application, initServers)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	eventBus := messaging.NewInMemoryEventBus()
	adapterxServer, cleanup3, err := NewServer(injector, server, eventBus)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return adapterxServer, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

const serviceName = "daemon"

// InitApplication is a function to initialize application.
func InitApplication(config *configx.Configuration) (*configx.Application, error) {
	return config.GetService(serviceName)
}

// InitTonClient is used to initialize the ton client.
func InitTonClient(config *configx.Configuration) (*tonx.Client, error) {
	settings, ok := config.Networks["ton"]
	if !ok {
		return nil, fmt.Errorf("network [ton] not found")
	}

	n := "mainnet"
	if settings.Testnet {
		n = "testnet"
	}

	return tonx.NewClient(tonx.Options{Network: n})
}
