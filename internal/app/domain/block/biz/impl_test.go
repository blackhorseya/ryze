package biz

import (
	"testing"

	"github.com/blackhorseya/ryze/internal/app/domain/block/biz/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	bb "github.com/blackhorseya/ryze/pkg/entity/domain/block/biz"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTest struct {
	suite.Suite

	logger *zap.Logger
	ctrl   *gomock.Controller

	biz  bb.IBiz
	repo *repo.MockIRepo
}

func (s *suiteTest) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.ctrl = gomock.NewController(s.T())
	s.repo = repo.NewMockIRepo(s.ctrl)
	s.biz = CreateBlockBiz(s.repo)
}

func (s *suiteTest) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTest))
}

func (s *suiteTest) Test_impl_ListenNewBlock() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.ListenNewBlock(contextx.BackgroundWithLogger(s.logger)); (err != nil) != tt.wantErr {
				t.Errorf("ListenNewBlock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
