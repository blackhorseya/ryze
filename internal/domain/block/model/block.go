package model

import (
	"time"
)

// Block 代表 TON 區塊鏈中的區塊（聚合根）
type Block struct {
	ID           string    // 區塊唯一識別碼
	Height       uint64    // 區塊高度
	Hash         string    // 區塊雜湊值
	PrevHash     string    // 前一區塊雜湊值
	CreatedAt    time.Time // 區塊建立時間
	Transactions []string  // 交易 ID 列表（如需可改為 Transaction Value Object）
	// 其他區塊屬性
}
