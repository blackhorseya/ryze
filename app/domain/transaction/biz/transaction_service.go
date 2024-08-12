package biz

import (
	"context"

	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	txM "github.com/blackhorseya/ryze/entity/domain/transaction/model"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.uber.org/zap"
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
	ctx, err := contextx.FromContext(stream.Context())
	if err != nil {
		return err
	}

	ctx, span := otelx.Span(ctx, "transaction.biz.ListTransactions")
	defer span.End()

	var txList []*txM.Transaction

	// TODO: 2024/8/12|sean|implement me

	for _, tx := range txList {
		if err = stream.Send(tx); err != nil {
			ctx.Error("send transaction error", zap.Error(err))
			return err
		}
	}

	return nil
}
