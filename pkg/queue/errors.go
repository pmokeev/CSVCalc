package queue

import "fmt"

type errNotFoundHeaderKey struct {
	value string
}

func (e *errNotFoundHeaderKey) Error() string {
	return fmt.Sprintf("expression doesn't contains header key: %v", e.value)
}

type errNotFoundVerticalKey struct {
	value string
}

func (e *errNotFoundVerticalKey) Error() string {
	return fmt.Sprintf("not found vertical key: %v", e.value)
}

type errInvalidExpression struct {
	value string
}

func (e *errInvalidExpression) Error() string {
	return fmt.Sprintf("invalid expression: %v", e.value)
}

type ErrUnknownOperationInExpression struct {
	Value string
}

func (e *ErrUnknownOperationInExpression) Error() string {
	return fmt.Sprintf("unknown operation in expression: %s", e.Value)
}
