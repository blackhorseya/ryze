package repo

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/ryze/internal/app/domain/block/biz/repo/dao"
	"github.com/blackhorseya/ryze/pkg/contextx"
	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/ethereum/go-ethereum/common"
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
	s.repo, _ = CreateTestRepo(s.logger, nil, sqlx.NewDb(db, "mysql"), nil)
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

func (s *suiteTest) Test_impl_ListBlocks() {
	columns := []string{"number", "hash", "parent_hash", "timestamp"}
	selection := `SELECT number, hash, parent_hash, timestamp FROM blocks`
	count := fmt.Sprintf(`SELECT COUNT(*) AS total FROM (%s) AS t`, selection)
	query := []string{selection}

	// order by
	query = append(query, `ORDER BY timestamp DESC`)

	type args struct {
		condition ListBlocksCondition
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
			name: "count blocks then error",
			args: args{condition: ListBlocksCondition{Limit: 10, Offset: 0}, mock: func() {
				s.rw.ExpectQuery(regexp.QuoteMeta(count)).
					WillReturnError(errors.New("error"))
			}},
			wantRecords: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "list blocks then error",
			args: args{condition: ListBlocksCondition{Limit: 10, Offset: 10}, mock: func() {
				s.rw.ExpectQuery(regexp.QuoteMeta(count)).
					WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(100))

				newQuery := query
				newQuery = append(newQuery, `LIMIT ? OFFSET ?`)
				stmt := strings.Join(newQuery, " ")

				s.rw.ExpectQuery(regexp.QuoteMeta(stmt)).
					WithArgs(10, 10).
					WillReturnError(errors.New("error"))
			}},
			wantRecords: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "not found then return total",
			args: args{condition: ListBlocksCondition{Limit: 10, Offset: 10}, mock: func() {
				s.rw.ExpectQuery(regexp.QuoteMeta(count)).
					WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(100))

				newQuery := query
				newQuery = append(newQuery, `LIMIT ? OFFSET ?`)
				stmt := strings.Join(newQuery, " ")

				s.rw.ExpectQuery(regexp.QuoteMeta(stmt)).
					WithArgs(10, 10).
					WillReturnRows(sqlmock.NewRows(columns))
			}},
			wantRecords: nil,
			wantTotal:   100,
			wantErr:     false,
		},
		{
			name: "ok",
			args: args{condition: ListBlocksCondition{Limit: 10, Offset: 10}, mock: func() {
				s.rw.ExpectQuery(regexp.QuoteMeta(count)).
					WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(100))

				newQuery := query
				newQuery = append(newQuery, `LIMIT ? OFFSET ?`)
				stmt := strings.Join(newQuery, " ")

				s.rw.ExpectQuery(regexp.QuoteMeta(stmt)).
					WithArgs(10, 10).
					WillReturnRows(sqlmock.NewRows(columns).
						AddRow(1, common.HexToHash("hash1").Bytes(), common.HexToHash("parent_hash1").Bytes(), timestamp1.AsTime()).
						AddRow(2, common.HexToHash("hash2").Bytes(), common.HexToHash("parent_hash2").Bytes(), timestamp2.AsTime()))
			}},
			wantRecords: []*bm.Block{
				{
					Number:           1,
					Hash:             common.HexToHash("hash1").Hex(),
					ParentHash:       common.HexToHash("parent_hash1").Hex(),
					Nonce:            "",
					Sha3Uncles:       "",
					LogsBloom:        "",
					TransactionsRoot: "",
					StateRoot:        "",
					ReceiptsRoot:     "",
					Miner:            "",
					Difficulty:       0,
					TotalDifficulty:  0,
					ExtraData:        "",
					Size:             0,
					GasLimit:         0,
					GasUsed:          0,
					Timestamp:        timestamp1,
					Transactions:     nil,
					Uncles:           nil,
				},
				{
					Number:           2,
					Hash:             common.HexToHash("hash2").Hex(),
					ParentHash:       common.HexToHash("parent_hash2").Hex(),
					Nonce:            "",
					Sha3Uncles:       "",
					LogsBloom:        "",
					TransactionsRoot: "",
					StateRoot:        "",
					ReceiptsRoot:     "",
					Miner:            "",
					Difficulty:       0,
					TotalDifficulty:  0,
					ExtraData:        "",
					Size:             0,
					GasLimit:         0,
					GasUsed:          0,
					Timestamp:        timestamp2,
					Transactions:     nil,
					Uncles:           nil,
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

			gotRecords, gotTotal, err := s.repo.ListBlocks(contextx.BackgroundWithLogger(s.logger), tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListBlocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("ListBlocks() \ngotRecords = %+v, \nwant %+v", gotRecords, tt.wantRecords)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListBlocks() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}

			s.assert(t)
		})
	}
}

func (s *suiteTest) Test_impl_GetBlockByHash() {
	hash := common.Hex2Bytes("0x1")
	columns := []string{"number", "hash", "parent_hash", "timestamp"}
	stmt := `SELECT number, hash, parent_hash, timestamp FROM blocks WHERE hash = ?`

	type args struct {
		hash []byte
		mock func()
	}
	tests := []struct {
		name       string
		args       args
		wantRecord *bm.Block
		wantErr    bool
	}{
		{
			name: "get block by hash then error",
			args: args{hash: hash, mock: func() {
				s.rw.ExpectQuery(regexp.QuoteMeta(stmt)).
					WithArgs(hash).
					WillReturnError(errors.New("error"))
			}},
			wantRecord: nil,
			wantErr:    true,
		},
		{
			name: "not found block then return nil",
			args: args{hash: hash, mock: func() {
				s.rw.ExpectQuery(regexp.QuoteMeta(stmt)).
					WithArgs(hash).
					WillReturnRows(sqlmock.NewRows(columns))
			}},
			wantRecord: nil,
			wantErr:    false,
		},
		{
			name: "ok",
			args: args{hash: hash, mock: func() {
				block := dao.NewBlock(block1)

				s.rw.ExpectQuery(regexp.QuoteMeta(stmt)).
					WithArgs(hash).
					WillReturnRows(sqlmock.NewRows(columns).AddRow(
						block.Number,
						block.Hash,
						block.ParentHash,
						block.Timestamp,
					))
			}},
			wantRecord: block1,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotRecord, err := s.repo.GetBlockByHash(contextx.BackgroundWithLogger(s.logger), tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockByHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("GetBlockByHash() gotRecord = %v, want %v", gotRecord, tt.wantRecord)
			}

			s.assert(t)
		})
	}
}
