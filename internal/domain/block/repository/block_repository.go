//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repository

import (
	"context"

	"github.com/blackhorseya/ryze/internal/domain/block/model"
)

// BlockRepository 定義區塊資料存取契約（DDD Repository interface）
type BlockRepository interface {
	FindByID(c context.Context, id string) (*model.Block, error)
	FindLatest(c context.Context, limit int) ([]*model.Block, error)
	// 其他查詢/儲存方法
}
