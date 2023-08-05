package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitIfExpr(v *ast.IfExpression) (result object.Object) {
	// Probably problem is here:
	condition := i.evaluate(v.Condition)
	if object.IsTruthy(condition) {
		return i.execute(v.Consequence)
	}
	if v.Alternative != nil {
		return i.execute(v.Alternative)
	}
	return object.NULL
}

func (i *Interpreter) VisitTernaryOperator(v *ast.TernaryOperatorExpression) (result object.Object) {
	condition := i.evaluate(v.Condition)
	if object.IsTruthy(condition) {
		return i.evaluate(v.Consequence)
	}

	return i.evaluate(v.Alternative)
}

func (i *Interpreter) VisitElvisOperator(v *ast.ElvisOperatorExpression) (result object.Object) {
	left := i.evaluate(v.Left)
	if object.IsTruthy(left) {
		return left
	}
	return i.evaluate(v.Right)
}
