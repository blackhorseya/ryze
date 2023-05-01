package dao

import (
	"time"

	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/ethereum/go-ethereum/common"
)

// Block declare the block record dao
type Block struct {
	Number     uint64      `json:"number" db:"number"`
	Hash       common.Hash `json:"hash" db:"hash"`
	ParentHash common.Hash `json:"parent_hash" db:"parent_hash"`
	Timestamp  time.Time   `json:"timestamp" db:"timestamp"`
}

// NewBlock serve caller to get a new Block instance
func NewBlock(block *bm.Block) *Block {
	return &Block{
		Number:     block.Number,
		Hash:       common.HexToHash(block.Hash),
		ParentHash: common.HexToHash(block.ParentHash),
		Timestamp:  block.Timestamp.AsTime().UTC(),
	}
}
