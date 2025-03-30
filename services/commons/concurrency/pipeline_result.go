package concurrency

import (
	"fmt"
)

type Result[T any] struct {
	err error
	val T
}

func NewResult[T any](value T) Result[T] {
	return Result[T]{
		val: value,
	}
}

func NewResultWithError[T any](value T, err error) Result[T] {
	return Result[T]{
		val: value,
		err: err,
	}
}

func (r *Result[T]) HasError() bool {
	return r.err != nil
}

func (r *Result[T]) Value() T {
	return r.val
}

func (r *Result[T]) Error() string {
	return r.err.Error()
}

func (r *Result[T]) Unwrap() error {
	return r.err
}

type PipelineError struct {
	Stage   Stage
	Message string
	Err     error
}

func (e *PipelineError) Error() string {
	return fmt.Sprintf("[%s] %s: %v", e.Stage, e.Message, e.Err)
}

func (e *PipelineError) Unwrap() error {
	return e.Err
}
