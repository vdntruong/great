package concurrency

import "context"

type Stage string

// generator

type Generator[T any] func(context.Context) <-chan Result[T]

// pipeline process stage

type Processor[I any, O any] func(context.Context, I) (O, error)

type ResultProcessor[I any, O any] func(context.Context, Result[I]) (Result[O], bool)

type StageProcessor[I any, O any] func(context.Context, <-chan Result[I]) <-chan Result[O]

// pipeline consume stage

type Consumer[T any] func(context.Context, T) error

type ResultConsumer[T any] func(context.Context, Result[T]) error

type StageConsumer[T any] func(context.Context, <-chan Result[T]) error
