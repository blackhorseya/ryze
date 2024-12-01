package ton

import (
	"context"

	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/internal/app/repo"
	"github.com/blackhorseya/ryze/internal/shared/tonx"
)

// BlockAdapterImpl is the implementation for block adapter.
type BlockAdapterImpl struct {
	client *tonx.Client
}

// NewBlockAdapterImpl is used to create a new block adapter implementation.
func NewBlockAdapterImpl(client *tonx.Client) *BlockAdapterImpl {
	return &BlockAdapterImpl{
		client: client,
	}
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
