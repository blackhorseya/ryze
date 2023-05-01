package biz

import (
	"github.com/blackhorseya/ryze/internal/app/domain/block/biz/repo"
	"github.com/blackhorseya/ryze/internal/pkg/errorx"
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

func (i *impl) ListBlocks(ctx contextx.Contextx, condition bb.ListBlocksCondition) (records []*bm.Block, total uint, err error) {
	if condition.Page == 0 {
		ctx.Error(errorx.ErrInvalidPage.Error(), zap.Uint("page", condition.Page))
		return nil, 0, errorx.ErrInvalidPage
	}

	if condition.Size == 0 {
		ctx.Error(errorx.ErrInvalidSize.Error(), zap.Uint("size", condition.Size))
		return nil, 0, errorx.ErrInvalidSize
	}

	repoCondition := newRepoCondition(condition)
	blocks, total, err := i.repo.ListBlocks(ctx, repoCondition)
	if err != nil {
		ctx.Error(errorx.ErrListBlocks.Error(), zap.Error(err), zap.Any("condition", repoCondition))
		return nil, 0, errorx.ErrListBlocks
	}

	return blocks, total, nil
}

func (i *impl) ListenNewBlock(ctx contextx.Contextx) error {
	blocks, err := i.repo.SubscribeNewBlock(ctx)
	if err != nil {
		ctx.Error("listen new block error", zap.Error(err))
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case block := <-blocks:
			ctx.Info("get new block", zap.Any("block", block))

			err = i.repo.CreateNewBlock(ctx, block)
			if err != nil {
				ctx.Error("create new block error", zap.Error(err))
				continue
			}
		}
	}
}

func newRepoCondition(condition bb.ListBlocksCondition) repo.ListBlocksCondition {
	return repo.ListBlocksCondition{
		Limit:  condition.Size,
		Offset: (condition.Page - 1) * condition.Size,
	}
}
