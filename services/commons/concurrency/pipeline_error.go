package concurrency

import (
	"context"
	"fmt"
	"sync"
)

type ErrorSeverity int

const (
	ErrorSeverityInfo     ErrorSeverity = iota + 1 // Log and continue
	ErrorSeverityWarning                           // Log, report, and continue
	ErrorSeverityCritical                          // Stop processing
)

type ErrorHandler interface {
	HandleError(context.Context, *PipelineError) ErrorSeverity
}

type DefaultErrorHandler struct {
	MaxErrors          int
	ErrorCount         int
	CriticalErrorTypes map[string]bool
	mu                 sync.Mutex
}

func NewDefaultErrorHandler() *DefaultErrorHandler {
	return &DefaultErrorHandler{
		MaxErrors:          50,
		CriticalErrorTypes: make(map[string]bool),
	}
}

func (h *DefaultErrorHandler) RegisterCriticalErrorType(errorType string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.CriticalErrorTypes[errorType] = true
}

func (h *DefaultErrorHandler) HandleError(ctx context.Context, err *PipelineError) ErrorSeverity {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.ErrorCount++

	errType := fmt.Sprintf("%T", err.Err)

	if h.CriticalErrorTypes[errType] {
		fmt.Printf("CRITICAL ERROR: %v\n", err)
		return ErrorSeverityCritical
	}

	if h.ErrorCount > h.MaxErrors {
		fmt.Printf("ERROR THRESHOLD EXCEEDED: %v\n", err)
		return ErrorSeverityCritical
	}

	fmt.Printf("ERROR: %v\n", err)
	return ErrorSeverityWarning
}
