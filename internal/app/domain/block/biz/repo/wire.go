//go:generate wire
//go:build wireinject

package repo

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateTestRepo(eht *ethclient.Client, rw *sqlx.DB) IRepo {
	panic(wire.Build(testProviderSet))
}
