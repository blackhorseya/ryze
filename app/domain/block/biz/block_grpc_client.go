package biz

import (
	"fmt"

	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/entity/domain/block/biz"
)

// NewBlockServiceClient is used to create a new block service client
func NewBlockServiceClient(client *grpcx.Client) (biz.BlockServiceClient, error) {
	conn, err := client.Dial("block-grpc")
	if err != nil {
		return nil, fmt.Errorf("failed to dial `block-grpc` error: %w", err)
	}

	return biz.NewBlockServiceClient(conn), nil
}
