package block

import (
	"fmt"

	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/internal/app/infra/transports/grpcx"
)

// NewBlockServiceClient is used to create a new block service tonClient
func NewBlockServiceClient(client *grpcx.Client) (biz.BlockServiceClient, error) {
	conn, err := client.Dial("daemon")
	if err != nil {
		return nil, fmt.Errorf("failed to dial `daemon` error: %w", err)
	}

	return biz.NewBlockServiceClient(conn), nil
}
