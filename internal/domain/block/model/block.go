package model

import (
	"fmt"
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
	Workchain    int32     // 工作鏈 ID
	Shard        int64     // 分片 ID
	SeqNo        uint32    // 區塊序號（在分片中的順序）

	// 其他區塊屬性
}

// NewBlock is used to create a new block.
func NewBlock(workchain int32, shard int64, seqno uint32) (*Block, error) {
	return &Block{
		ID:        fmt.Sprintf("%d:%d:%d", workchain, shard, seqno),
		Workchain: workchain,
		Shard:     shard,
		SeqNo:     seqno,
	}, nil
}
