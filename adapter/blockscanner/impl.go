package blockscanner

import (
	"context"
	"errors"
	"io"

	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/internal/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.uber.org/zap"
)

type impl struct {
	injector *Injector
	server   *grpcx.Server
}

// NewServer is used to create a new scan server
func NewServer(injector *Injector, server *grpcx.Server) adapterx.Server {
	return &impl{
		injector: injector,
		server:   server,
	}
}

//nolint:funlen,gocognit // it's ok
func (i *impl) Start(c context.Context) error {
	ctx := contextx.WithContext(c)

	if err := i.server.Start(ctx); err != nil {
		ctx.Error("failed to start server", zap.Error(err))
		return err
	}

	blockScanner, err := i.injector.blockClient.ScanBlock(ctx, &biz.ScanBlockRequest{
		StartHeight: 0,
		EndHeight:   0,
	})
	if err != nil {
		ctx.Error("failed to scan block", zap.Error(err))
		return err
	}

	foundNewBlock, err := i.injector.blockClient.FoundNewBlock(ctx)
	if err != nil {
		ctx.Error("failed to found new block", zap.Error(err))
		return err
	}

	processBlock, err := i.injector.txClient.ProcessBlockTransactions(ctx)
	if err != nil {
		ctx.Error("failed to process block transactions", zap.Error(err))
		return err
	}

	go func() {
		ctx.Info("start to receive block")
		for {
			newBlockEvent, err2 := blockScanner.Recv()
			if errors.Is(err2, io.EOF) {
				break
			}
			if err2 != nil {
				ctx.Error("failed to receive block", zap.Error(err2))
				continue
			}
			ctx.Info("received block", zap.String("block_id", newBlockEvent.Id))

			err2 = foundNewBlock.Send(newBlockEvent)
			if err2 != nil {
				ctx.Error("failed to send block", zap.Error(err2))
				continue
			}
		}
	}()

	go func() {
		ctx.Info("start to receive new block")
		for {
			newBlock, err2 := foundNewBlock.Recv()
			if errors.Is(err2, io.EOF) {
				break
			}
			if err2 != nil {
				ctx.Error("failed to receive new block", zap.Error(err2))
				continue
			}
			ctx.Info("received new block", zap.String("block_id", newBlock.Id))

			err2 = processBlock.Send(newBlock)
			if err2 != nil {
				ctx.Error("failed to send new block", zap.Error(err2))
				continue
			}
		}
	}()

	go func() {
		ctx.Info("start to receive transaction")
		for {
			tx, err2 := processBlock.Recv()
			if errors.Is(err2, io.EOF) {
				break
			}
			if err2 != nil {
				ctx.Error("failed to receive transaction", zap.Error(err2))
				continue
			}
			ctx.Info("received transaction", zap.Any("transaction", &tx))
		}
	}()

	return nil
}

func (i *impl) Shutdown(c context.Context) error {
	ctx := contextx.WithContext(c)
	ctx.Info("server shutdown")

	if err := i.server.Stop(ctx); err != nil {
		ctx.Error("failed to stop server", zap.Error(err))
	}

	return nil
}
