package event

import (
	"context"
	"time"
)

// DomainEvent 定義領域事件介面，用於標識事件類型與發生時間
type DomainEvent interface {
	EventType() string     // 事件類型識別
	OccurredAt() time.Time // 事件發生時間
}

// EventBus 定義事件總線介面，用於發佈領域事件
type EventBus interface {
	// Publish 發佈領域事件，若失敗返回錯誤
	Publish(ctx context.Context, evt DomainEvent) error
}
