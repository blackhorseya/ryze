//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/contextx"
)

// ListCondition is the condition for list.
type ListCondition struct {
	Limit int
	Skip  int
}

// IBlockRepo is the interface for block repository.
type IBlockRepo interface {
	GetByID(ctx contextx.Contextx, id string) (item *model.Block, err error)
	Create(ctx contextx.Contextx, item *model.Block) (err error)
	List(ctx contextx.Contextx, condition ListCondition) (items []*model.Block, total int, err error)
}
