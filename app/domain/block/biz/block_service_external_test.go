//go:build external

package biz

import (
	"testing"

	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteExternal struct {
	suite.Suite

	biz model.BlockServiceServer
}

func (s *suiteExternal) SetupTest() {
	client, err := tonx.NewClient(tonx.Options{Network: "mainnet"})
	s.Require().NoError(err)

	s.biz = NewBlockService(client)
}

func (s *suiteExternal) TearDownTest() {
}

func TestExternalAll(t *testing.T) {
	suite.Run(t, new(suiteExternal))
}

func (s *suiteExternal) Test_impl_GetBlock() {
	ctx := contextx.Background()
	block, err := s.biz.GetBlock(ctx, &model.GetBlockRequest{
		Workchain: -1,
		Shard:     8000000000000000,
		SeqNo:     39382597,
	})
	s.Require().NoError(err)

	ctx.Debug("get block", zap.Any("block", &block))
}

func (s *suiteExternal) Test_impl_FetchAndStoreBlock() {
	ctx := contextx.Background()
	block, err := s.biz.FetchAndStoreBlock(ctx, &model.FetchAndStoreBlockRequest{
		Workchain: -1,
		Shard:     8000000000000000,
		SeqNo:     39382597,
	})
	s.Require().NoError(err)

	ctx.Debug("fetch and store block", zap.Any("block", &block))
}
