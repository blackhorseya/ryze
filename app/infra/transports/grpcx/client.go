package grpcx

import (
	"errors"
	"fmt"
	"time"

	"github.com/blackhorseya/ryze/app/infra/configx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
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

	// Retry options
	retryOptions := []grpc_retry.CallOption{
		grpc_retry.WithMax(3),                           // Retry up to 3 times
		grpc_retry.WithPerRetryTimeout(2 * time.Second), // Timeout for each retry
	}

	target := fmt.Sprintf("%s:%d", app.GRPC.URL, app.GRPC.Port)
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_prometheus.UnaryClientInterceptor,
			grpc_retry.UnaryClientInterceptor(retryOptions...),
		)),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_prometheus.StreamClientInterceptor,
			grpc_retry.StreamClientInterceptor(retryOptions...),
		)),
	}

	return grpc.NewClient(target, options...)
}
