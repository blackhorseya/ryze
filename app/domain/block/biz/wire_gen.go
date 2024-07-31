// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package biz

import (
	"github.com/blackhorseya/ryze/app/domain/block/repo/block"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func NewExternalBlockService(client *tonx.Client, rw *mongo.Client) model.BlockServiceServer {
	iBlockRepo := block.NewMongoDB(rw)
	blockServiceServer := NewBlockService(client, iBlockRepo)
	return blockServiceServer
}

// wire.go:

// ProviderBlockServiceSet is used to provide a new model.BlockServiceServer
var ProviderBlockServiceSet = wire.NewSet(NewBlockService, block.NewMongoDB)
