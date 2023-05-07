//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/ryze/internal/adapter/restful"
	"github.com/blackhorseya/ryze/internal/app/domain/block/biz"
	"github.com/blackhorseya/ryze/internal/app/domain/block/biz/repo"
	"github.com/blackhorseya/ryze/internal/pkg/config"
	"github.com/blackhorseya/ryze/internal/pkg/httpx"
	"github.com/blackhorseya/ryze/internal/pkg/log"
	"github.com/blackhorseya/ryze/internal/pkg/storage/mariadb"
	"github.com/blackhorseya/ryze/internal/pkg/transports/kafkax"
	"github.com/blackhorseya/ryze/pkg/app"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infra
	config.ProviderSet,
	log.ProviderSet,

	// transport
	httpx.ServerSet,
	kafkax.WriterSet,

	// storage
	mariadb.ProviderSet,

	// adapter
	restful.ProviderSet,

	// domain
	biz.BlockSet,
	repo.BlockSet,

	// main
	NewService,
)

// CreateApplication serve caller to create application instance
func CreateApplication(path string) (app.Servicer, error) {
	panic(wire.Build(providerSet))
}
