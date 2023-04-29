package block

import (
	"github.com/blackhorseya/ryze/pkg/adapter"
	"github.com/google/wire"
)

type impl struct {
}

// NewImpl returns a new block listener implementation.
func NewImpl() adapter.Listener {
	return &impl{}
}

func (i *impl) Start() error {
	// todo: 2023/4/29|sean|impl me
	return nil
}

func (i *impl) Stop() error {
	// todo: 2023/4/29|sean|impl me
	return nil
}

// ListenerSet presents a listener set.
var ListenerSet = wire.NewSet(NewImpl)
