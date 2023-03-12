package queue

import "fmt"

// notFoundHeaderKeyError represents NotFoundHeaderKey error.
type notFoundHeaderKeyError struct {
	value string
}

// Error represents Error interface.
func (e *notFoundHeaderKeyError) Error() string {
	return fmt.Sprintf("expression doesn't contains header key: %v", e.value)
}

// notFoundVerticalKeyError represents NotFoundVerticalKey error.
type notFoundVerticalKeyError struct {
	value string
}

// Error represents Error interface.
func (e *notFoundVerticalKeyError) Error() string {
	return fmt.Sprintf("not found vertical key: %v", e.value)
}

// invalidExpressionError represents InvalidExpression error.
type invalidExpressionError struct {
	value string
}

// Error represents Error interface.
func (e *invalidExpressionError) Error() string {
	return fmt.Sprintf("invalid expression: %v", e.value)
}

// UnknownOperationInExpressionError represents UnknownOperation error.
type UnknownOperationInExpressionError struct {
	Value string
}

// Error represents Error interface.
func (e *UnknownOperationInExpressionError) Error() string {
	return fmt.Sprintf("unknown operation in expression: %s", e.Value)
}
