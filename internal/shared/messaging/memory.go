package messaging

import (
	"errors"
	"sync"

	"github.com/blackhorseya/ryze/pkg/eventx"
)

// InMemoryEventBus 是一个基于内存的事件总线实现
type InMemoryEventBus struct {
	mu       sync.Mutex
	handlers []eventx.EventHandler
}

// NewInMemoryEventBus 创建一个新的基于内存的 EventBus 实例
func NewInMemoryEventBus() eventx.EventBus {
	return &InMemoryEventBus{
		handlers: make([]eventx.EventHandler, 0),
	}
}

// Subscribe 用于订阅事件处理器
func (bus *InMemoryEventBus) Subscribe(handler eventx.EventHandler) error {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	bus.handlers = append(bus.handlers, handler)
	return nil
}

// Publish 用于发布事件，通知所有订阅的处理器
func (bus *InMemoryEventBus) Publish(event eventx.DomainEvent) error {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	if len(bus.handlers) == 0 {
		return errors.New("no handlers subscribed for event: " + event.GetName())
	}

	// 通知所有订阅的处理器
	for _, handler := range bus.handlers {
		go handler.Handle(event) // 使用 goroutine 异步处理事件
	}

	return nil
}
