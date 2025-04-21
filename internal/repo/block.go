//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"context"
	"time"

	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
)

// ListBlockCondition is the condition for list block.
type ListBlockCondition struct {
	Start  time.Time
	End    time.Time
	Limit  int64
	Offset int64
}

type (
	// BlockCreator is the creator for block.
	BlockCreator interface {
		Create(c context.Context, block *biz.Block) error
	}

	// BlockGetter is the getter for block.
	BlockGetter interface {
		GetByID(c context.Context, id string) (*biz.Block, error)
		List(c context.Context, cond ListBlockCondition) ([]*biz.Block, int, error)
	}

	// BlockRepo is the repository for block.
	BlockRepo interface {
		BlockCreator
		BlockGetter
	}
)

// ScanBlockRequest is the request for scan block.
type ScanBlockRequest struct {
	StartHeight uint32
	EndHeight   uint32
}

type (
	BlockAdapter interface {
		ScanBlock(c context.Context, req ScanBlockRequest, blockCh chan<- *biz.Block) error
	}
)

// ListCondition is the condition for list.
type ListCondition struct {
	Limit int64
	Skip  int64
}

// IBlockRepo is the interface for block repository.
type IBlockRepo interface {
	GetByID(c context.Context, id string) (item *model.Block, err error)
	Create(c context.Context, item *model.Block) (err error)
	List(c context.Context, condition ListCondition) (items []*model.Block, total int, err error)
}
