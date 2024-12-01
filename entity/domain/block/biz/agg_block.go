package biz

import (
	"fmt"

	"github.com/blackhorseya/ryze/entity/domain/block/model"
)

// Block is the aggregate root of block domain
type Block struct {
	model.Block
}

// NewBlock is used to create a new block.
func NewBlock(workchain int32, shard int64, seqno uint32) (*Block, error) {
	return &Block{
		Block: model.Block{
			Id:        fmt.Sprintf("%d:%d:%d", workchain, shard, seqno),
			Workchain: workchain,
			Shard:     shard,
			SeqNo:     seqno,
		},
	}, nil
}
