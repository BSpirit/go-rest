package core

import (
	"fmt"
	"runtime"
)

type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se *StatusError) Error() string {
	return se.Err.Error()
}

func (se *StatusError) Unwrap() error {
	return se.Err
}

func Trace(err error) error {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	return fmt.Errorf("%s line %d:\n\t%s", f.Name(), line, err)
}
