package pgx

import (
	"context"
	"fmt"

	"github.com/blackhorseya/ryze/entity/domain/transaction/model"
	"github.com/blackhorseya/ryze/internal/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.uber.org/zap"
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
	ctx, span := contextx.StartSpan(c, "storage.pgx.transaction_repo.Create")
	defer span.End()

	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	err = i.rw.WithContext(timeout).Create(item).Error
	if err != nil {
		ctx.Error("create transaction to gormDB failed", zap.Error(err), zap.Any("transaction", &item))
		span.RecordError(err)
		return err
	}

	return nil
}

func (i *transactionRepo) GetByID(c context.Context, id string) (item *model.Transaction, err error) {
	ctx, span := contextx.StartSpan(c, "storage.pgx.transaction_repo.GetByID")
	defer span.End()

	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	err = i.rw.WithContext(timeout).Where("id = ?", id).First(&item).Error
	if err != nil {
		ctx.Error("get transaction by id from gormDB failed", zap.Error(err), zap.String("id", id))
		span.RecordError(err)
		return nil, fmt.Errorf("get transaction by id from gormDB failed: %w", err)
	}

	return item, nil
}

func (i *transactionRepo) List(
	c context.Context,
	cond repo.ListTransactionsCondition,
) (items []*model.Transaction, total int, err error) {
	ctx, span := contextx.StartSpan(c, "storage.pgx.transaction_repo.List")
	defer span.End()

	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	query := i.rw.WithContext(timeout).Model(&model.Transaction{})

	// limit and offset
	limit, offset := defaultLimit, 0
	if 0 < cond.Limit && cond.Limit <= defaultMaxLimit {
		limit = cond.Limit
	}
	if 0 < cond.Offset {
		offset = cond.Offset
	}
	query = query.Limit(limit).Offset(offset)

	// order by
	query = query.Order("timestamp desc")

	var count int64
	err = query.Count(&count).Find(&items).Error
	if err != nil {
		ctx.Error("list transaction from gormDB failed", zap.Error(err))
		span.RecordError(err)
		return nil, 0, err
	}

	return items, int(count), nil
}

func (i *transactionRepo) ListByAccount(
	c context.Context,
	accountID string,
	cond repo.ListTransactionsCondition,
) (items []*model.Transaction, total int, err error) {
	ctx, span := contextx.StartSpan(c, "storage.pgx.transaction_repo.ListByAccount")
	defer span.End()

	timeout, cancelFunc := context.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	query := i.rw.WithContext(timeout).
		Model(&model.Transaction{}).
		Where(`"from"::bytea = decode(?, 'base64') OR "to"::bytea = decode(?, 'base64')`, accountID, accountID)

	// limit and offset
	limit, offset := defaultLimit, 0
	if 0 < cond.Limit && cond.Limit <= defaultMaxLimit {
		limit = cond.Limit
	}
	if 0 < cond.Offset {
		offset = cond.Offset
	}
	query = query.Limit(limit).Offset(offset)

	// order by
	query = query.Order("timestamp desc")

	var count int64
	err = query.Count(&count).Find(&items).Error
	if err != nil {
		ctx.Error("list transactions by account from gormDB failed", zap.Error(err), zap.String("account_id", accountID))
		span.RecordError(err)
		return nil, 0, fmt.Errorf("list transactions by account from gormDB failed: %w", err)
	}

	return items, int(count), nil
}
