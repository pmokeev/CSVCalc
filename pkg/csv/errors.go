package csv

import (
	"errors"
	"fmt"
)

var (
	errDivisionByZero      = errors.New("division by zero")
	errEmptyValueInHeader  = errors.New("empty value in header cells")
	errUnequalLinesLengths = errors.New("unequal header and line lengths")
	errCyclicDependency    = errors.New("cyclic dependency of cells")
)

type errUnknownValueInLine struct {
	value string
}

func (e *errUnknownValueInLine) Error() string {
	return fmt.Sprintf("unknown value in line: %s", e.value)
}
