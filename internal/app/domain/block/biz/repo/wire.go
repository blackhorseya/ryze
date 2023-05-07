//go:generate wire
//go:build wireinject

package repo

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateTestRepo(logger *zap.Logger, eth *ethclient.Client, rw *sqlx.DB, m *migrate.Migrate, writer *kafka.Writer) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}
