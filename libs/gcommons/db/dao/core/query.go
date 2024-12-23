package core

import (
	"reflect"
	"strings"
)

type Filter struct {
	Field string
	Op    string
	Value interface{}
}

type Sort struct {
	Field string
	Desc  bool
}

type Pagination struct {
	Page     int
	PageSize int
}

type QueryOptions struct {
	Filters    []Filter
	Sort       []Sort
	Pagination *Pagination
}

type ColumnInfo struct {
	Name      string
	FieldName string
	Type      reflect.Type
	Tag       reflect.StructTag
}

// GetColumnInfo extracts column information from struct tags
func GetColumnInfo(t reflect.Type) []ColumnInfo {
	var columns []ColumnInfo
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get the db tag, fallback to field name in snake_case
		dbTag := field.Tag.Get("db")
		if dbTag == "-" {
			continue
		}
		if dbTag == "" {
			dbTag = toSnakeCase(field.Name)
		}

		columns = append(columns, ColumnInfo{
			Name:      dbTag,
			FieldName: field.Name,
			Type:      field.Type,
			Tag:       field.Tag,
		})
	}
	return columns
}

func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(toLower(r))
	}
	return result.String()
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}
