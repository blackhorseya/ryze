package daemon

import (
	"context"

	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/eventx"
)

type impl struct {
	injector *Injector
	bus      *eventx.EventBus
}

// NewServer is a function to create a new server.
func NewServer(injector *Injector, bus *eventx.EventBus) (adapterx.Server, func(), error) {
	return &impl{
		injector: injector,
		bus:      bus,
	}, func() {}, nil
}

func (i *impl) Start(c context.Context) error {
	ctx := contextx.WithContext(c)
	ctx.Info("server start")

	// i.bus.SubscribeHandler(event.NewFoundBlockHandler())

	return nil
}

func (i *impl) Shutdown(c context.Context) error {
	ctx := contextx.WithContext(c)
	ctx.Info("server shutdown")

	return nil
}
