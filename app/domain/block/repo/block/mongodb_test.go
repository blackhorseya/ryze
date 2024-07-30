package block

import (
	"testing"

	"github.com/blackhorseya/ryze/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/entity/domain/block/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteMongodbTester struct {
	suite.Suite

	container *mongodbx.Container
	rw        *mongo.Client
	repo      repo.IBlockRepo
}

func (s *suiteMongodbTester) SetupTest() {
	container, err := mongodbx.NewContainer(contextx.Background())
	s.Require().NoError(err)
	s.container = container

	rw, err := container.RW(contextx.Background())
	s.Require().NoError(err)
	s.rw = rw

	s.repo = NewMongoDB(rw)
}

func (s *suiteMongodbTester) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}

	if s.container != nil {
		_ = s.container.Terminate(contextx.Background())
	}
}

func TestMongodbAll(t *testing.T) {
	suite.Run(t, new(suiteMongodbTester))
}
