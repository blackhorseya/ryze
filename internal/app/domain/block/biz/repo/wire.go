//go:generate wire
//go:build wireinject

package repo

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateTestRepo(o *EthOptions, rw *sqlx.DB) IRepo {
	panic(wire.Build(testProviderSet))
}
