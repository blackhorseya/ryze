package grpc

import (
	"github.com/blackhorseya/ryze/adapter/block/wirex"
	"github.com/blackhorseya/ryze/pkg/adapterx"
)

type grpc struct {
}

// NewGRPC creates a new grpc service.
func NewGRPC(injector *wirex.Injector) adapterx.Service {
	return &grpc{}
}

func (i *grpc) Start() error {
	// TODO: 2024/7/29|sean|add grpc logic here
	panic("implement me")
}

func (i *grpc) AwaitSignal() error {
	// TODO: 2024/7/29|sean|add grpc logic here
	panic("implement me")
}
