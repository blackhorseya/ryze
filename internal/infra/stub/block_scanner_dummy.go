package stub

import (
	"context"

	"github.com/blackhorseya/ryze/internal/domain/block/model"
	"github.com/blackhorseya/ryze/internal/domain/block/service"
)

// DummyBlockScanner 為僅用於開發期的空實作，直接回傳 nil。

type DummyBlockScanner struct{}

// NewDummyBlockScanner 建立 DummyBlockScanner 實例
func NewDummyBlockScanner() service.BlockScanner {
	return &DummyBlockScanner{}
}

// ScanBlocks 不做任何事，直接結束。
func (d *DummyBlockScanner) ScanBlocks(_ context.Context, _ chan<- *model.Block, _ ...service.ScanBlockOption) error {
	return nil
}
