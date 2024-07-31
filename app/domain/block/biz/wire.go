//go:build wireinject

//go:generate wire

package biz

import (
	"github.com/blackhorseya/ryze/app/domain/block/repo/block"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProviderBlockServiceSet is used to provide a new model.BlockServiceServer
var ProviderBlockServiceSet = wire.NewSet(NewBlockService, block.NewMongoDB)

func NewExternalBlockService(client *tonx.Client, rw *mongo.Client) model.BlockServiceServer {
	panic(wire.Build(ProviderBlockServiceSet))
}
