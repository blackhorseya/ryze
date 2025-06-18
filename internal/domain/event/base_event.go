package event

import "time"

// BaseEvent 提供所有領域事件的基礎實作
type BaseEvent struct {
	eventType string
	data      map[string]interface{}
	metadata  map[string]interface{}
}

// NewBaseEvent 建立一個基礎事件
func NewBaseEvent(eventType string, data map[string]interface{}) *BaseEvent {
	return &BaseEvent{
		eventType: eventType,
		data:      data,
		metadata: map[string]interface{}{
			"occurred_at": time.Now().Format(time.RFC3339),
		},
	}
}

// Type 返回事件類型
func (e *BaseEvent) Type() string {
	return e.eventType
}

// Data 返回事件資料
func (e *BaseEvent) Data() map[string]interface{} {
	return e.data
}

// Metadata 返回事件中繼資料
func (e *BaseEvent) Metadata() map[string]interface{} {
	return e.metadata
}

// WithMetadata 設定額外中繼資料並返回事件本身
func (e *BaseEvent) WithMetadata(key string, value interface{}) *BaseEvent {
	e.metadata[key] = value
	return e
}
