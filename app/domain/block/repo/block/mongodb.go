package block

import (
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/entity/domain/block/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongoDB is used to create an implementation of the block repository.
func NewMongoDB(rw *mongo.Client) repo.IBlockRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Block, err error) {
	// TODO: 2024/7/31|sean|implement me
	panic("implement me")
}

func (i *mongodb) Create(ctx contextx.Contextx, item *model.Block) (err error) {
	// TODO: 2024/7/31|sean|implement me
	panic("implement me")
}

func (i *mongodb) Update(ctx contextx.Contextx, item *model.Block) (err error) {
	// TODO: 2024/7/31|sean|implement me
	panic("implement me")
}

func (i *mongodb) Delete(ctx contextx.Contextx, id string) (err error) {
	// TODO: 2024/7/31|sean|implement me
	panic("implement me")
}
