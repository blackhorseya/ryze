package dao

import (
	"time"

	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Block declare the block record dao
type Block struct {
	Number     uint64    `json:"number" db:"number"`
	Hash       []byte    `json:"hash" db:"hash"`
	ParentHash []byte    `json:"parent_hash" db:"parent_hash"`
	Timestamp  time.Time `json:"timestamp" db:"timestamp"`
}

// NewBlock serve caller to get a new Block instance
func NewBlock(block *bm.Block) *Block {
	return &Block{
		Number:     block.Number,
		Hash:       common.HexToHash(block.Hash).Bytes(),
		ParentHash: common.HexToHash(block.ParentHash).Bytes(),
		Timestamp:  block.Timestamp.AsTime().UTC(),
	}
}

// ToEntity serve caller to convert Block to *bm.Block
func (b *Block) ToEntity() *bm.Block {
	return &bm.Block{
		Number:     b.Number,
		Hash:       common.BytesToHash(b.Hash).Hex(),
		ParentHash: common.BytesToHash(b.ParentHash).Hex(),
		Timestamp:  timestamppb.New(b.Timestamp),
	}
}

// Blocks declare the block record dao slice
type Blocks []*Block

// ToEntity serve caller to convert Blocks to []*bm.Block
func (slice Blocks) ToEntity() []*bm.Block {
	var ret []*bm.Block
	for _, block := range slice {
		ret = append(ret, block.ToEntity())
	}

	return ret
}
