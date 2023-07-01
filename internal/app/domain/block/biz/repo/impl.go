package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/blackhorseya/ryze/internal/app/domain/block/biz/repo/dao"
	"github.com/blackhorseya/ryze/internal/pkg/config"
	"github.com/blackhorseya/ryze/pkg/contextx"
	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	topicNewBlock = "new_block"
)

// NewEthClient serve caller to get a new ethclient.Client instance
func NewEthClient(cfg *config.Config, logger *zap.Logger) (*ethclient.Client, error) {
	if cfg.ETH.Websocket == "" {
		return nil, nil
	}

	client, err := ethclient.Dial(cfg.ETH.Websocket)
	if err != nil {
		return nil, errors.Wrap(err, "dial eth client failed")
	}

	logger.Info("dial eth client success")

	return client, nil
}

type impl struct {
	rw     *sqlx.DB
	socket *ethclient.Client
	writer *kafka.Writer
}

// NewImpl serve caller to get a new IRepo implementation instance
func NewImpl(logger *zap.Logger, rw *sqlx.DB, socket *ethclient.Client, m *migrate.Migrate, writer *kafka.Writer) (IRepo, error) {
	if m != nil {
		err := m.Up()
		if err != nil && err != migrate.ErrNoChange {
			if !errors.Is(err, migrate.ErrNoChange) {
				return nil, errors.Wrap(err, "migrate up error")
			}
		}
	}

	return &impl{
		socket: socket,
		rw:     rw,
		writer: writer,
	}, nil
}

func (i *impl) GetBlockByHash(ctx contextx.Contextx, hash []byte) (record *bm.Block, err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, 3*time.Second)
	defer cancelFunc()

	stmt := `SELECT number, hash, parent_hash, timestamp FROM blocks WHERE hash = ?`

	var got dao.Block
	err = i.rw.GetContext(timeout, &got, stmt, hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		ctx.Error("select block by hash failed", zap.Error(err), zap.String("stmt", stmt), zap.Any("args", hash))
		return nil, err
	}

	return got.ToEntity(), nil
}

func (i *impl) ListBlocks(ctx contextx.Contextx, condition ListBlocksCondition) (records []*bm.Block, total uint, err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, 3*time.Second)
	defer cancelFunc()

	selection := `SELECT number, hash, parent_hash, timestamp FROM blocks`
	count := fmt.Sprintf("SELECT COUNT(*) AS total FROM (%s) AS t", selection)

	err = i.rw.QueryRowxContext(timeout, count).Scan(&total)
	if err != nil {
		ctx.Error("count total blocks failed", zap.Error(err), zap.String("stmt", count))
		return nil, 0, err
	}

	query := []string{selection}
	var args []interface{}

	// append order by
	query = append(query, `ORDER BY timestamp DESC`)

	// append limit
	if condition.Limit > 0 {
		query = append(query, `LIMIT ?`)
		args = append(args, condition.Limit)
	}

	// append offset
	if condition.Offset > 0 {
		query = append(query, `OFFSET ?`)
		args = append(args, condition.Offset)
	}

	// concat query to stmt
	stmt := strings.Join(query, " ")

	var got dao.Blocks
	err = i.rw.SelectContext(timeout, &got, stmt, args...)
	if err != nil {
		ctx.Error("select blocks failed", zap.Error(err), zap.String("stmt", stmt), zap.Any("args", args))
		return nil, 0, err
	}

	ret := got.ToEntity()

	return ret, total, nil
}

func (i *impl) CreateNewBlock(ctx contextx.Contextx, newBlock *bm.Block) error {
	timeout, cancelFunc := contextx.WithTimeout(ctx, 1*time.Second)
	defer cancelFunc()

	created := dao.NewBlock(newBlock)
	stmt := `INSERT INTO blocks (number, hash, parent_hash, timestamp) VALUES (:number, :hash, :parent_hash, :timestamp)`

	_, err := i.rw.NamedExecContext(timeout, stmt, created)
	if err != nil {
		return err
	}

	return nil
}

func (i *impl) SubscribeNewBlock(ctx contextx.Contextx) (newBlockChan <-chan *bm.Block, err error) {
	headers := make(chan *types.Header)
	sub, err := i.socket.SubscribeNewHead(ctx, headers)
	if err != nil {
		return nil, err
	}

	blocks := make(chan *bm.Block)
	go func() {
		for {
			select {
			case err = <-sub.Err():
				ctx.Warn("subscribe new block failed", zap.Error(err))
			case header := <-headers:
				block := &bm.Block{
					Number:           header.Number.Uint64(),
					Hash:             header.Hash().Hex(),
					ParentHash:       header.ParentHash.Hex(),
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
					Timestamp:        timestamppb.New(time.Unix(int64(header.Time), 0)),
					Transactions:     nil,
					Uncles:           nil,
				}

				blocks <- block
			}
		}
	}()

	return blocks, nil
}

func (i *impl) PublishNewBlock(ctx contextx.Contextx, newBlock *bm.Block) error {
	value, err := json.Marshal(newBlock)
	if err != nil {
		return errors.Wrap(err, "marshal new block failed")
	}

	err = i.writer.WriteMessages(ctx, kafka.Message{
		Topic: topicNewBlock,
		Key:   []byte(newBlock.Hash),
		Value: value,
		Time:  time.Now(),
	})
	if err != nil {
		return errors.Wrap(err, "write new block to kafka failed")
	}

	return nil
}
