//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"context"

	"github.com/blackhorseya/ryze/entity/domain/block/model"
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
