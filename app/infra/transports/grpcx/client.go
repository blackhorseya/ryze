package grpcx

import (
	"fmt"

	"github.com/blackhorseya/ryze/app/infra/configx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client is the grpc client
type Client struct {
	services map[string]*configx.Application
}

// NewClient is used to create a new grpc client
func NewClient(config *configx.Configuration) (*Client, error) {
	return &Client{
		services: config.Services,
	}, nil
}

// Dial is used to dial the grpc service
func (c *Client) Dial(service string) (*grpc.ClientConn, error) {
	app, ok := c.services[service]
	if !ok {
		return nil, fmt.Errorf("service: [%s] not found", service)
	}

	target := fmt.Sprintf("localhost:%d", app.GRPC.Port)
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_prometheus.UnaryClientInterceptor,
		)),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_prometheus.StreamClientInterceptor,
		)),
	}

	return grpc.NewClient(target, options...)
}
