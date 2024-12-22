package core

import "errors"

var (
	ErrNotFound         = errors.New("entity not found")
	ErrInvalidEntity    = errors.New("invalid entity")
	ErrConnectionFailed = errors.New("database connection failed")
	ErrOperationFailed  = errors.New("operation failed")
)
