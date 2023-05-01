package biz

import (
	"errors"
	"reflect"
	"testing"

	"github.com/blackhorseya/ryze/internal/app/domain/block/biz/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	bb "github.com/blackhorseya/ryze/pkg/entity/domain/block/biz"
	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
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

func (s *suiteTest) Test_impl_ListBlocks() {
	condition := bb.ListBlocksCondition{Page: 1, Size: 10}
	newCondition := newRepoCondition(condition)

	type args struct {
		condition bb.ListBlocksCondition
		mock      func()
	}
	tests := []struct {
		name        string
		args        args
		wantRecords []*bm.Block
		wantTotal   uint
		wantErr     bool
	}{
		{
			name:        "invalid page then error",
			args:        args{condition: bb.ListBlocksCondition{Page: 0, Size: 10}},
			wantRecords: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name:        "invalid size then error",
			args:        args{condition: bb.ListBlocksCondition{Page: 1, Size: 0}},
			wantRecords: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "list blocks then error",
			args: args{condition: condition, mock: func() {
				s.repo.EXPECT().ListBlocks(gomock.Any(), newCondition).Return(nil, uint(0), errors.New("error")).Times(1)
			}},
			wantRecords: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "ok",
			args: args{condition: condition, mock: func() {
				s.repo.EXPECT().ListBlocks(gomock.Any(), newCondition).Return([]*bm.Block{
					{
						Number: 1,
					},
				}, uint(100), nil).Times(1)
			}},
			wantRecords: []*bm.Block{
				{
					Number: 1,
				},
			},
			wantTotal: 100,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotRecords, gotTotal, err := s.biz.ListBlocks(contextx.BackgroundWithLogger(s.logger), tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListBlocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("ListBlocks() gotRecords = %v, want %v", gotRecords, tt.wantRecords)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListBlocks() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
