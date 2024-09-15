package event

import (
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/eventx"
	"go.uber.org/zap"
)

type foundBlockHandler struct {
}

// NewFoundBlockHandler creates a new found block handler.
func NewFoundBlockHandler() eventx.EventHandler {
	return &foundBlockHandler{}
}

func (h *foundBlockHandler) Handle(event eventx.DomainEvent) {
	ctx := contextx.Background()

	blockEvent, ok := event.(*model.FoundBlockEvent)
	if !ok {
		ctx.Error("failed to cast event to FoundBlockEvent")
		return
	}

	ctx.Info("found block", zap.String("block", blockEvent.Block.String()))
	// TODO: 2024/9/15|sean|do something with the block
}
