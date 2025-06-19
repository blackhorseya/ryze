package platform

import (
	"context"
	"fmt"

	"github.com/blackhorseya/ryze/internal/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.uber.org/zap"
)

type impl struct {
	injector *Injector
	server   *grpcx.Server
}

// NewServer creates a new impl service.
func NewServer(injector *Injector, server *grpcx.Server) adapterx.Server {
	return &impl{
		injector: injector,
		server:   server,
	}
}

func (i *impl) Start(c context.Context) error {
	ctx := contextx.WithContext(c)

	err := i.server.Start(ctx)
	if err != nil {
		ctx.Error(
			"Failed to start grpc server",
			zap.Error(err),
			zap.String("addr", i.injector.A.GRPC.GetAddr()),
		)
		return err
	}

	ctx.Info("start grpc server")

	return nil
}

func (i *impl) Shutdown(c context.Context) error {
	ctx := contextx.WithContext(c)
	ctx.Info("receive signal to stop server")

	if err := i.server.Stop(ctx); err != nil {
		ctx.Error("Failed to stop server", zap.Error(err))
		return fmt.Errorf("failed to stop server: %w", err)
	}

	return nil
}
