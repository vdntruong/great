package cassandra

import (
	"commons/db/dao/core"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

type QueryBuilder[T core.Entity] struct {
	*core.BaseDAO[T]
	tableName string
}

func NewCassandraQueryBuilder[T core.Entity](tableName string) *QueryBuilder[T] {
	return &QueryBuilder[T]{
		BaseDAO:   core.NewBaseDAO[T](tableName),
		tableName: tableName,
	}
}

func (qb *QueryBuilder[T]) BuildInsertQuery(columns []core.ColumnInfo) string {
	columnNames := make([]string, 0, len(columns))
	placeholders := make([]string, 0, len(columns))

	for range columns {
		placeholders = append(placeholders, "?")
	}

	for _, col := range columns {
		columnNames = append(columnNames, col.Name)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		qb.tableName,
		strings.Join(columnNames, ", "),
		strings.Join(placeholders, ", "),
	)

	return query
}

func (qb *QueryBuilder[T]) BuildUpdateQuery(columns []core.ColumnInfo, whereColumns []string) string {
	setStatements := make([]string, 0, len(columns))
	whereStatements := make([]string, 0, len(whereColumns))

	for _, col := range columns {
		if slices.Contains(whereColumns, col.Name) {
			continue
		}
		setStatements = append(setStatements, fmt.Sprintf("%s = ?", col.Name))
	}

	for _, col := range whereColumns {
		whereStatements = append(whereStatements, fmt.Sprintf("%s = ?", col))
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		qb.tableName,
		strings.Join(setStatements, ", "),
		strings.Join(whereStatements, " AND "),
	)

	return query
}

func (qb *QueryBuilder[T]) GetEntityValues(entity T) []interface{} {
	val := reflect.ValueOf(entity)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	columns := core.GetColumnInfo(val.Type())
	values := make([]interface{}, 0, len(columns))

	for _, col := range columns {
		values = append(values, val.FieldByName(col.FieldName).Interface())
	}

	return values
}
