package block

import (
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/entity/domain/block/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"gorm.io/gorm"
)

type rdbms struct {
	rw *gorm.DB
}

// NewGORM is used to create an implementation of the block repository.
func NewGORM(rw *gorm.DB) repo.IBlockRepo {
	return &rdbms{rw: rw}
}

func (i *rdbms) GetByID(ctx contextx.Contextx, id string) (item *model.Block, err error) {
	// TODO: 2024/8/13|sean|implement me
	panic("implement me")
}

func (i *rdbms) Create(ctx contextx.Contextx, item *model.Block) (err error) {
	// TODO: 2024/8/13|sean|implement me
	panic("implement me")
}

func (i *rdbms) List(ctx contextx.Contextx, cond repo.ListCondition) (items []*model.Block, total int, err error) {
	// TODO: 2024/8/13|sean|implement me
	panic("implement me")
}
