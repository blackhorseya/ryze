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

type foundBlockHandlerV2 struct {
	blockClient blockB.BlockServiceClient
	txClient    txB.TransactionServiceClient
}

// NewFoundBlockHandlerV2 creates a new-found block handler.
func NewFoundBlockHandlerV2(
	blockClient blockB.BlockServiceClient,
	txClient txB.TransactionServiceClient,
) eventx.EventHandler {
	return &foundBlockHandlerV2{
		blockClient: blockClient,
		txClient:    txClient,
	}
}

func (i *foundBlockHandlerV2) Handle(event eventx.DomainEvent) {
	ctx, span := contextx.StartSpan(context.Background(), "event.found_block_handler_v2.Handle")
	defer span.End()

	blockEvent, ok := event.(*model.FoundBlockEvent)
	if !ok {
		ctx.Error("failed to cast event to FoundBlockEvent", zap.Any("event", event))
		return
	}

	ctx.Info("start handle new block event", zap.String("block_id", blockEvent.Block.Id))

	// handle new block by id
	newBlock, err := i.blockClient.FoundNewBlockNonStream(ctx, &blockB.FoundNewBlockRequest{
		Workchain: blockEvent.Block.Workchain,
		Shard:     blockEvent.Block.Shard,
		SeqNo:     blockEvent.Block.SeqNo,
	})
	if err != nil {
		ctx.Error("failed to handle new block", zap.Error(err))
		return
	}
	ctx.Debug(
		"found new block",
		zap.String("block_id", newBlock.Id),
		zap.Time("block_time", newBlock.Timestamp.AsTime()),
	)

	// handle new block transactions
	stream, err := i.txClient.ProcessBlockTransactionsNonStream(ctx, newBlock)
	if err != nil {
		ctx.Error("failed to handle new block transactions", zap.Error(err))
		return
	}

	for {
		tx, err2 := stream.Recv()
		if errors.Is(err2, io.EOF) {
			break
		}
		if err2 != nil {
			ctx.Error("failed to receive new block transaction", zap.Error(err2))
			break
		}
		ctx.Debug("found new block transaction", zap.String("tx_id", string(tx.Id)))
	}
}
