//go:build external

package mongodbx

import (
	"testing"

	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/entity/domain/block/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteExternalMongodb struct {
	suite.Suite

	rw   *mongo.Client
	repo repo.IBlockRepo
}

func (s *suiteExternalMongodb) SetupTest() {
	config, err := configx.NewConfiguration(viper.GetViper())
	s.Require().NoError(err)

	app, err := config.GetService("block-grpc")
	s.Require().NoError(err)

	rw, err := NewClient(app)
	s.Require().NoError(err)

	s.rw = rw
	s.repo = NewBlockRepo(rw)
}

func (s *suiteExternalMongodb) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}
}

func TestExternalMongodbAll(t *testing.T) {
	suite.Run(t, new(suiteExternalMongodb))
}

func (s *suiteExternalMongodb) Test_mongodb_Create() {
	block, err := model.NewBlock(-1, 8000000000000000, 39382597)
	s.Require().NoError(err)

	err = s.repo.Create(contextx.Background(), block)
	s.Require().NoError(err)
}

func (s *suiteExternalMongodb) Test_mongodb_GetByID() {
	block, err := model.NewBlock(-1, 8000000000000000, 39382597)
	s.Require().NoError(err)

	err = s.repo.Create(contextx.Background(), block)
	s.Require().NoError(err)

	item, err := s.repo.GetByID(contextx.Background(), block.Id)
	s.Require().NoError(err)
	s.Require().Equal(block, item)
}
