package mongodbx

import (
	"time"

	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/entity/domain/block/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	defaultTimeout = 5 * time.Second
	dbName         = "ryze"
	collName       = "blocks"
)

type mongodbBlockRepo struct {
	rw *mongo.Client
}

// NewBlockRepo is used to create an implementation of the block repository.
func NewBlockRepo(rw *mongo.Client) repo.IBlockRepo {
	return &mongodbBlockRepo{rw: rw}
}

func (i *mongodbBlockRepo) GetByID(ctx contextx.Contextx, id string) (item *model.Block, err error) {
	ctx, span := otelx.Span(ctx, "block.biz.block.mongodbBlockRepo.GetByID")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	var got blockDocument
	filter := bson.M{"metadata._id": id}
	err = i.rw.Database(dbName).Collection(collName).FindOne(timeout, filter).Decode(&got)
	if err != nil {
		ctx.Error("failed to find a block from mongodbBlockRepo", zap.Error(err), zap.Any("id", id))
		return nil, err
	}

	return got.Metadata, nil
}

func (i *mongodbBlockRepo) Create(ctx contextx.Contextx, item *model.Block) (err error) {
	ctx, span := otelx.Span(ctx, "block.biz.block.mongodbBlockRepo.Create")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	doc := newBlockDocument(item)
	_, err = i.rw.Database(dbName).Collection(collName).InsertOne(timeout, doc)
	if err != nil {
		ctx.Error("failed to insert a block to mongodbBlockRepo", zap.Error(err))
		return err
	}

	return nil
}

func (i *mongodbBlockRepo) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.Block, total int, err error) {
	ctx, span := otelx.Span(ctx, "block.biz.block.mongodbBlockRepo.List")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{}

	limit := condition.Limit
	if limit == 0 {
		limit = 10
	}
	skip := condition.Skip
	if skip < 0 {
		skip = 0
	}
	opts := options.Find().SetSort(bson.M{"timestamp": -1}).SetLimit(int64(limit)).SetSkip(int64(skip))

	cur, err := i.rw.Database(dbName).Collection(collName).Find(timeout, filter, opts)
	if err != nil {
		ctx.Error("failed to find blocks from mongodbBlockRepo", zap.Error(err))
		return nil, 0, err
	}
	defer cur.Close(timeout)

	for cur.Next(timeout) {
		var got blockDocument
		err = cur.Decode(&got)
		if err != nil {
			ctx.Error("failed to decode a block from mongodbBlockRepo", zap.Error(err))
			return nil, 0, err
		}

		items = append(items, got.Metadata)
	}

	count, err := i.rw.Database(dbName).Collection(collName).CountDocuments(timeout, filter)
	if err != nil {
		ctx.Error("failed to count blocks from mongodbBlockRepo", zap.Error(err))
		return nil, 0, err
	}

	return items, int(count), nil
}
