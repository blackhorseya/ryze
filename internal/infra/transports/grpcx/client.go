package grpcx

import (
	"errors"
	"fmt"

	"github.com/blackhorseya/ryze/internal/infra/configx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
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

	if app.GRPC.URL == "" || app.GRPC.Port == 0 {
		return nil, errors.New("grpc url or port is empty")
	}

	target := fmt.Sprintf("%s:%d", app.GRPC.URL, app.GRPC.Port)
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_prometheus.UnaryClientInterceptor,
		)),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_prometheus.StreamClientInterceptor,
		)),
	}

	return grpc.NewClient(target, options...)
}
