package concurrency

import (
	"context"
	"sync"
)

func PushValueToChan[T any](ctx context.Context, ch chan<- T, v T) bool {
	if ch == nil {
		return false
	}

	select {
	case <-ctx.Done():
		return false
	case ch <- v:
	}

	return true
}

func PushValuesToChan[T any](ctx context.Context, ch chan<- T, v ...T) bool {
	if ch == nil {
		return false
	}

	for _, item := range v {
		select {
		case <-ctx.Done():
			return false
		case ch <- item:
		}
	}

	return true
}

func RelayDataFromChan[T any](ctx context.Context, ch <-chan T) <-chan T {
	return RelayDataFromChanWithBuffer(ctx, ch, 0)
}

func RelayDataFromChanWithBuffer[T any](ctx context.Context, ch <-chan T, size int) <-chan T {
	relayCh := make(chan T, size)
	go func() {
		defer close(relayCh)

		for {
			select {
			case <-ctx.Done():
				DrainChannel(ch)
				return
			case v, ok := <-ch:
				if !ok {
					return
				}

				select {
				case <-ctx.Done():
					DrainChannel(ch)
					return
				case relayCh <- v:
				}
			}
		}
	}()
	return relayCh
}

func DrainChannel[T any](ch <-chan T) {
	for range ch {
		// just drain
	}
}

// fan-in

func FanIn[T any](ctx context.Context, inputs ...<-chan T) <-chan T {
	return FanInWithBuffer[T](ctx, 0, inputs...)
}

func FanInWithBuffer[T any](ctx context.Context, size int, inputs ...<-chan T) <-chan T {
	var output = make(chan T, size)

	var wg sync.WaitGroup

	wg.Add(len(inputs))
	for _, in := range inputs {
		go func() {
			defer wg.Done()

			for v := range in {
				select {
				case <-ctx.Done():
					DrainChannel(in)
					return
				case output <- v:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
