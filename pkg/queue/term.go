package queue

import (
	"errors"
	"strings"
)

var (
	operations = []string{"+", "-", "*", "/"}
)

type Cell struct {
	XValue string
	YValue string
}

func NewCell(expression string, header map[string]int) (*Cell, error) {
	for key := range header {
		if strings.Contains(expression, key) && strings.Index(expression, key) == 0 {
			return &Cell{
				XValue: expression[len(key):],
				YValue: expression[:len(key)],
			}, nil
		}
	}

	return nil, errors.New("expression doesn't contains header key")
}

func (c *Cell) PickValue(records map[string][]string, header map[string]int) (string, error) {
	values, ok := records[c.YValue]
	if !ok {
		return "", errors.New("vertical key doesn't exist")
	}

	index := header[c.XValue]

	return values[index], nil
}

type Term struct {
	XKey      string
	YKey      int
	LeftCell  *Cell
	RightCell *Cell
	Operation string
}

func NewTerm(expression, xKey string, yKey int, header map[string]int) (*Term, error) {
	if !checkExpressionCorrectness(expression) {
		return nil, errors.New("expression isn't correct")
	}

	term := &Term{
		XKey: xKey,
		YKey: yKey,
	}

	var operation int
	for _, op := range operations {
		if strings.Contains(expression, op) {
			operation = strings.Index(expression, "+")
			term.Operation = op
			break
		}
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
	if expression[0] != '=' {
		return false
	}

	return strings.Count(expression, "+")+
		strings.Count(expression, "-")+
		strings.Count(expression, "*")+
		strings.Count(expression, "/") == 1
}