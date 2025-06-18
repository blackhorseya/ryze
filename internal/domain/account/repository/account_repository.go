//go:generate mockgen -destination=./mock_account_repository.go -package=repository -source=account_repository.go
package repository

import (
	"context"

	"github.com/blackhorseya/ryze/internal/domain/account/model"
	txM "github.com/blackhorseya/ryze/internal/domain/transaction/model"
)

// AccountRepository 定義帳戶資料存取契約（DDD Repository interface）
type AccountRepository interface {
	FindByAddress(c context.Context, address string) (*model.Account, error)
	GetBalance(c context.Context, address string) (uint64, error)
	GetTransactionHistory(c context.Context, address string, limit int) ([]*txM.Transaction, error)
	// 其他查詢/儲存方法
}
