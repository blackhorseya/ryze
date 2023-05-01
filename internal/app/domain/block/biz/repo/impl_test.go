package repo

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/ryze/internal/app/domain/block/biz/repo/dao"
	"github.com/blackhorseya/ryze/pkg/contextx"
	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTest struct {
	suite.Suite

	logger *zap.Logger

	rw   sqlmock.Sqlmock
	repo IRepo
}

func (s *suiteTest) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	db, mock, _ := sqlmock.New()
	s.rw = mock
	s.repo = CreateTestRepo(nil, sqlx.NewDb(db, "mysql"))
}

func (s *suiteTest) assert(t *testing.T) {
	err := s.rw.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTest))
}

func (s *suiteTest) Test_impl_CreateNewBlock() {
	stmt := `INSERT INTO blocks (number, hash, parent_hash, timestamp) VALUES (?, ?, ?, ?)`

	type args struct {
		newBlock *bm.Block
		mock     func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "insert new block then error",
			args: args{newBlock: block1, mock: func() {
				block := dao.NewBlock(block1)

				s.rw.ExpectExec(regexp.QuoteMeta(stmt)).
					WithArgs(
						block.Number,
						block.Hash,
						block.ParentHash,
						block.Timestamp,
					).
					WillReturnError(errors.New("error"))
			}},
			wantErr: true,
		},
		{
			name: "insert new block then error",
			args: args{newBlock: block1, mock: func() {
				block := dao.NewBlock(block1)

				s.rw.ExpectExec(regexp.QuoteMeta(stmt)).
					WithArgs(
						block.Number,
						block.Hash,
						block.ParentHash,
						block.Timestamp,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.CreateNewBlock(contextx.BackgroundWithLogger(s.logger), tt.args.newBlock); (err != nil) != tt.wantErr {
				t.Errorf("CreateNewBlock() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.assert(t)
		})
	}
}
