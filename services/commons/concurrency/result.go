package concurrency

type Result[T any] struct {
	err error
	val T
}

func NewResult[T any](value T) Result[T] {
	return Result[T]{
		val: value,
	}
}

func NewErrorResult[T any](err error) Result[T] {
	return Result[T]{
		err: err,
	}
}

func (r *Result[T]) HasError() bool {
	return r.err != nil
}

func (r *Result[T]) Value() T {
	return r.val
}
