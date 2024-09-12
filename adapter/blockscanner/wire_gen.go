// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package blockscanner

import (
	"fmt"
	"github.com/blackhorseya/ryze/app/domain/block"
	"github.com/blackhorseya/ryze/app/domain/transaction"
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/eventx"
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
	sdk, cleanup, err := otelx.NewSDK(application)
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
	injector := &Injector{
		C:           configuration,
		A:           application,
		OTel:        sdk,
		blockClient: blockServiceClient,
	}
	tonxClient, err := InitTonClient(configuration)
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
	transactionServiceClient, err := transaction.NewTransactionServiceClient(client)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	blockServiceServer := block.NewBlockService(tonxClient, iBlockRepo, eventBus, transactionServiceClient)
	initServers := NewInitServersFn(blockServiceServer)
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

const serviceName = "block-scanner"

// NewInitServersFn creates a new grpc server initializer.
func NewInitServersFn(blockServer biz.BlockServiceServer) grpcx.InitServers {
	return func(s *grpc.Server) {

		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_SERVING)
		reflection.Register(s)
		biz.RegisterBlockServiceServer(s, blockServer)
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

// InitTonClient is used to initialize the ton client.
func InitTonClient(config *configx.Configuration) (*tonx.Client, error) {
	settings, ok := config.Networks["ton"]
	if !ok {
		return nil, fmt.Errorf("network [ton] not found")
	}

	network := "mainnet"
	if settings.Testnet {
		network = "testnet"
	}

	return tonx.NewClient(tonx.Options{Network: network})
}