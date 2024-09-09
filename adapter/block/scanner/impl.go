package scanner

import (
	"context"

	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
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

	// stream, err := i.injector.blockClient.ScanBlock(ctx, &biz.ScanBlockRequest{
	// 	StartHeight: 0,
	// 	EndHeight:   0,
	// })
	// if err != nil {
	// 	ctx.Error("failed to scan block", zap.Error(err))
	// 	return err
	// }
	//
	// for {
	// 	block, err := stream.Recv()
	// 	if err != nil && errors.Is(err, io.EOF) {
	// 		ctx.Error("failed to receive block", zap.Error(err))
	// 		return err
	// 	}
	//
	// 	ctx.Info("received block", zap.Any("block", &block))
	// }

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
