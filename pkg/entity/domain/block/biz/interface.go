//go:generate mockgen -destination=./mock_${GOFILE} -package=biz -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/ryze/pkg/contextx"
	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
)

type ListBlocksCondition struct {
	Page uint
	Size uint
}

type IBiz interface {
	// GetBlockByHash serve caller to given block hash to get a block
	GetBlockByHash(ctx contextx.Contextx, hash []byte) (record *bm.Block, err error)

	// ListBlocks serve caller to given conditions to list all blocks
	ListBlocks(ctx contextx.Contextx, condition ListBlocksCondition) (records []*bm.Block, total int, err error)

	// ListenNewBlock serve caller to listen new block
	ListenNewBlock(ctx contextx.Contextx) error
}
