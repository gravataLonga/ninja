package evaluator

import (
	"ninja/ast"
	"ninja/object"
)

func evalIfExpression(
	ie *ast.IfExpression,
	env *object.Environment,
) object.Object {
	condition := Eval(ie.Condition, env)
	if object.IsError(condition) {
		return condition
	}

	if object.IsTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return object.NULL
	}
}
