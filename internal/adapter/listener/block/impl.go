package block

import (
	"context"

	"github.com/blackhorseya/ryze/pkg/adapter"
	"github.com/blackhorseya/ryze/pkg/contextx"
	bb "github.com/blackhorseya/ryze/pkg/entity/domain/block/biz"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type impl struct {
	ctx    contextx.Contextx
	cancel context.CancelFunc
	logger *zap.Logger
	biz    bb.IBiz
}

// NewImpl returns a new block listener implementation.
func NewImpl(logger *zap.Logger, biz bb.IBiz) adapter.Listener {
	ctx, cancel := contextx.WithCancel(contextx.BackgroundWithLogger(logger))

	return &impl{
		ctx:    ctx,
		cancel: cancel,
		logger: logger,
		biz:    biz,
	}
}

func (i *impl) Start() error {
	i.logger.Info("start block listener")

	go i.biz.ListenNewBlock(i.ctx)

	return nil
}

func (i *impl) Stop() error {
	i.logger.Info("stop block listener")

	i.cancel()

	return nil
}

// ListenerSet presents a listener set.
var ListenerSet = wire.NewSet(NewImpl)
