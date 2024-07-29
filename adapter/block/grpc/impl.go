package grpc

import (
	"github.com/blackhorseya/ryze/adapter/block/wirex"
	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"google.golang.org/grpc"
)

type impl struct {
	injector *wirex.Injector
	server   *grpcx.Server
}

// NewGRPC creates a new impl service.
func NewGRPC(injector *wirex.Injector, server *grpcx.Server) adapterx.Service {
	return &impl{
		injector: injector,
		server:   server,
	}
}

func (i *impl) Start() error {
	// TODO: 2024/7/29|sean|add impl logic here
	panic("implement me")
}

func (i *impl) AwaitSignal() error {
	// TODO: 2024/7/29|sean|add impl logic here
	panic("implement me")
}

// NewInitServersFn creates a new impl server init function.
func NewInitServersFn(injector *wirex.Injector) grpcx.InitServers {
	return func(s *grpc.Server) {
		model.RegisterBlockServiceServer(s, injector.BlockService)
	}
}
