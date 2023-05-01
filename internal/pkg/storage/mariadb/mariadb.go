package mariadb

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import MySQL driver
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql" // import MySQL driver
	_ "github.com/golang-migrate/migrate/v4/source/github"  // import GitHub source
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	defaultConns = 100
)

// Options is configuration of database
type Options struct {
	URL    string `json:"url" yaml:"url"`
	Debug  bool   `json:"debug" yaml:"debug"`
	Conns  int    `json:"conns" yaml:"conns"`
	Source string `json:"source" yaml:"source"`
}

// NewOptions serve caller to create an Options
func NewOptions(v *viper.Viper) (*Options, error) {
	o := new(Options)
	err := v.UnmarshalKey("db", o)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	return o, nil
}

// NewMariadb init mariadb client
func NewMariadb(o *Options) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", o.URL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect database")
	}

	conns := o.Conns
	if conns == 0 {
		conns = defaultConns
	}

	db.SetConnMaxLifetime(15 * time.Minute)
	db.SetMaxOpenConns(conns)
	db.SetMaxIdleConns(conns)

	return db, nil
}

// NewMigration init migration
func NewMigration(o *Options, rw *sqlx.DB) (*migrate.Migrate, error) {
	instance, err := mysql.WithInstance(rw.DB, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(o.Source, "mysql", instance)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// ProviderSet is a provider set for mariadb client
var ProviderSet = wire.NewSet(NewOptions, NewMariadb, NewMigration)
