package block

import (
	"time"

	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/entity/domain/block/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	defaultTimeout = 5 * time.Second
	dbName         = "ryze"
	collName       = "blocks"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongoDB is used to create an implementation of the block repository.
func NewMongoDB(rw *mongo.Client) repo.IBlockRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Block, err error) {
	ctx, span := otelx.Span(ctx, "block.biz.block.mongodb.GetByID")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	var got blockDocument
	filter := bson.M{}
	err = i.rw.Database(dbName).Collection(collName).FindOne(timeout, filter).Decode(&got)
	if err != nil {
		ctx.Error("failed to find a block from mongodb", zap.Error(err), zap.Any("id", id))
		return nil, err
	}

	return got.Metadata, nil
}

func (i *mongodb) Create(ctx contextx.Contextx, item *model.Block) (err error) {
	ctx, span := otelx.Span(ctx, "block.biz.block.mongodb.create")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	doc := newBlockDocument(item)
	_, err = i.rw.Database(dbName).Collection(collName).InsertOne(timeout, doc)
	if err != nil {
		ctx.Error("failed to insert a block to mongodb", zap.Error(err))
		return err
	}

	return nil
}

func (i *mongodb) Update(ctx contextx.Contextx, item *model.Block) (err error) {
	// TODO: 2024/7/31|sean|implement me
	panic("implement me")
}

func (i *mongodb) Delete(ctx contextx.Contextx, id string) (err error) {
	// TODO: 2024/7/31|sean|implement me
	panic("implement me")
}
