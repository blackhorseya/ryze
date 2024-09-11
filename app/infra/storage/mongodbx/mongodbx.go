package mongodbx

import (
	"context"
	"fmt"
	"time"

	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	defaultTimeout  = 5 * time.Second
	defaultLimit    = int64(10)
	defaultMaxLimit = int64(100)

	dbName = "ryze"
)

// NewClientWithDSN returns a new mongo client with dsn.
func NewClientWithDSN(dsn string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(dsn).
		SetMaxPoolSize(500).
		SetMinPoolSize(10).
		SetMaxConnIdleTime(10 * time.Minute).
		SetConnectTimeout(10 * time.Second).
		SetRetryWrites(true).
		SetServerSelectionTimeout(5 * time.Second)

	client, err := mongo.Connect(contextx.Background(), opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewClient returns a new mongo client.
func NewClient(app *configx.Application) (*mongo.Client, error) {
	return NewClientWithDSN(app.Storage.Mongodb.DSN)
}

// Container is used to represent a mongodb container.
type Container struct {
	*mongodb.MongoDBContainer
}

// NewContainer returns a new mongodb container.
func NewContainer(ctx contextx.Contextx) (*Container, error) {
	container, err := mongodb.Run(ctx, "mongo:7")
	if err != nil {
		return nil, fmt.Errorf("run mongodb container: %w", err)
	}

	return &Container{
		MongoDBContainer: container,
	}, nil
}

// RW returns a read-write client.
func (c *Container) RW(ctx contextx.Contextx) (*mongo.Client, error) {
	dsn, err := c.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}

	return mongo.Connect(ctx, options.Client().ApplyURI(dsn))
}

func initTimeSeriesByName(rw *mongo.Client, dbName, collName string) error {
	c, cancelFunc := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelFunc()

	ctx := contextx.Background()

	db := rw.Database(dbName)

	exists, err := collectionOrViewExists(db, collName)
	if err != nil {
		ctx.Error("failed to check collection exists", zap.Error(err))
		return err
	}
	ctx.Debug("collection exists", zap.String("collection", collName), zap.Bool("exists", exists))

	if !exists {
		tsOptions := options.CreateCollection().SetTimeSeriesOptions(
			options.TimeSeries().
				SetTimeField("timestamp").
				SetMetaField("metadata").
				SetGranularity("hours"),
		)

		err = db.CreateCollection(c, collName, tsOptions)
		if err != nil {
			ctx.Error("failed to create collection", zap.Error(err))
			return err
		}
	}

	ctx.Info("collections are initialized")
	return nil
}

func collectionOrViewExists(db *mongo.Database, collectionName string) (bool, error) {
	c, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	ctx := contextx.Background()

	filter := bson.D{{Key: "name", Value: collectionName}}
	collections, err := db.ListCollections(c, filter)
	if err != nil {
		return false, err
	}

	var result bson.M
	if collections.Next(c) {
		if err = collections.Decode(&result); err != nil {
			return false, err
		}
		// Check if it is a view or a collection
		if kind, ok := result["type"].(string); ok && kind == "view" {
			ctx.Debug("The collection %s is a view.", zap.String("collection", collectionName))
		}
		return true, nil
	}

	return false, nil
}
