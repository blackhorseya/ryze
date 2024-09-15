package transaction

import (
	"fmt"

	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/entity/domain/transaction/biz"
)

// NewTransactionServiceClient is used to create a new transaction service client.
func NewTransactionServiceClient(client *grpcx.Client) (biz.TransactionServiceClient, error) {
	conn, err := client.Dial("daemon")
	if err != nil {
		return nil, fmt.Errorf("failed to dial `daemon` error: %w", err)
	}

	return biz.NewTransactionServiceClient(conn), nil
}
