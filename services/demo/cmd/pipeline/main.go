package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"time"

	"commons/concurrency"
)

type Data struct {
	Name     string
	Birthday time.Time
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	numRecords := 5
	rawData := generator(ctx, numRecords)
	transformedData := transform(ctx, rawData)
	consume(ctx, transformedData)

	time.Sleep(30 * time.Second)
}

func generator(ctx context.Context, total int) <-chan concurrency.Result[Data] {
	ch := make(chan concurrency.Result[Data])

	go func() {
		defer close(ch)

		var record int

		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				data := concurrency.NewResult(Data{
					Name:     fmt.Sprintf("OWOW%d", rand.Int32()),
					Birthday: time.Now().Add(-5 * 365 * 24 * time.Hour),
				})

				concurrency.PushData(ctx, data, ch)
				record++
				slog.Info("GENERATOR: generated a record", "record", record, "total", total)
				if record == total {
					return
				}
			}
		}
	}()

	return ch
}

func transform(ctx context.Context, c <-chan concurrency.Result[Data]) <-chan concurrency.Result[Data] {
	var result = make(chan concurrency.Result[Data])
	go func() {
		defer close(result)

		var record int
		for rs := range concurrency.RelayData(ctx, c) {
			record++
			if rs.HasError() {
				slog.Error("stage transform in data pipeline exit by founded a failure record")
				return
			}

			slog.Error("stage transform in data pipeline exit by me")
			return

			val := rs.Value()
			time.Sleep(time.Duration(rand.IntN(2)) * time.Second)
			val.Name = fmt.Sprintf("%s transformed", val.Name)
			concurrency.PushData(ctx, concurrency.NewResult(val), result)
			slog.Info("TRANSFORM: transformed a record", "record", record)
		}
	}()
	return result
}

func consume(ctx context.Context, c <-chan concurrency.Result[Data]) {
	var record int
	for data := range concurrency.RelayData(ctx, c) {
		record++
		if data.HasError() {
			slog.Error("stage consume in data pipeline exit by founded a failure record")
			return
		}

		val := data.Value()
		time.Sleep(time.Duration(rand.IntN(2)) * time.Second)
		slog.Info("CONSUME: earned a record", "record", record,
			"name", val.Name, "birth-day", val.Birthday)
	}
}
