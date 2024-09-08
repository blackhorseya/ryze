package scan

import (
	"context"

	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
)

type impl struct {
}

// NewServer is used to create a new scan server
func NewServer() adapterx.Server {
	return &impl{}
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
