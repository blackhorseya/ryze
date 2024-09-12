package eventx

import (
	"sync"
	"time"
)

// DomainEvent is the interface for domain events.
type DomainEvent interface {
	GetOccurredAt() time.Time
	GetName() string
	GetVersion() int
}

// EventBus is the interface for event bus.
type EventBus struct {
	subscribers []chan DomainEvent
	mu          sync.Mutex
}

// NewEventBus is used to create a new event bus.
func NewEventBus() *EventBus {
	return &EventBus{}
}

// Subscribe is used to subscribe to the event bus.
func (eb *EventBus) Subscribe(ch chan DomainEvent) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.subscribers = append(eb.subscribers, ch)
}

// Publish is used to publish an event to the event bus.
func (eb *EventBus) Publish(event DomainEvent) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	for _, subscriber := range eb.subscribers {
		go func(sub chan DomainEvent) {
			sub <- event
		}(subscriber)
	}
}
