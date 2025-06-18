package block

import (
	"context"

	"github.com/blackhorseya/ryze/internal/domain/block/model"
)

// LoggingMiddleware 提供日誌功能
// TODO: 實作日誌 middleware
type LoggingMiddleware struct {
	next Service
}

// NewLoggingMiddleware 建立日誌 middleware
func NewLoggingMiddleware(next Service) Service {
	// TODO: 回傳 LoggingMiddleware 實例
	return nil
}

func (mw *LoggingMiddleware) GetBlockByID(c context.Context, id string) (*model.Block, error) {
	// TODO: 實作日誌邏輯
	return nil, nil
}

func (mw *LoggingMiddleware) GetLatestBlocks(c context.Context, limit int) ([]*model.Block, error) {
	// TODO: 實作日誌邏輯
	return nil, nil
}
