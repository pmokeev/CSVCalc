package queue

import (
	"fmt"
	"strings"
)

var (
	operations = []string{"+", "-", "*", "/"}
)

// Cell represents cell in csv table.
type Cell struct {
	XValue int
	YValue string
}

// NewCell returns new instance of Cell.
func NewCell(expression string, header map[string]int) (*Cell, error) {
	for key, ind := range header {
		if strings.Contains(expression, key) && strings.Index(expression, key) == 0 {
			return &Cell{
				YValue: expression[len(key):],
				XValue: ind,
			}, nil
		}
	}

	return nil, &notFoundHeaderKeyError{value: expression}
}

// String represents Stringer interface for Cell struct.
func (c *Cell) String() string {
	return fmt.Sprintf("%v:%v", c.YValue, c.XValue)
}

// PickValue picks value from records by cell values.
func (c *Cell) PickValue(records map[string][]string) (string, error) {
	values, ok := records[c.YValue]
	if !ok {
		return "", &notFoundVerticalKeyError{value: c.YValue}
	}

	return values[c.XValue], nil
}

// Term represents expression from csv table.
type Term struct {
	XKey      int
	YKey      string
	LeftCell  *Cell
	RightCell *Cell
	Operation string
}

// NewTerm returns new instance of Term.
func NewTerm(expression, yKey string, xKey int, header map[string]int) (*Term, error) {
	if !checkExpressionCorrectness(expression) {
		return nil, &invalidExpressionError{value: expression}
	}

	term := &Term{
		XKey: xKey,
		YKey: yKey,
	}

	var operation int
	for _, op := range operations {
		if strings.Contains(expression, op) {
			operation = strings.Index(expression, op)
			term.Operation = op
			break
		}
	}

	if operation == -1 {
		return nil, &UnknownOperationInExpressionError{Value: expression}
	}

	leftCell, err := NewCell(expression[1:operation], header)
	if err != nil {
		return nil, err
	}
	rightCell, err := NewCell(expression[operation+1:], header)
	if err != nil {
		return nil, err
	}

	term.LeftCell = leftCell
	term.RightCell = rightCell

	return term, nil
}

// checkExpressionCorrectness checks term expression on correctness.
func checkExpressionCorrectness(expression string) bool {
	if strings.Index(expression, "=") != 0 {
		return false
	}

	return strings.Count(expression, "+")+
		strings.Count(expression, "-")+
		strings.Count(expression, "*")+
		strings.Count(expression, "/") == 1
}
