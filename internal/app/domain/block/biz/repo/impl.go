package repo

import (
	"time"

	"github.com/blackhorseya/ryze/internal/app/domain/block/biz/repo/dao"
	"github.com/blackhorseya/ryze/pkg/contextx"
	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// EthOptions declare the options of ethereum client
type EthOptions struct {
	Endpoint  string `json:"endpoint" yaml:"endpoint"`
	Websocket string `json:"websocket" yaml:"websocket"`
}

// NewEthOptions serve caller to get a new EthOptions instance
func NewEthOptions(v *viper.Viper, logger *zap.Logger) (*EthOptions, error) {
	o := new(EthOptions)
	err := v.UnmarshalKey("eth", o)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal eth options failed")
	}

	logger.Info("get eth options success")

	return o, nil
}

type impl struct {
	o      *EthOptions
	rw     *sqlx.DB
	socket *ethclient.Client
}

// NewImpl serve caller to get a new IRepo implementation instance
func NewImpl(o *EthOptions, rw *sqlx.DB) IRepo {
	return &impl{
		o:      o,
		socket: nil,
		rw:     rw,
	}
}

func (i *impl) GetBlockByHash(ctx contextx.Contextx, hash []byte) (record *bm.Block, err error) {
	// todo: 2023/4/29|sean|implement me
	panic("implement me")
}

func (i *impl) GetBlockByHeight(ctx contextx.Contextx, height uint64) (record *bm.Block, err error) {
	// todo: 2023/4/29|sean|implement me
	panic("implement me")
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
	sub, err := i.getSocket().SubscribeNewHead(ctx, headers)
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

func (i *impl) getSocket() *ethclient.Client {
	if i.socket == nil {
		socket, err := ethclient.Dial(i.o.Websocket)
		if err != nil {
			panic(err)
		}

		i.socket = socket
	}

	return i.socket
}
