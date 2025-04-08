package main

import (
	"context"
	"fmt"
	"time"

	"commons/concurrency"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	pipeline(ctx)
}

func pipeline(ctx context.Context) {
	var (
		totalTask = 3000
		maxRandom = 50
	)

	numberTasks := concurrency.IntegerTaskGenerator(totalTask, maxRandom)(ctx)
	odd := concurrency.ProcessFilterOddNumbersStage(ctx, numberTasks)
	square := concurrency.ProcessSquareNumbersStage(ctx, odd)
	half := concurrency.ProcessHalfNumbersStage(ctx, square)

	if err := concurrency.ProcessPrintNumbersStage(ctx, half); err != nil {
		fmt.Println("ERROR:", err)
	}
}

func enhancedPipeline(ctx context.Context) {
	pp := concurrency.NewEnhancedPipeline[concurrency.Task, concurrency.Task]()

	handler := pp.ErrorHandler.(*concurrency.DefaultErrorHandler)
	handler.MaxErrors = 50
	handler.RegisterCriticalErrorType("*MyCustomCriticalError")

	// Create and run pp
	var (
		totalTask = 3000
		maxRandom = 50
	)

	numberTasks := concurrency.IntegerTaskGenerator(totalTask, maxRandom)(ctx)
	var oddstage concurrency.Stage = "ODD Stage"
	odd := pp.ProcessStage(ctx, oddstage, numberTasks, concurrency.ProcessCheckOddNumber)
	square := pipeline.ProcessSquareNumbersStage(childCtx, odd)
	half := pipeline.ProcessHalfNumbersStage(childCtx, square)

	err := pipeline.ProcessPrintNumbersStage(childCtx, half)
}
