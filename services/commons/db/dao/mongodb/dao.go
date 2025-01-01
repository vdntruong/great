package mongodb

import (
	"context"

	"commons/db/dao/core"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDAO[T core.Entity] struct {
	*core.BaseDAO[T]
	collection *mongo.Collection
}

func NewMongoDAO[T core.Entity](collection *mongo.Collection) *MongoDAO[T] {
	return &MongoDAO[T]{
		BaseDAO:    core.NewBaseDAO[T](),
		collection: collection,
	}
}

func (dao *MongoDAO[T]) Create(ctx context.Context, entity T) error {
	if err := entity.Validate(); err != nil {
		return err
	}

	result, err := dao.collection.InsertOne(ctx, entity)
	if err != nil {
		return err
	}

	entity.SetID(result.InsertedID)
	return nil
}
