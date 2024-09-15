package model

import (
	"fmt"
	"time"

	"github.com/blackhorseya/ryze/pkg/eventx"
)

// NewBlock is used to create a new block.
func NewBlock(workchain int32, shard int64, seqno uint32) (*Block, error) {
	return &Block{
		Id:        fmt.Sprintf("%d:%x:%d", workchain, uint64(shard), seqno),
		Workchain: workchain,
		Shard:     shard,
		SeqNo:     seqno,
		// TODO: 2024/7/31|sean|add timestamp field
	}, nil
}

// Born is used to born a block.
func (x *Block) Born() eventx.DomainEvent {
	return &FoundBlockEvent{
		Block:      x,
		OccurredAt: time.Now(),
		Version:    1,
	}
}
