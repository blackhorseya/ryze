package repo

import (
	"github.com/blackhorseya/ryze/pkg/contextx"
	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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
	socket *ethclient.Client
}

// NewImpl serve caller to get a new IRepo implementation instance
func NewImpl(o *EthOptions) (IRepo, error) {
	socket, err := ethclient.Dial(o.Websocket)
	if err != nil {
		return nil, errors.Wrap(err, "dial eth websocket failed")
	}

	instance := &impl{
		socket: socket,
	}

	return instance, nil
}

func (i *impl) GetBlockByHash(ctx contextx.Contextx, hash []byte) (record *bm.Block, err error) {
	// todo: 2023/4/29|sean|implement me
	panic("implement me")
}

func (i *impl) GetBlockByHeight(ctx contextx.Contextx, height uint64) (record *bm.Block, err error) {
	// todo: 2023/4/29|sean|implement me
	panic("implement me")
}

func (i *impl) ListenNewBlock(ctx contextx.Contextx) (newBlockChan <-chan *bm.Block, err error) {
	// todo: 2023/4/29|sean|implement me
	panic("implement me")
}
