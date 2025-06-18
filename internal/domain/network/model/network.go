package model

import "time"

// NetworkStats 代表 TON 區塊鏈網路統計資料（Value Object）
type NetworkStats struct {
	BlockHeight      uint64    // 區塊高度
	TPS              float64   // 每秒交易數
	ActiveValidators uint64    // 活躍驗證者數
	TotalAccounts    uint64    // 帳戶總數
	UpdatedAt        time.Time // 統計資料更新時間
}

// NodeStatus 代表網路節點狀態（Value Object）
type NodeStatus struct {
	ID          string    // 節點 ID
	IsActive    bool      // 是否活躍
	Version     string    // 節點版本
	LastSeen    time.Time // 最後上線時間
	BlockHeight uint64    // 節點區塊高度
}
