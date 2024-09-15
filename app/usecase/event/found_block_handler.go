package event

import (
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
	ctx.Info("called found block handler", zap.String("event", event.GetName()))
}
