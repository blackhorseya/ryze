package mariadb

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import db driver
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
	URL   string `json:"url" yaml:"url"`
	Debug bool   `json:"debug" yaml:"debug"`
	Conns int    `json:"conns" yaml:"conns"`
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

// ProviderSet is a provider set for mariadb client
var ProviderSet = wire.NewSet(NewOptions, NewMariadb)
