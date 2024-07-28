package biz

import (
	"testing"
	"time"

	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type suiteTester struct {
	suite.Suite

	ctrl   *gomock.Controller
	client *tonx.Client
	biz    model.BlockServiceServer
}

func (s *suiteTester) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.client, _ = tonx.NewClient(tonx.Options{Network: "testnet"})
	s.biz = NewBlockService(s.client)
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_ScanBlock() {
	stream := model.NewMockBlockService_ScanBlockServer(s.ctrl)

	timeout, cancelFunc := contextx.WithTimeout(contextx.Background(), 2*time.Second)
	defer cancelFunc()

	stream.EXPECT().Context().Return(timeout).Times(1)
	_ = s.biz.ScanBlock(&model.ScanBlockRequest{}, stream)
}
