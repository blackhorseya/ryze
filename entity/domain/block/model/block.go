package model

import (
	"fmt"
	"time"

	"github.com/blackhorseya/ryze/pkg/eventx"
)

// NewBlock is used to create a new block.
func NewBlock(workchain int32, shard int64, seqno uint32) (*Block, error) {
	return &Block{
		Id:        fmt.Sprintf("%d:%d:%d", workchain, shard, seqno),
		Workchain: workchain,
		Shard:     shard,
		SeqNo:     seqno,
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
