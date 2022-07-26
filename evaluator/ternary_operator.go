package evaluator

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func evalTernaryOperatorExpression(
	to *ast.TernaryOperatorExpression,
	env *object.Environment,
) object.Object {
	condition := Eval(to.Condition, env)

	if object.IsError(condition) {
		return condition
	}

	if object.IsTruthy(condition) {
		return Eval(to.Consequence, env)
	}
	return Eval(to.Alternative, env)
}
