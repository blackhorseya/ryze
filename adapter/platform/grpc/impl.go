package grpc

import (
	"github.com/blackhorseya/ryze/adapter/platform/wirex"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type impl struct {
}

// NewGRPC creates a new impl service.
func NewGRPC(injector *wirex.Injector, server *grpcx.Server) adapterx.Service {
	return &impl{}
}

func (i *impl) Start(ctx contextx.Contextx) error {
	// TODO: 2024/8/12|sean|implement me
	panic("implement me")
}

func (i *impl) AwaitSignal(ctx contextx.Contextx) error {
	// TODO: 2024/8/12|sean|implement me
	panic("implement me")
}

// NewInitServersFn creates a new impl server init function.
func NewInitServersFn() grpcx.InitServers {
	return func(s *grpc.Server) {
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serviceName, grpc_health_v1.HealthCheckResponse_SERVING)

		reflection.Register(s)
	}
}
