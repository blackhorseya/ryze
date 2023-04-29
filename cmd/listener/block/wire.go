//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/ryze/internal/adapter/listener/block"
	"github.com/blackhorseya/ryze/internal/pkg/config"
	"github.com/blackhorseya/ryze/internal/pkg/log"
	"github.com/blackhorseya/ryze/pkg/app"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infra
	config.ProviderSet,
	log.ProviderSet,

	// adapter
	block.ListenerSet,

	// main
	NewService,
)

// CreateApplication serve caller to create application instance
func CreateApplication(path string) (app.Servicer, error) {
	panic(wire.Build(providerSet))
}
