package platform

import (
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	accountB "github.com/blackhorseya/ryze/entity/domain/account/biz"
	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	netB "github.com/blackhorseya/ryze/entity/domain/network/biz"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// NewInitServersFn is a function to create a new init servers function.
func NewInitServersFn(
	blockServer blockB.BlockServiceServer,
	networkServer netB.NetworkServiceServer,
	txServer txB.TransactionServiceServer,
	accountServer accountB.AccountServiceServer,
) grpcx.InitServers {
	return func(s *grpc.Server) {
		// register health server
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_SERVING)

		// register reflection service on gRPC server.
		reflection.Register(s)

		// register our services
		blockB.RegisterBlockServiceServer(s, blockServer)
		netB.RegisterNetworkServiceServer(s, networkServer)
		txB.RegisterTransactionServiceServer(s, txServer)
		accountB.RegisterAccountServiceServer(s, accountServer)
	}
}
