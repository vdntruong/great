package postgre

import (
	"commons/db/dao/core"
	"context"
	"database/sql"
)

type PostgresDAO[T core.Entity] struct {
	*core.BaseDAO[T]
	db *sql.DB
}

func NewPostgresDAO[T core.Entity](db *sql.DB, tableName string) *PostgresDAO[T] {
	return &PostgresDAO[T]{
		BaseDAO: core.NewBaseDAO[T](tableName),
		db:      db,
	}
}

func (dao *PostgresDAO[T]) buildInsertQuery() string {
	qb := NewPostgresQueryBuilder[T](dao.GetTableName())
	columns := core.GetColumnInfo(dao.GetEntityType())
	return qb.BuildInsertQuery(columns)
}

func (dao *PostgresDAO[T]) getEntityValues(entity T) []interface{} {
	qb := NewPostgresQueryBuilder[T](dao.GetTableName())
	return qb.GetEntityValues(entity)
}

func (dao *PostgresDAO[T]) Create(ctx context.Context, entity T) error {
	if err := entity.Validate(); err != nil {
		return err
	}

	query := dao.buildInsertQuery()
	values := dao.getEntityValues(entity)

	result, err := dao.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	entity.SetID(id)
	return nil
}
