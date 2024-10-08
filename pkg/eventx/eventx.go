package eventx

import (
	"time"
)

// DomainEvent represents a generic event with basic properties.
type DomainEvent interface {
	GetOccurredAt() time.Time
	GetName() string
	GetVersion() int
}

// EventHandler is the interface for handling domain events.
type EventHandler interface {
	Handle(event DomainEvent)
}

// EventBus is the interface for a pluggable event bus.
type EventBus interface {
	Subscribe(handler EventHandler) error
	Publish(event DomainEvent) error
}

// // EventBus is the structure that handles event subscriptions and publishing.
// type EventBus struct {
// 	subscribers []chan DomainEvent
// 	handlers    []EventHandler
// 	mu          sync.Mutex
// }
//
// // NewEventBus creates a new event bus.
// func NewEventBus() *EventBus {
// 	return &EventBus{
// 		subscribers: make([]chan DomainEvent, 0),
// 		handlers:    make([]EventHandler, 0),
// 	}
// }
//
// // SubscribeChannel allows channels to subscribe to the event bus.
// func (eb *EventBus) SubscribeChannel(ch chan DomainEvent) {
// 	eb.mu.Lock()
// 	defer eb.mu.Unlock()
// 	eb.subscribers = append(eb.subscribers, ch)
// }
//
// // SubscribeHandler allows event handlers to subscribe to the event bus.
// func (eb *EventBus) SubscribeHandler(handler EventHandler) {
// 	eb.mu.Lock()
// 	defer eb.mu.Unlock()
// 	eb.handlers = append(eb.handlers, handler)
// }
//
// // Publish publishes an event to all subscribed channels and handlers.
// func (eb *EventBus) Publish(event DomainEvent) {
// 	eb.mu.Lock()
// 	defer eb.mu.Unlock()
//
// 	// Send to all channels asynchronously
// 	for _, sub := range eb.subscribers {
// 		go func(ch chan DomainEvent) {
// 			ch <- event
// 		}(sub)
// 	}
//
// 	// Send to all handlers
// 	for _, handler := range eb.handlers {
// 		go handler.Handle(event)
// 	}
// }
