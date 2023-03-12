package queue

import "fmt"

// errNotFoundHeaderKey represents NotFoundHeaderKey error.
type errNotFoundHeaderKey struct {
	value string
}

// Error represents Error interface.
func (e *errNotFoundHeaderKey) Error() string {
	return fmt.Sprintf("expression doesn't contains header key: %v", e.value)
}

// errNotFoundVerticalKey represents NotFoundVerticalKey error.
type errNotFoundVerticalKey struct {
	value string
}

// Error represents Error interface.
func (e *errNotFoundVerticalKey) Error() string {
	return fmt.Sprintf("not found vertical key: %v", e.value)
}

// errInvalidExpression represents InvalidExpression error.
type errInvalidExpression struct {
	value string
}

// Error represents Error interface.
func (e *errInvalidExpression) Error() string {
	return fmt.Sprintf("invalid expression: %v", e.value)
}

// ErrUnknownOperationInExpression represents UnknownOperation error.
type ErrUnknownOperationInExpression struct {
	Value string
}

// Error represents Error interface.
func (e *ErrUnknownOperationInExpression) Error() string {
	return fmt.Sprintf("unknown operation in expression: %s", e.Value)
}
