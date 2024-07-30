package block

import (
	"testing"

	"github.com/blackhorseya/ryze/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
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

func (s *suiteMongodbTester) Test_mongodb_Create() {
	block, err := model.NewBlock(-1, 8000000000000000, 39382597)
	s.Require().NoError(err)

	type args struct {
		ctx  contextx.Contextx
		item *model.Block
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{item: block},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err = s.repo.Create(tt.args.ctx, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
