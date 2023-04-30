package biz

import (
	"github.com/blackhorseya/ryze/pkg/contextx"
	bb "github.com/blackhorseya/ryze/pkg/entity/domain/block/biz"
	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
)

type impl struct {
}

// NewImpl will create an object that implement IBiz interface
func NewImpl() bb.IBiz {
	return &impl{}
}

func (i *impl) GetBlockByHash(ctx contextx.Contextx, hash []byte) (record *bm.Block, err error) {
	// todo: 2023/4/30|sean|impl me
	panic("implement me")
}

func (i *impl) ListBlocks(ctx contextx.Contextx, condition bb.ListBlocksCondition) (records []*bm.Block, total int, err error) {
	// todo: 2023/4/30|sean|impl me
	panic("implement me")
}

func (i *impl) ListenNewBlock(ctx contextx.Contextx) error {
	// todo: 2023/4/30|sean|impl me
	panic("implement me")
}
