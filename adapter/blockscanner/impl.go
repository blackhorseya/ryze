package blockscanner

import (
	"context"
	"errors"
	"io"

	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/entity/domain/block/biz"
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

func (i *impl) Start(c context.Context) error {
	ctx := contextx.WithContext(c)

	if err := i.server.Start(ctx); err != nil {
		ctx.Error("failed to start server", zap.Error(err))
		return err
	}

	blockStream, err := i.injector.blockClient.ScanBlock(ctx, &biz.ScanBlockRequest{
		StartHeight: 0,
		EndHeight:   0,
	})
	if err != nil {
		ctx.Error("failed to scan block", zap.Error(err))
		return err
	}

	txStream, err := i.injector.txClient.ProcessBlockTransactions(ctx)
	if err != nil {
		ctx.Error("failed to process block transactions", zap.Error(err))
		return err
	}

	go func() {
		ctx.Info("start to receive block")
		for {
			block, err2 := blockStream.Recv()
			if errors.Is(err2, io.EOF) {
				break
			}
			if err2 != nil {
				ctx.Error("failed to receive block", zap.Error(err2))
				return
			}

			ctx.Info("received block", zap.Any("block", &block))
			block, err2 = i.injector.blockClient.FoundNewBlock(ctx, &biz.FoundNewBlockRequest{
				Workchain: block.Workchain,
				Shard:     block.Shard,
				SeqNo:     block.SeqNo,
			})
			if err2 != nil {
				ctx.Error("failed to found new block", zap.Error(err2))
				return
			}
			ctx.Info("found new block", zap.Any("block", &block))

			err2 = txStream.Send(block)
			if err2 != nil {
				ctx.Error("failed to send block", zap.Error(err2))
				return
			}
		}
	}()

	go func() {
		for {
			tx, err2 := txStream.Recv()
			if errors.Is(err2, io.EOF) {
				break
			}
			if err2 != nil {
				ctx.Error("failed to receive transaction", zap.Error(err2))
				return
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
