package concurrency

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
)

// examples

const (
	prefixGenNumber    = "1-GENERATE-NUMBER-STAGE"
	prefixOddNumber    = "2-ODD-NUMBER-STAGE"
	prefixSquareNumber = "3-SQUARE-NUMBER-STAGE"
	prefixHalfNumber   = "4-HALF-NUMBER-STAGE"
	prefixPrintNumber  = "5-PRINT-NUMBER-STAGE"
)

type Task struct {
	Name  string
	Value int
}

func IntegerTaskGenerator(total int, max int) func(context.Context) <-chan Result[Task] {
	return func(ctx context.Context) <-chan Result[Task] {
		ch := make(chan Result[Task])
		go func() {
			defer close(ch)

			for i := 1; i <= total; i++ {
				select {
				case <-ctx.Done():
					return
				default:
					var (
						name = fmt.Sprintf("task id T-%d", i)
						task = NewResult(Task{Name: name, Value: random(ctx, max)})
					)

					if !PushValueToChan[Result[Task]](ctx, ch, task) {
						fmt.Printf("ERROR Failed to generate new task %v\n", task.Value().Name)
					}
				}
			}
		}()
		return ch
	}
}

// Odd stage

func checkOddNumber(ctx context.Context, n int) (int, error) {
	_ = ctx
	if isOdd(n) {
		return n, nil
	}
	return 0, fmt.Errorf("%d is not an odd number", n)
}

func ProcessCheckOddNumber(ctx context.Context, r Result[Task]) (Result[Task], bool) {
	_ = ctx

	// 1 - check previous stage err
	if r.HasError() {
		// this stage don't want to keep processing the failed record so mask it not processed and handle the error
		fmt.Println("err: ", r.Error())
		return Result[Task]{}, false
	}

	// 2 - extract actual job
	var task = r.Value()

	// 3 - process job
	oddNumber, err := checkOddNumber(ctx, task.Value)
	if err != nil { // 4 - if fail to process job then return result with an error
		var stageErr = fmt.Errorf("%s failed: %w", prefixOddNumber, err)
		var result = NewResultWithError(Task{}, stageErr)
		return result, true
	}

	// 5 - if success then return new value (this example is the task)
	var result = NewResult(Task{Name: task.Name, Value: oddNumber})
	return result, true
}

func ProcessFilterOddNumbersStage(ctx context.Context, in <-chan Result[Task]) <-chan Result[Task] {
	out := make(chan Result[Task])
	go func() {
		defer close(out)

		for tr := range RelayDataFromChan(ctx, in) {
			result, ok := ProcessCheckOddNumber(ctx, tr)
			if !ok {
				continue
			}

			if !PushValueToChan(ctx, out, result) {
				fmt.Printf("%s failed to push task to next stage %+v\n", prefixOddNumber, result.Value())
			}
		}
	}()
	return out
}

// Square stage

func squareNumber(ctx context.Context, n int) (int, error) {
	_ = ctx
	return int(math.Pow(float64(n), 2)), nil
}

func processSquareNumber(ctx context.Context, r Result[Task]) (Result[Task], bool) {
	_ = ctx

	// 1 - check previous stage err
	if r.HasError() {
		// this stage don't want to keep processing the failed record so mask it not processed and handle the error
		fmt.Println("err: ", r.Error())
		return Result[Task]{}, false
	}

	// 2 - extract actual job
	var task = r.Value()

	// 3 - process job
	sqNumber, err := squareNumber(ctx, task.Value)
	if err != nil {
		var stageErr = fmt.Errorf("%s failed: %w", prefixSquareNumber, err)
		var result = NewResultWithError(Task{}, stageErr)
		return result, true
	}

	var result = NewResult(Task{Name: task.Name, Value: sqNumber})
	return result, true
}

func ProcessSquareNumbersStage(ctx context.Context, in <-chan Result[Task]) <-chan Result[Task] {
	out := make(chan Result[Task])
	go func() {
		defer close(out)

		for tr := range RelayDataFromChan(ctx, in) {
			result, ok := processSquareNumber(ctx, tr)
			if !ok {
				continue
			}

			if !PushValueToChan(ctx, out, result) {
				fmt.Printf("%s failed to push task to next stage %+v\n", prefixSquareNumber, result.Value())
			}
		}
	}()
	return out
}

// Half stage

func halfNumber(ctx context.Context, n int) (int, error) {
	_ = ctx
	return n / 2, nil
}

func processHalfNumber(ctx context.Context, r Result[Task]) (Result[Task], bool) {
	_ = ctx

	// 1 - check previous stage err
	if r.HasError() {
		// this stage don't want to keep processing the failed record so mask it not processed and handle the error
		fmt.Println("err: ", r.Error())
		return Result[Task]{}, false
	}

	// 2 - extract actual job
	var task = r.Value()

	// 3 - process job
	halfNum, err := halfNumber(ctx, task.Value)
	if err != nil {
		var stageErr = fmt.Errorf("%s failed: %w", prefixHalfNumber, err)
		var result = NewResultWithError(Task{}, stageErr)
		return result, true
	}

	var result = NewResult(Task{Name: task.Name, Value: halfNum})
	return result, true
}

func ProcessHalfNumbersStage(ctx context.Context, in <-chan Result[Task]) <-chan Result[Task] {
	out := make(chan Result[Task])
	go func() {
		defer close(out)

		for tr := range RelayDataFromChan(ctx, in) {
			result, ok := processHalfNumber(ctx, tr)
			if !ok {
				continue
			}

			if !PushValueToChan(ctx, out, result) {
				fmt.Printf("%s failed to push task to next stage %+v\n", prefixHalfNumber, result.Value())
			}
		}
	}()
	return out
}

// Print stage

func printNumber(ctx context.Context, n int) error {
	_ = ctx
	fmt.Println("======> VALUE: ", n)
	return nil
}

func processPrintNumber(ctx context.Context, r Result[Task]) error {
	_ = ctx

	// 1 - check previous stage err
	if r.HasError() {
		// this stage don't want to keep processing the failed record so mask it not processed and handle the error
		fmt.Println("err: ", r.Error())
		return nil
	}

	// 2 - extract actual job
	var task = r.Value()

	// 3 - consume job
	return printNumber(ctx, task.Value)
}

func ProcessPrintNumbersStage(ctx context.Context, in <-chan Result[Task]) error {
	for tr := range RelayDataFromChan(ctx, in) {
		if err := processPrintNumber(ctx, tr); err != nil {
			return err
		}
	}
	return nil
}

func isOdd(i int) bool {
	return i%2 == 0
}

func isEven(i int) bool {
	return i%2 == 1
}

func random(ctx context.Context, max int) int {
	select {
	case <-ctx.Done():
		return 0
	default:
		return rand.IntN(max)
	}
}
