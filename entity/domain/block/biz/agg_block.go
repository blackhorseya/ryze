package biz

import (
	"github.com/blackhorseya/ryze/entity/domain/block/model"
)

// Block is the aggregate root of block domain
type Block struct {
	model.Block
}

func NewBlock(workchain int32, shard int64, seqno uint32) (*Block, error) {
	b, err := model.NewBlock(workchain, shard, seqno)
	if err != nil {
		return nil, err
	}

	return &Block{
		Block: *b,
	}, nil
}
