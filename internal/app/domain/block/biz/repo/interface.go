//go:generate mockgen -destination=./mock_${GOFILE} -package=repo -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/ryze/pkg/contextx"
	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/google/wire"
)

// IRepo declare the interface of repo
type IRepo interface {
	// GetBlockByHash serve caller to given block hash to get block info
	GetBlockByHash(ctx contextx.Contextx, hash []byte) (record *bm.Block, err error)

	// GetBlockByHeight serve caller to given block height to get block info
	GetBlockByHeight(ctx contextx.Contextx, height uint64) (record *bm.Block, err error)

	// ListenNewBlock serve caller to listen new block
	ListenNewBlock(ctx contextx.Contextx) (newBlockChan <-chan *bm.Block, err error)
}

// BlockSet is the provider set of repo
var BlockSet = wire.NewSet(NewEthOptions, NewImpl)
