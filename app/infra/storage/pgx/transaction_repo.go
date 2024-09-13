package pgx

import (
	"context"
	"fmt"

	"github.com/blackhorseya/ryze/entity/domain/transaction/model"
	"github.com/blackhorseya/ryze/entity/domain/transaction/repo"
	"gorm.io/gorm"
)

type transactionRepo struct {
	rw *gorm.DB
}

// NewTransactionRepo create and return a new transactionRepo.
func NewTransactionRepo(rw *gorm.DB) (repo.ITransactionRepo, error) {
	err := rw.AutoMigrate(&model.Transaction{})
	if err != nil {
		return nil, fmt.Errorf("auto migrate transaction error: %w", err)
	}

	return &transactionRepo{
		rw: rw,
	}, nil
}

func (i *transactionRepo) Create(c context.Context, item *model.Transaction) (err error) {
	// TODO: 2024/9/13|sean|implement me
	panic("implement me")
}

func (i *transactionRepo) GetByID(c context.Context, id string) (item *model.Transaction, err error) {
	// TODO: 2024/9/13|sean|implement me
	panic("implement me")
}

func (i *transactionRepo) List(
	c context.Context,
	cond repo.ListTransactionsCondition,
) (items []*model.Transaction, total int, err error) {
	// TODO: 2024/9/13|sean|implement me
	panic("implement me")
}

func (i *transactionRepo) Update(c context.Context, item *model.Transaction) (err error) {
	// TODO: 2024/9/13|sean|implement me
	panic("implement me")
}

func (i *transactionRepo) Delete(c context.Context, id string) (err error) {
	// TODO: 2024/9/13|sean|implement me
	panic("implement me")
}
