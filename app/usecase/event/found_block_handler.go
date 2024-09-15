package event

import (
	"context"
	"errors"
	"io"

	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/eventx"
	"go.uber.org/zap"
)

type foundBlockHandler struct {
	blockClient blockB.BlockServiceClient
	txClient    txB.TransactionServiceClient
}

// NewFoundBlockHandler creates a new-found block handler.
func NewFoundBlockHandler(
	blockClient blockB.BlockServiceClient,
	txClient txB.TransactionServiceClient,
) eventx.EventHandler {
	return &foundBlockHandler{
		blockClient: blockClient,
		txClient:    txClient,
	}
}

func (i *foundBlockHandler) Handle(event eventx.DomainEvent) {
	ctx := contextx.Background()

	blockEvent, ok := event.(*model.FoundBlockEvent)
	if !ok {
		ctx.Error("failed to cast event to FoundBlockEvent")
		return
	}

	block := blockEvent.Block

	// Call block service to handle the block via block client
	blockStream, err := i.blockClient.FoundNewBlock(ctx)
	if err != nil {
		ctx.Error("failed to call block service", zap.Error(err))
		return
	}
	err = blockStream.Send(block)
	if err != nil {
		ctx.Error("failed to send block to block service", zap.Error(err))
		return
	}
	_ = blockStream.CloseSend()

	// Wait for block service to receive the block
	block, err = blockStream.Recv()
	if err != nil {
		ctx.Error("failed to receive response from block service", zap.Error(err))
		return
	}

	ctx.Info("found block", zap.String("block", block.String()))

	// Call transaction service to handle the block via transaction client
	transactions, err := i.txClient.ProcessBlockTransactions(ctx)
	if err != nil {
		ctx.Error("failed to call transaction service", zap.Error(err))
		return
	}
	err = transactions.Send(block)
	if err != nil {
		ctx.Error("failed to send block to transaction service", zap.Error(err))
		return
	}
	_ = transactions.CloseSend()

	// Wait for transaction service to receive the block
	for {
		transaction, err2 := transactions.Recv()
		if errors.Is(err2, io.EOF) || errors.Is(err2, context.Canceled) {
			break
		}
		if err2 != nil {
			ctx.Error("failed to receive response from transaction service", zap.Error(err2))
			return
		}

		ctx.Info("found transaction", zap.String("transaction", transaction.String()))
	}
}
