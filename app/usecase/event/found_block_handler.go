package event

import (
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/eventx"
	"go.uber.org/zap"
)

type foundBlockHandler struct {
}

func (h *foundBlockHandler) Handle(event eventx.DomainEvent) {
	ctx := contextx.Background()
	ctx.Info("called found block handler", zap.String("event", event.GetName()))
}

// NewFoundBlockHandler creates a new found block handler.
func NewFoundBlockHandler() eventx.EventHandler {
	return &foundBlockHandler{}
}
