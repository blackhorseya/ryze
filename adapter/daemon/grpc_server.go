package daemon

import (
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// NewInitServersFn is a function to create a new init servers function.
func NewInitServersFn() grpcx.InitServers {
	return func(s *grpc.Server) {
		// register health server
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_SERVING)

		// register reflection service on gRPC server.
		reflection.Register(s)

		// TODO: 2024/9/14|sean|register grpc server
	}
}
