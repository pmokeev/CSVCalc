package csv

import (
	"errors"
	"fmt"
)

var (
	errDivisionByZero      = errors.New("division by zero")
	errEmptyValueInHeader  = errors.New("empty value in header cells")
	errUnequalLinesLengths = errors.New("unequal header and line lengths")
	errEmptyHeader         = errors.New("empty header of csv file")
)

// unknownValueInLine represents UnknownValue error.
type unknownValueInLineError struct {
	value string
}

// Error represents Error interface.
func (e *unknownValueInLineError) Error() string {
	return fmt.Sprintf("unknown value in line: %s", e.value)
}
