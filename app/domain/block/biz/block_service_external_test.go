//go:build external

package biz

import (
	"testing"

	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type suiteExternal struct {
	suite.Suite

	rw  *mongo.Client
	biz biz.BlockServiceServer
}

func (s *suiteExternal) SetupTest() {
	config, err := configx.NewConfiguration(viper.GetViper())
	s.Require().NoError(err)

	app, ok := config.Services["block-grpc"]
	s.Require().True(ok)

	client, err := tonx.NewClient(tonx.Options{Network: "mainnet"})
	s.Require().NoError(err)

	rw, err := mongodbx.NewClient(app)
	s.Require().NoError(err)
	s.rw = rw

	s.biz = NewExternalBlockService(client, s.rw)
}

func (s *suiteExternal) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}
}

func TestExternalAll(t *testing.T) {
	suite.Run(t, new(suiteExternal))
}

func (s *suiteExternal) Test_impl_GetBlock() {
	ctx := contextx.Background()
	block, err := s.biz.GetBlock(ctx, &biz.GetBlockRequest{
		Workchain: -1,
		Shard:     8000000000000000,
		SeqNo:     39382597,
	})
	s.Require().NoError(err)

	ctx.Debug("get block", zap.Any("block", &block))
}

func (s *suiteExternal) Test_impl_FetchAndStoreBlock() {
	ctx := contextx.Background()
	block, err := s.biz.FetchAndStoreBlock(ctx, &biz.FetchAndStoreBlockRequest{
		Workchain: -1,
		Shard:     8000000000000000,
		SeqNo:     39382597,
	})
	s.Require().NoError(err)

	ctx.Debug("fetch and store block", zap.Any("block", &block))
}
