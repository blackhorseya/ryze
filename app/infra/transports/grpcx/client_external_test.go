//go:build external

package grpcx

import (
	"testing"

	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/spf13/viper"
)

func TestNewClient(t *testing.T) {
	config, err := configx.NewConfiguration(viper.GetViper())
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	client, err := NewClient(config)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	conn, err := client.Dial("block-grpc")
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	blockService := model.NewBlockServiceClient(conn)
	block, err := blockService.GetBlock(contextx.Background(), &model.GetBlockRequest{
		Workchain: -1,
		Shard:     8000000000000000,
		SeqNo:     39346131,
	})
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	t.Logf("Block: %v", block)
}
