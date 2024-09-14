package transaction

import (
	"github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	"github.com/blackhorseya/ryze/internal/infra/transports/grpcx"
)

// NewTransactionServiceClient is used to create a new transaction service client.
func NewTransactionServiceClient(client *grpcx.Client) (biz.TransactionServiceClient, error) {
	conn, err := client.Dial("block-scanner")
	if err != nil {
		return nil, err
	}

	return biz.NewTransactionServiceClient(conn), nil
}
