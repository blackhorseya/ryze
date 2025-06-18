package model

import "time"

// Account 代表 TON 區塊鏈中的帳戶（聚合根）
type Account struct {
	Address    string    // 帳戶地址
	Balance    uint64    // 帳戶餘額
	LastActive time.Time // 最後活躍時間
	// 其他帳戶屬性
}
