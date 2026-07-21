package xdb

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/jindasoft/jinda-platform/xauth"
	"github.com/jindasoft/jinda-platform/xentities"
	"github.com/stoewer/go-strcase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	errDataNotFound        = "data not found: %w"
	errFailedToDecode      = "failed to decode data: %w"
	errFailedToPrepare     = "failed to prepare data: %w"
	errFailedToCloseCursor = "failed to close cursor: %v\n"
)

func (s *service) IsNotFound(err error) bool {
	return errors.Is(err, mongo.ErrNoDocuments)
}

func (s *service) FindOne(ctx context.Context, filter bson.D, entity any) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	// check record is not soft deleted
	filter = append(filter, bson.E{Key: "deleted_at", Value: nil})

	if result := collection.FindOne(ctx, filter).Decode(entity); result != nil {
		return wrapFindOneError(result)
	}

	return nil
}

func (s *service) FindByID(ctx context.Context, oid primitive.ObjectID, entity any) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	// check record is not soft deleted
	filter := bson.D{
		{Key: "_id", Value: oid},
		{Key: "deleted_at", Value: nil},
	}

	if result := collection.FindOne(ctx, filter).Decode(entity); result != nil {
		return wrapFindOneError(result)
	}

	return nil
}

func (s *service) FindPaging(ctx context.Context, filter bson.D, sort bson.D, offset, limit int64, entity any) error {
	entityType := reflect.TypeOf(entity).Elem().Elem()
	entityName := entityType.Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	// check record is not soft deleted
	filter = append(filter, bson.E{Key: "deleted_at", Value: nil})

	findOptions := options.Find()
	findOptions.SetSkip(offset)
	findOptions.SetLimit(limit)
	findOptions.SetSort(sort)

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return fmt.Errorf("failed to find data: %w", err)
	}
	defer func() {
		if cerr := cursor.Close(ctx); cerr != nil {
			fmt.Printf(errFailedToCloseCursor, cerr)
		}
	}()

	if err := cursor.All(ctx, entity); err != nil {
		return fmt.Errorf(errFailedToDecode, err)
	}

	return nil
}

func (s *service) Find(ctx context.Context, filter bson.D, entity any, opts ...*options.FindOptions) error {
	entityType := reflect.TypeOf(entity).Elem().Elem()
	entityName := entityType.Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	// check record is not soft deleted
	filter = append(filter, bson.E{Key: "deleted_at", Value: nil})

	var findOptions *options.FindOptions
	if len(opts) > 0 {
		findOptions = opts[0]
	}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return fmt.Errorf("failed to find data: %w", err)
	}
	defer func() {
		if cerr := cursor.Close(ctx); cerr != nil {
			fmt.Printf(errFailedToCloseCursor, cerr)
		}
	}()

	if err := cursor.All(ctx, entity); err != nil {
		return fmt.Errorf(errFailedToDecode, err)
	}

	return nil
}

func (s *service) AggregateSample(ctx context.Context, filter bson.D, sampleSize int, entity any) error {
	entityType := reflect.TypeOf(entity).Elem().Elem()
	entityName := entityType.Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	// check record is not soft deleted
	filter = append(filter, bson.E{Key: "deleted_at", Value: nil})
	match := bson.D{{Key: "$match", Value: filter}}
	sample := bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: sampleSize}}}}
	pipeline := mongo.Pipeline{match, sample}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return fmt.Errorf("failed to aggregate data: %w", err)
	}
	defer func() {
		if cerr := cursor.Close(ctx); cerr != nil {
			fmt.Printf(errFailedToCloseCursor, cerr)
		}
	}()

	if err := cursor.All(ctx, entity); err != nil {
		return fmt.Errorf(errFailedToDecode, err)
	}

	return nil
}

func (s *service) InsertOne(ctx context.Context, entity any) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	if err := entity.(xentities.MongoBefore).BeforeCreate(ctx); err != nil {
		return fmt.Errorf(errFailedToPrepare, err)
	}

	if _, err := collection.InsertOne(ctx, entity); err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	return nil
}

func (s *service) InsertOneNoActor(ctx context.Context, entity any) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	if err := entity.(xentities.MongoBefore).BeforeCreateNoActor(ctx); err != nil {
		return fmt.Errorf(errFailedToPrepare, err)
	}

	if _, err := collection.InsertOne(ctx, entity); err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	return nil
}

func (s *service) UpdateOne(ctx context.Context, filter bson.D, entity any) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	if err := entity.(xentities.MongoBefore).BeforeUpdate(ctx); err != nil {
		return fmt.Errorf(errFailedToPrepare, err)
	}

	// check record is not soft deleted
	filter = append(filter, bson.E{Key: "deleted_at", Value: nil})

	update := bson.M{"$set": entity}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("failed to update data: %w", err)
	}

	return nil
}

func (s *service) UpdateOneNoActor(ctx context.Context, filter bson.D, entity any) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	if err := entity.(xentities.MongoBefore).BeforeUpdateNoActor(ctx); err != nil {
		return fmt.Errorf(errFailedToPrepare, err)
	}

	// check record is not soft deleted
	filter = append(filter, bson.E{Key: "deleted_at", Value: nil})

	update := bson.M{"$set": entity}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("failed to update data: %w", err)
	}

	return nil
}

func (s *service) UpdateIncrement(ctx context.Context, filter bson.D, field string, value int64, entity any) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	update := bson.M{"$inc": bson.M{field: value}}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("failed to increment field: %w", err)
	}

	return nil
}

func (s *service) SoftDeleteOne(ctx context.Context, filter bson.D, entity any) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	accountID, err := xauth.GetAccountOID(ctx)
	if err != nil {
		return err
	}

	// check record is not soft deleted
	filter = append(filter, bson.E{Key: "deleted_at", Value: nil})

	update := bson.M{
		"$set": bson.M{
			"status":     xentities.StatusDeleted,
			"deleted_by": accountID,
			"deleted_at": time.Now(),
		},
	}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("failed to soft delete data: %w", err)
	}

	return nil
}

func (s *service) SoftDeleteMany(ctx context.Context, filter bson.D, entity any) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	accountID, err := xauth.GetAccountOID(ctx)
	if err != nil {
		return err
	}

	// check record is not soft deleted
	filter = append(filter, bson.E{Key: "deleted_at", Value: nil})
	update := bson.M{
		"$set": bson.M{
			"status":     xentities.StatusDeleted,
			"deleted_by": accountID,
			"deleted_at": time.Now(),
		},
	}

	if _, err := collection.UpdateMany(ctx, filter, update); err != nil {
		return fmt.Errorf("failed to soft delete data: %w", err)
	}

	return nil
}

func (s *service) ForceDeleteOne(ctx context.Context, filter bson.D, entity any) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("failed to force delete data: %w", err)
	}

	return nil
}

func (s *service) SetIndexTtl(ctx context.Context, entity any, field string, expireAfter int32) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: field, Value: 1},
		},
		Options: options.Index().SetExpireAfterSeconds(expireAfter), // x seconds after "expireAt"
	}

	if _, err := collection.Indexes().CreateOne(ctx, indexModel); err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}

	return nil
}

func (s *service) SetIndexSearch(ctx context.Context, entity any, keys bson.D, unique bool) error {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	indexModel := mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetUnique(unique),
	}

	if _, err := collection.Indexes().CreateOne(ctx, indexModel); err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}

	return nil
}

func (s *service) SetIndexUnique(ctx context.Context, entity any, keys bson.D) error {
	return s.SetIndexSearch(ctx, entity, keys, true)
}

func (s *service) Count(ctx context.Context, filter bson.D, entity any) (int64, error) {
	entityName := reflect.TypeOf(entity).Elem().Name()
	entitySnake := strcase.SnakeCase(entityName)
	collection := s.mongo.Collection(entitySnake)

	filter = append(filter, bson.E{Key: "deleted_at", Value: nil})

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("failed to count documents: %w", err)
	}

	return count, nil
}
