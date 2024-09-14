package grpcx

import (
	"github.com/blackhorseya/ryze/pkg/eventx"
)

type eventBus struct {
}

// NewEventBus creates a new gRPC event bus.
func NewEventBus() eventx.EventBus {
	return &eventBus{}
}

func (bus *eventBus) Subscribe(handler eventx.EventHandler) error {
	// TODO: 2024/9/14|sean|implement me
	panic("implement me")
}

func (bus *eventBus) Publish(event eventx.DomainEvent) error {
	// TODO: 2024/9/14|sean|implement me
	panic("implement me")
}
