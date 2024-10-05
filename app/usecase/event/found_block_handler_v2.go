package event

import (
	"context"

	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
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
	// TODO: 2024/10/5|sean|implement me
	ctx, span := contextx.StartSpan(context.Background(), "event.found_block_handler_v2.Handle")
	defer span.End()

	ctx.Info("found block event", zap.Any("event", &event))
}
