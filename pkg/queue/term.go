package queue

import (
	"fmt"
	"strings"
)

var (
	operations = []string{"+", "-", "*", "/"}
)

type Cell struct {
	XValue int
	YValue string
}

func NewCell(expression string, header map[string]int) (*Cell, error) {
	for key, ind := range header {
		if strings.Contains(expression, key) && strings.Index(expression, key) == 0 {
			return &Cell{
				YValue: expression[len(key):],
				XValue: ind,
			}, nil
		}
	}

	return nil, &errNotFoundHeaderKey{value: expression}
}

func (c *Cell) String() string {
	return fmt.Sprintf("%v:%v", c.YValue, c.XValue)
}

func (c *Cell) PickValue(records map[string][]string) (string, error) {
	values, ok := records[c.YValue]
	if !ok {
		return "", &errNotFoundVerticalKey{value: c.YValue}
	}

	return values[c.XValue], nil
}

type Term struct {
	XKey      int
	YKey      string
	LeftCell  *Cell
	RightCell *Cell
	Operation string
}

func NewTerm(expression, yKey string, xKey int, header map[string]int) (*Term, error) {
	if !checkExpressionCorrectness(expression) {
		return nil, &errInvalidExpression{value: expression}
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
		return nil, &ErrUnknownOperationInExpression{Value: expression}
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

func checkExpressionCorrectness(expression string) bool {
	if strings.Index(expression, "=") != 0 {
		return false
	}

	return strings.Count(expression, "+")+
		strings.Count(expression, "-")+
		strings.Count(expression, "*")+
		strings.Count(expression, "/") == 1
}
