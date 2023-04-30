package biz

import (
	"github.com/blackhorseya/ryze/internal/app/domain/block/biz/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	bb "github.com/blackhorseya/ryze/pkg/entity/domain/block/biz"
	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// BlockSet is the provider set of biz
var BlockSet = wire.NewSet(NewImpl)

type impl struct {
	repo repo.IRepo
}

// NewImpl will create an object that implement IBiz interface
func NewImpl(repo repo.IRepo) bb.IBiz {
	return &impl{
		repo: repo,
	}
}

func (i *impl) GetBlockByHash(ctx contextx.Contextx, hash []byte) (record *bm.Block, err error) {
	// todo: 2023/4/30|sean|impl me
	panic("implement me")
}

func (i *impl) ListBlocks(ctx contextx.Contextx, condition bb.ListBlocksCondition) (records []*bm.Block, total int, err error) {
	// todo: 2023/4/30|sean|impl me
	panic("implement me")
}

func (i *impl) ListenNewBlock(ctx contextx.Contextx) (newBlockChan <-chan *bm.Block, err error) {
	blocks, err := i.repo.SubscribeNewBlock(ctx)
	if err != nil {
		ctx.Error("listen new block error", zap.Error(err))
		return nil, err
	}

	for {
		select {
		case <-ctx.Done():
			return
		case block := <-blocks:
			ctx.Info("get new block", zap.Any("block", block))

			// todo: 2023/4/30|sean|breakpoint: save block to db
		}
	}
}
