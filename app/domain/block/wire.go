package block

import (
	"github.com/blackhorseya/ryze/app/infra/storage/mongodbx"
	"github.com/google/wire"
)

// ProviderSet is used to provide a new model.BlockServiceServer
var ProviderSet = wire.NewSet(
	NewBlockService,
	mongodbx.NewBlockRepo,
)
