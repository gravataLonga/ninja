package evaluator

import (
	"ninja/ast"
	"ninja/object"
)

// prefix expression
func evalPostfixExpression(node *ast.PostfixExpression, env *object.Environment) object.Object {

	left := Eval(node.Left, env)
	if object.IsError(left) {
		return left
	}

	switch node.Operator {
	case "++":
		return evalPostfixExpressionAndAssing(node, left, env)
	case "--":
		return evalPostfixExpressionAndAssing(node, left, env)
	}
	return object.NewErrorFormat("unknown operator: %s%s", node.Operator, left.Type())
}

func evalPostfixExpressionAndAssing(node *ast.PostfixExpression, left object.Object, env *object.Environment) object.Object {
	astIdent, ok := node.Left.(*ast.Identifier)

	var result object.Object
	switch node.Operator {
	case "++":
		result = evalIncrementExpression(left)
	case "--":
		result = evalDecrementExpression(left)
	}

	// We are dealing with digits only..
	if !ok {
		return result
	}

	ident := astIdent.Token
	env.Set(ident.Literal, result)

	return left
}
