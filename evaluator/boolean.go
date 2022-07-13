package evaluator

import "github.com/gravataLonga/ninja/object"

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return object.TRUE
	}
	return object.FALSE
}
