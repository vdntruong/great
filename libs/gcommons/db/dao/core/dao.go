package core

import (
	"reflect"
)

func NewBaseDAO[T Entity](tableName string) *BaseDAO[T] {
	var zero T
	return &BaseDAO[T]{
		entityType: reflect.TypeOf(zero),
		tableName:  tableName,
	}
}

func (b BaseDAO[T]) GetEntityType() reflect.Type {
	return b.entityType
}

func (b BaseDAO[T]) GetTableName() string {
	return b.tableName
}
