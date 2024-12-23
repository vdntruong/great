package core

import (
	"context"
	"reflect"
)

type Entity interface {
	GetID() interface{}
	SetID(id interface{})
	Validate() error
}

type DAO[T Entity] interface {
	// Basic CRUD

	Create(ctx context.Context, entity T) error
	GetByID(ctx context.Context, id interface{}) (T, error)
	Update(ctx context.Context, entity T) error
	Delete(ctx context.Context, id interface{}) error

	// Batch operations

	CreateMany(ctx context.Context, entities []T) error
	GetMany(ctx context.Context, ids []interface{}) ([]T, error)
	UpdateMany(ctx context.Context, entities []T) error
	DeleteMany(ctx context.Context, ids []interface{}) error

	// Query operations

	Find(ctx context.Context, opts QueryOptions) ([]T, error)
	Count(ctx context.Context, opts QueryOptions) (int64, error)

	// Transaction support

	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}

type BaseDAO[T Entity] struct {
	entityType reflect.Type
	tableName  string
}
