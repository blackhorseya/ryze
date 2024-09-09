package scan

import (
	"context"

	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
)

type impl struct {
	injector *Injector
	server   *grpcx.Server
}

// NewServer is used to create a new scan server
func NewServer(injector *Injector, server *grpcx.Server) adapterx.Server {
	return &impl{
		injector: injector,
		server:   server,
	}
}

func (i *impl) Start(c context.Context) error {
	// TODO: 2024/9/8|sean|implement me
	ctx := contextx.WithContext(c)
	ctx.Info("scan server start")

	return nil
}

func (i *impl) Shutdown(c context.Context) error {
	// TODO: 2024/9/8|sean|implement me
	ctx := contextx.WithContext(c)
	ctx.Info("scan server shutdown")

	return nil
}
