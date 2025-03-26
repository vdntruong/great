package concurrency

import "context"

func SliceToChan[T any](ctx context.Context, values []T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)

		for i := 0; i < len(values); i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- values[i]:
			}
		}
	}()
	return ch
}
