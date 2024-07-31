package model

import (
	"fmt"
)

// NewBlock is used to create a new block.
func NewBlock(workchain int32, shard int64, seqno uint32) (*Block, error) {
	return &Block{
		Id:        fmt.Sprintf("%d:%d:%d", workchain, shard, seqno),
		Workchain: workchain,
		Shard:     shard,
		SeqNo:     seqno,
		// TODO: 2024/7/31|sean|add timestamp field
	}, nil
}
