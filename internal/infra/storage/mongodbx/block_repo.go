package mongodbx

import (
	"context"
	"fmt"

	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/entity/domain/block/repo"
	"github.com/blackhorseya/ryze/internal/infra/otelx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type mongodbBlockRepo struct {
	coll *mongo.Collection
}

// NewBlockRepo is used to create an implementation of the block repository.
func NewBlockRepo(rw *mongo.Client) (repo.IBlockRepo, error) {
	collName := "blocks"

	err := initTimeSeriesByName(rw, dbName, collName)
	if err != nil {
		return nil, fmt.Errorf("failed to init collections: %w", err)
	}

	coll := rw.Database(dbName).Collection(collName)

	return &mongodbBlockRepo{
		coll: coll,
	}, nil
}

func (i *mongodbBlockRepo) GetByID(c context.Context, id string) (item *model.Block, err error) {
	_, span := otelx.Tracer.Start(c, "block.biz.block.mongodbBlockRepo.GetByID")
	defer span.End()

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.WithContext(c)

	var got blockDocument
	filter := bson.M{"metadata._id": id}
	err = i.coll.FindOne(timeout, filter).Decode(&got)
	if err != nil {
		ctx.Error("failed to find a block from mongodbBlockRepo", zap.Error(err), zap.Any("id", id))
		return nil, err
	}

	return got.Metadata, nil
}

func (i *mongodbBlockRepo) Create(c context.Context, item *model.Block) (err error) {
	_, span := otelx.Tracer.Start(c, "block.biz.block.mongodbBlockRepo.Create")
	defer span.End()

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.WithContext(c)

	doc := newBlockDocument(item)
	_, err = i.coll.InsertOne(timeout, doc)
	if err != nil {
		ctx.Error("failed to insert a block to mongodbBlockRepo", zap.Error(err))
		return err
	}

	return nil
}

func (i *mongodbBlockRepo) List(
	c context.Context,
	cond repo.ListCondition,
) (items []*model.Block, total int, err error) {
	_, span := otelx.Tracer.Start(c, "block.biz.block.mongodbBlockRepo.List")
	defer span.End()

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.WithContext(c)

	filter := bson.M{}

	limit, skip := defaultLimit, int64(0)
	if 0 < cond.Limit && cond.Limit <= defaultMaxLimit {
		limit = cond.Limit
	}
	if cond.Skip > 0 {
		skip = cond.Skip
	}
	opts := options.Find().SetSort(bson.M{"timestamp": -1}).SetLimit(limit).SetSkip(skip)

	cur, err := i.coll.Find(timeout, filter, opts)
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

	count, err := i.coll.CountDocuments(timeout, filter)
	if err != nil {
		ctx.Error("failed to count blocks from mongodbBlockRepo", zap.Error(err))
		return nil, 0, err
	}

	return items, int(count), nil
}
