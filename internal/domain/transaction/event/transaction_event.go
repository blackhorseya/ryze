package event

import "github.com/blackhorseya/ryze/internal/domain/event"

var _ event.Event = (*TransactionCreatedEvent)(nil)

// TransactionCreatedEvent 代表新交易產生的事件
// 可依實際需求擴充欄位
type TransactionCreatedEvent struct {
	TransactionID string
}

// Data implements event.Event.
func (t *TransactionCreatedEvent) Data() map[string]interface{} {
	// TODO: implement
	panic("unimplemented")
}

// Metadata implements event.Event.
func (t *TransactionCreatedEvent) Metadata() map[string]interface{} {
	// TODO: implement
	panic("unimplemented")
}

// Type implements event.Event.
func (t *TransactionCreatedEvent) Type() string {
	// TODO: implement
	panic("unimplemented")
}
