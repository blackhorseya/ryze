package repo

import (
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
)

type impl struct {
}

// NewImpl serve caller to get a new IRepo implementation instance
func NewImpl() IRepo {
	return &impl{}
}

func (i *impl) GetBlockByHash(ctx contextx.Contextx, hash []byte) (record *model.Block, err error) {
	// todo: 2023/4/29|sean|implement me
	panic("implement me")
}

func (i *impl) GetBlockByHeight(ctx contextx.Contextx, height uint64) (record *model.Block, err error) {
	// todo: 2023/4/29|sean|implement me
	panic("implement me")
}

func (i *impl) ListenNewBlock(ctx contextx.Contextx) (newBlockChan <-chan *model.Block, err error) {
	// todo: 2023/4/29|sean|implement me
	panic("implement me")
}
