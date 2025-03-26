package concurrency

import "context"

func RelayData[T any](ctx context.Context, ch <-chan T) <-chan T {
	relayCh := make(chan T)
	go func() {
		defer close(relayCh)

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch:
				if !ok {
					return
				}

				select {
				case <-ctx.Done():
					return
				case relayCh <- v:
				}
			}
		}
	}()
	return relayCh
}

func PushData[T any](ctx context.Context, v T, c chan<- T) {
	select {
	case <-ctx.Done():
		return
	case c <- v:
	}
}
