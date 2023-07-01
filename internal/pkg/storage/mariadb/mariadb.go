package mariadb

import (
	"time"

	"github.com/blackhorseya/ryze/internal/pkg/config"
	_ "github.com/go-sql-driver/mysql" // import MySQL driver
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql" // import MySQL driver
	_ "github.com/golang-migrate/migrate/v4/source/github"  // import GitHub source
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// NewMariadb init mariadb client
func NewMariadb(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", cfg.DB.URL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect database")
	}

	conns := cfg.DB.Conns
	if conns == 0 {
		conns = 100
	}

	db.SetConnMaxLifetime(15 * time.Minute)
	db.SetMaxOpenConns(conns)
	db.SetMaxIdleConns(conns)

	return db, nil
}

// NewMigration init migration
func NewMigration(cfg *config.Config, rw *sqlx.DB) (*migrate.Migrate, error) {
	instance, err := mysql.WithInstance(rw.DB, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(cfg.DB.Source, "mysql", instance)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// ProviderSet is a provider set for mariadb client
var ProviderSet = wire.NewSet(NewMariadb, NewMigration)
