package event

import "github.com/blackhorseya/ryze/internal/domain/event"

var _ event.Event = (*BlockCreatedEvent)(nil)

// BlockCreatedEvent 代表新區塊產生的事件
// 可依實際需求擴充欄位
// Example: event for event sourcing or domain notification
type BlockCreatedEvent struct {
	BlockID string
}

// Data implements event.Event.
func (b *BlockCreatedEvent) Data() map[string]interface{} {
	// TODO: implement
	panic("unimplemented")
}

// Metadata implements event.Event.
func (b *BlockCreatedEvent) Metadata() map[string]interface{} {
	// TODO: implement
	panic("unimplemented")
}

// Type implements event.Event.
func (b *BlockCreatedEvent) Type() string {
	// TODO: implement
	panic("unimplemented")
}
