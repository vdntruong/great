package errs

import (
	"fmt"
	"path/filepath"
	"runtime"
)

const maxTraceback = 10

func trace() string {
	pc := make([]uintptr, maxTraceback)
	skip := 3
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("\t%s:%d | %s", filepath.Base(frame.File), frame.Line, frame.Function)
}

func ErrWithTrace(err error) error {
	return fmt.Errorf("%w \n at %s", err, trace())
}
