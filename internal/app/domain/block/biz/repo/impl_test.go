package repo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

func (s *suiteTest) assert() {
	err := s.rw.ExpectationsWereMet()
	if err != nil {
		s.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTest))
}
