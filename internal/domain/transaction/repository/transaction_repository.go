//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repository

import (
	"context"

	"github.com/blackhorseya/ryze/internal/domain/transaction/model"
)

// TransactionRepository 定義交易資料存取契約（DDD Repository interface）
type TransactionRepository interface {
	FindByID(c context.Context, id string) (*model.Transaction, error)
	FindByBlockID(c context.Context, blockID string) ([]*model.Transaction, error)
	// 其他查詢/儲存方法
}
