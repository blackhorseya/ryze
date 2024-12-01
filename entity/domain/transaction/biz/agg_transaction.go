package biz

import (
	"github.com/blackhorseya/ryze/entity/domain/transaction/model"
)

// Transaction is the aggregate root of transaction
type Transaction struct {
	model.Transaction
}
