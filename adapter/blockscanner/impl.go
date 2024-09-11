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

	stream, err := i.injector.blockClient.ScanBlock(ctx, &biz.ScanBlockRequest{
		StartHeight: 0,
		EndHeight:   0,
	})
	if err != nil {
		ctx.Error("failed to scan block", zap.Error(err))
		return err
	}

	go func() {
		for {
			block, err2 := stream.Recv()
			if errors.Is(err2, io.EOF) {
				break
			}
			if err2 != nil {
				ctx.Error("failed to receive block", zap.Error(err2))
				return
			}

			ctx.Info("received block", zap.Any("block", &block))
			_, err2 = i.injector.blockClient.FoundNewBlock(ctx, &biz.FoundNewBlockRequest{
				Workchain: block.Workchain,
				Shard:     block.Shard,
				SeqNo:     block.SeqNo,
			})
			if err2 != nil {
				ctx.Error("failed to found new block", zap.Error(err2))
				return
			}
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
