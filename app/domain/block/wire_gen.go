// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package block

import (
	"github.com/blackhorseya/ryze/app/infra/storage/mongodbx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func NewExternalBlockService(client *tonx.Client, rw *mongo.Client) biz.BlockServiceServer {
	iBlockRepo := mongodbx.NewBlockRepo(rw)
	blockServiceServer := NewBlockService(client, iBlockRepo)
	return blockServiceServer
}

// wire.go:

// ProviderSet is used to provide a new model.BlockServiceServer
var ProviderSet = wire.NewSet(
	NewBlockService, mongodbx.NewBlockRepo,
)