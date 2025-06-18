package event

import "context"

// Event 定義領域事件的基本介面
type Event interface {
	// Type 返回事件類型
	Type() string

	// Data 返回事件的資料
	Data() map[string]interface{}

	// Metadata 返回事件的中繼資料
	Metadata() map[string]interface{}
}

// EventCollector 定義事件收集器介面
type EventCollector interface {
	// Collect 收集領域事件
	Collect(c context.Context, e Event) error
}
