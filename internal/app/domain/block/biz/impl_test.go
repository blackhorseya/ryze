package biz

import (
	"testing"

	"github.com/blackhorseya/ryze/internal/app/domain/block/biz/repo"
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
