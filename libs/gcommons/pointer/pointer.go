package pointer

func FromValue[T any](value T) *T {
	return &value
}
