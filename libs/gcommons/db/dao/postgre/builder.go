package postgre

import (
	"fmt"
	"reflect"
	"strings"

	"gcommons/db/dao/core"
)

type PostgresQueryBuilder[T core.Entity] struct {
	*core.BaseDAO[T]
	tableName string
}

func NewPostgresQueryBuilder[T core.Entity](tableName string) *PostgresQueryBuilder[T] {
	return &PostgresQueryBuilder[T]{
		BaseDAO:   core.NewBaseDAO[T](tableName),
		tableName: tableName,
	}
}

func (qb *PostgresQueryBuilder[T]) BuildInsertQuery(columns []core.ColumnInfo) string {
	columnNames := make([]string, 0, len(columns))
	placeholders := make([]string, 0, len(columns))

	for i, col := range columns {
		if col.Name == "id" {
			continue // Skip ID for auto-increment
		}
		columnNames = append(columnNames, col.Name)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING id",
		qb.tableName,
		strings.Join(columnNames, ", "),
		strings.Join(placeholders, ", "),
	)

	return query
}

func (qb *PostgresQueryBuilder[T]) BuildUpdateQuery(columns []core.ColumnInfo) string {
	setStatements := make([]string, 0, len(columns))

	for i, col := range columns {
		if col.Name == "id" {
			continue
		}
		setStatements = append(setStatements, fmt.Sprintf("%s = $%d", col.Name, i+1))
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = $%d",
		qb.tableName,
		strings.Join(setStatements, ", "),
		len(columns),
	)

	return query
}

func (qb *PostgresQueryBuilder[T]) GetEntityValues(entity T) []interface{} {
	val := reflect.ValueOf(entity)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	columns := core.GetColumnInfo(val.Type())
	values := make([]interface{}, 0, len(columns))

	for _, col := range columns {
		if col.Name == "id" {
			continue
		}
		values = append(values, val.FieldByName(col.FieldName).Interface())
	}

	return values
}
