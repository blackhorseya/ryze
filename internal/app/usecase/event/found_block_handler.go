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
	blockClient  blockB.BlockServiceClient
	txClient     txB.TransactionServiceClient
	blockStream  blockB.BlockService_FoundNewBlockClient               // 儲存 block 連線
	transactions txB.TransactionService_ProcessBlockTransactionsClient // 儲存 transaction 連線
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

// setupConnections 初始化連線
func (i *foundBlockHandler) setupConnections(ctx context.Context) error {
	var err error

	if i.blockStream == nil {
		i.blockStream, err = i.blockClient.FoundNewBlock(ctx)
		if err != nil {
			return err
		}
	}

	if i.transactions == nil {
		i.transactions, err = i.txClient.ProcessBlockTransactions(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *foundBlockHandler) Handle(event eventx.DomainEvent) {
	ctx := contextx.WithContext(context.Background())

	blockEvent, ok := event.(*model.FoundBlockEvent)
	if !ok {
		ctx.Error("failed to cast event to FoundBlockEvent")
		return
	}
	block := blockEvent.Block

	// 初始化連線
	if err := i.setupConnections(ctx); err != nil {
		ctx.Error("failed to setup connections", zap.Error(err))
		return
	}

	// Send block to block service
	err := i.blockStream.Send(block)
	if err != nil {
		ctx.Error("failed to send block to block service", zap.Error(err))
		return
	}
	block, err = i.blockStream.Recv()
	if err != nil {
		ctx.Error("failed to receive response from block service", zap.Error(err))
		return
	}
	ctx.Info("found block", zap.String("block", block.String()))

	// Send block to transaction service
	err = i.transactions.Send(block)
	if err != nil {
		ctx.Error("failed to send block to transaction service", zap.Error(err))
		return
	}

	for {
		transaction, err2 := i.transactions.Recv()
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
