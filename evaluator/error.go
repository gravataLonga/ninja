package evaluator

import (
	"fmt"
	"ninja/object"
)

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
