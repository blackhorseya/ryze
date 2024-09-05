//go:build wireinject

//go:generate wire

package block

import (
	"github.com/blackhorseya/ryze/app/domain/block/repo/block"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProviderSet is used to provide a new model.BlockServiceServer
var ProviderSet = wire.NewSet(
	NewBlockService,
	block.NewMongoDB,
)

func NewExternalBlockService(client *tonx.Client, rw *mongo.Client) biz.BlockServiceServer {
	panic(wire.Build(ProviderSet))
}
