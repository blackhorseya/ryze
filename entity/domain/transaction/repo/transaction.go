//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"context"

	"github.com/blackhorseya/ryze/entity/domain/transaction/model"
)

// ListTransactionsCondition is the condition for list.
type ListTransactionsCondition struct {
	Limit  int
	Offset int
}

// ITransactionRepo is the interface for transaction repository.
type ITransactionRepo interface {
	Create(c context.Context, item *model.Transaction) (err error)
	GetByID(c context.Context, id string) (item *model.Transaction, err error)
	List(c context.Context, cond ListTransactionsCondition) (items []*model.Transaction, total int, err error)
	Update(c context.Context, item *model.Transaction) (err error)
	Delete(c context.Context, id string) (err error)

	ListByAccount(
		c context.Context,
		accountID string,
		cond ListTransactionsCondition,
	) (items []*model.Transaction, total int, err error)
}
