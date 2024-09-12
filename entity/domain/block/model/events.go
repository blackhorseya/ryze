package model

import (
	"time"

	"github.com/blackhorseya/ryze/pkg/eventx"
)

var _ eventx.DomainEvent = (*NewBlockEvent)(nil)

// NewBlockEvent is used to create a new block event.
type NewBlockEvent struct {
	BlockID   string
	Workchain int32
	Shard     int64
	SeqNo     uint32

	OccurredAt time.Time
	Version    int
}

func (x *NewBlockEvent) GetOccurredAt() time.Time {
	return x.OccurredAt
}

func (x *NewBlockEvent) GetName() string {
	return "NewBlockEvent"
}

func (x *NewBlockEvent) GetVersion() int {
	return x.Version
}
