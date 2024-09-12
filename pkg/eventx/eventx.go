package eventx

import (
	"time"
)

// DomainEvent is the interface for domain events.
type DomainEvent interface {
	GetOccurredAt() time.Time
	GetName() string
	GetVersion() int
}
