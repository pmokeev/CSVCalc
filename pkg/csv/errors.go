package csv

import (
	"errors"
	"fmt"
)

var (
	errDivisionByZero      = errors.New("division by zero")
	errEmptyValueInHeader  = errors.New("empty value in header cells")
	errUnequalLinesLengths = errors.New("unequal header and line lengths")
)

// errUnknownValueInLine represents UnknownValue error.
type errUnknownValueInLine struct {
	value string
}

// Error represents Error interface.
func (e *errUnknownValueInLine) Error() string {
	return fmt.Sprintf("unknown value in line: %s", e.value)
}
