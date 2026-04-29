package xdb

import (
	"context"
	"fmt"

	"github.com/jindasoft/jinda-platform/xlogger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoService interface {
	Find(ctx context.Context, filter bson.D, entity any, opts ...*options.FindOptions) error
	FindByID(ctx context.Context, oid primitive.ObjectID, entity any) error
	FindOne(ctx context.Context, filter bson.D, entity any) error
	FindPaging(ctx context.Context, filter bson.D, sort bson.D, offset, limit int64, entity any) error
	AggregateSample(ctx context.Context, filter bson.D, sampleSize int, entity any) error
	InsertOne(ctx context.Context, entity any) error
	InsertOneNoActor(ctx context.Context, entity any) error
	UpdateOne(ctx context.Context, filter bson.M, entity any) error
	UpdateOneNoActor(ctx context.Context, filter bson.M, entity any) error
	UpdateIncrement(ctx context.Context, filter bson.M, field string, value int64, entity any) error
	SoftDeleteOne(ctx context.Context, filter bson.M, entity any) error
	SoftDeleteMany(ctx context.Context, filter bson.M, entity any) error
	ForceDeleteOne(ctx context.Context, filter bson.M, entity any) error
	SetIndexTtl(ctx context.Context, entity any, field string, expireAfter int32) error
	SetIndexSearch(ctx context.Context, entity any, keys bson.D, unique bool) error
	Count(ctx context.Context, filter bson.M, entity any) (int64, error)
}

type MongoConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string `json:"-"`
	Options  string
	IsDebug  bool
}

type service struct {
	mongo *mongo.Database
}

func NewMongoService(ctx context.Context, cfg *MongoConfig) (*service, error) {
	client, err := mongoClient(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	mongo := client.Database(cfg.Database)
	if mongo == nil {
		return nil, fmt.Errorf("failed to get MongoDB database")
	}

	return &service{mongo: mongo}, nil
}

func mongoClient(ctx context.Context, cfg *MongoConfig) (*mongo.Client, error) {
	connection := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	if cfg.Options != "" {
		connection += fmt.Sprintf("/?%s", cfg.Options)
	}

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(connection).
		SetServerAPIOptions(serverAPI)

	if cfg.IsDebug {
		cmdMonitor := &event.CommandMonitor{
			Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
				xlogger.SysInfof("MongoDB CommandStartedEvent: %s", evt.Command.String())
			},
		}
		opts = opts.SetMonitor(cmdMonitor)
	}

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		xlogger.SysInfof("Error disconnecting MongoDB client: %v", err)
	// 	}
	// }()

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		xlogger.SysErrorf("Failed to ping MongoDB: %v", err)
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	xlogger.SysInfof("MongoDB initialized")

	return client, nil
}
