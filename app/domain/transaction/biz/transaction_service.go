package biz

import (
	"context"

	"github.com/blackhorseya/ryze/app/infra/tonx"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	txM "github.com/blackhorseya/ryze/entity/domain/transaction/model"
)

type txService struct {
	client *tonx.Client
}

// NewTransactionService is used to create a new transaction service
func NewTransactionService(client *tonx.Client) txB.TransactionServiceServer {
	return &txService{
		client: client,
	}
}

func (i *txService) GetTransaction(c context.Context, req *txB.GetTransactionRequest) (*txM.Transaction, error) {
	// TODO: 2024/8/12|sean|implement me
	panic("implement me")
}

func (i *txService) ListTransactions(
	req *txB.ListTransactionsRequest,
	stream txB.TransactionService_ListTransactionsServer,
) error {
	// TODO: 2024/8/12|sean|implement me
	panic("implement me")
}
