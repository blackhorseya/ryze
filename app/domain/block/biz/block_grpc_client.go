package biz

import (
	"fmt"

	"github.com/blackhorseya/ryze/app/infra/transports/grpcx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
)

// NewBlockServiceClient is used to create a new block service client
func NewBlockServiceClient(client *grpcx.Client) (model.BlockServiceClient, error) {
	conn, err := client.Dial("block-grpc")
	if err != nil {
		return nil, fmt.Errorf("failed to dial `block-grpc` error: %w", err)
	}

	return model.NewBlockServiceClient(conn), nil
}
