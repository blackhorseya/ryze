package tonx

import (
	"context"

	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/internal/app/repo"
)

// BlockAdapterImpl is the implementation for block adapter.
type BlockAdapterImpl struct {
}

// NewBlockAdapterImpl is used to create a new block adapter implementation.
func NewBlockAdapterImpl() *BlockAdapterImpl {
	return &BlockAdapterImpl{}
}

// NewBlockAdapter is used to create a new block adapter.
func NewBlockAdapter(impl *BlockAdapterImpl) repo.BlockAdapter {
	return impl
}

func (i *BlockAdapterImpl) ScanBlock(
	c context.Context,
	req repo.ScanBlockRequest,
	blockCh chan<- *blockB.Block,
) error {
	// TODO: 2024/12/1|sean|implement me
	panic("implement me")
}
