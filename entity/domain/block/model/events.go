package model

import (
	"time"

	"github.com/blackhorseya/ryze/pkg/eventx"
)

var _ eventx.DomainEvent = (*FoundBlockEvent)(nil)

// FoundBlockEvent is used to create a new block event.
type FoundBlockEvent struct {
	Block *Block

	OccurredAt time.Time
	Version    int
}

func (x *FoundBlockEvent) GetOccurredAt() time.Time {
	return x.OccurredAt
}

func (x *FoundBlockEvent) GetName() string {
	return "FoundBlockEvent"
}

func (x *FoundBlockEvent) GetVersion() int {
	return x.Version
}
