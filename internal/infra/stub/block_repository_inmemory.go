package stub

import (
	"context"
	"sync"

	"github.com/blackhorseya/ryze/internal/domain/block/model"
	"github.com/blackhorseya/ryze/internal/domain/block/repository"
)

// InMemoryBlockRepository 為簡易記憶體實作，僅供測試/開發期 wire 注入使用。
// 不建議於生產使用。

type InMemoryBlockRepository struct {
	mu    sync.RWMutex
	store []*model.Block
}

// NewInMemoryBlockRepository 建立新的記憶體實作。
func NewInMemoryBlockRepository() repository.BlockRepository {
	return &InMemoryBlockRepository{
		store: make([]*model.Block, 0),
	}
}

// FindByID 依 id 取得區塊
func (r *InMemoryBlockRepository) FindByID(_ context.Context, id string) (*model.Block, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, b := range r.store {
		if b.ID == id {
			return b, nil
		}
	}
	return nil, nil // not found
}

// FindLatest 取得最新 limit 筆區塊
func (r *InMemoryBlockRepository) FindLatest(_ context.Context, limit int) ([]*model.Block, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if limit <= 0 || limit > len(r.store) {
		limit = len(r.store)
	}
	// 簡單 slice copy
	latest := make([]*model.Block, limit)
	copy(latest, r.store[len(r.store)-limit:])
	return latest, nil
}
