//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/ryze/internal/pkg/config"
	"github.com/blackhorseya/ryze/internal/pkg/log"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infra
	config.ProviderSet,
	log.ProviderSet,

	// main
	NewService,
)

// CreateService serve caller to create service instance
func CreateService(path string) (*Service, error) {
	panic(wire.Build(providerSet))
}
