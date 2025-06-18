package model

import "time"

// Transaction 代表 TON 區塊鏈中的交易（聚合根）
type Transaction struct {
	ID        string            // 交易唯一識別碼
	BlockID   string            // 所屬區塊 ID
	FromAddr  string            // 來源地址
	ToAddr    string            // 目標地址
	Amount    uint64            // 交易金額
	Fee       uint64            // 交易手續費
	Status    TransactionStatus // 交易狀態
	Timestamp time.Time         // 交易時間
	// 其他交易屬性
}

// TransactionStatus 代表交易狀態（Value Object）
type TransactionStatus string

const (
	StatusPending   TransactionStatus = "pending"
	StatusCompleted TransactionStatus = "completed"
	StatusFailed    TransactionStatus = "failed"
)
