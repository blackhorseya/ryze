package event

import (
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

// NewFoundBlockHandler creates a new found block handler.
func NewFoundBlockHandler(
	blockClient blockB.BlockServiceClient,
	txClient txB.TransactionServiceClient,
) eventx.EventHandler {
	return &foundBlockHandler{
		blockClient: blockClient,
		txClient:    txClient,
	}
}

func (h *foundBlockHandler) Handle(event eventx.DomainEvent) {
	ctx := contextx.Background()

	blockEvent, ok := event.(*model.FoundBlockEvent)
	if !ok {
		ctx.Error("failed to cast event to FoundBlockEvent")
		return
	}

	ctx.Info("found block", zap.String("block", blockEvent.Block.String()))

	// TODO: 2024/9/15|sean|call block service to handle the block via block client

	// TODO: 2024/9/15|sean|call transaction service to handle the block via transaction client
}
