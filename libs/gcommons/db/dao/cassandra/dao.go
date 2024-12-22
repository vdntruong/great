package cassandra

import (
	"context"
	"gcommons/db/dao/core"
	"github.com/gocql/gocql"
)

type DAO[T core.Entity] struct {
	*core.BaseDAO[T]
	session *gocql.Session
}

func NewCassandraDAO[T core.Entity](session *gocql.Session, tableName string) *DAO[T] {
	return &DAO[T]{
		BaseDAO: core.NewBaseDAO[T](tableName),
		session: session,
	}
}

func (dao *DAO[T]) buildInsertQuery() string {
	qb := NewCassandraQueryBuilder[T](dao.GetTableName())
	columns := core.GetColumnInfo(dao.GetEntityType())
	return qb.BuildInsertQuery(columns)
}

func (dao *DAO[T]) getEntityValues(entity T) []interface{} {
	qb := NewCassandraQueryBuilder[T](dao.GetTableName())
	return qb.GetEntityValues(entity)
}

func (dao *DAO[T]) Create(ctx context.Context, entity T) error {
	if err := entity.Validate(); err != nil {
		return err
	}

	query := dao.buildInsertQuery()
	values := dao.getEntityValues(entity)

	return dao.session.Query(query, values...).WithContext(ctx).Exec()
}
