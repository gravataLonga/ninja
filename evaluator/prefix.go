package evaluator

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

// prefix expression
func evalPrefixExpression(node *ast.PrefixExpression, env *object.Environment) object.Object {

	right := Eval(node.Right, env)
	if object.IsError(right) {
		return right
	}

	switch node.Operator {
	case "++":
		return evalPrefixExpressionAndAssing(node, right, env)
	case "--":
		return evalPrefixExpressionAndAssing(node, right, env)
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(node, right)
	}

	return object.NewErrorFormat("unknown operator: %s%s %s", node.Operator, right.Type(), node.Token)
}

func evalPrefixExpressionAndAssing(node *ast.PrefixExpression, right object.Object, env *object.Environment) object.Object {
	astIdent, ok := node.Right.(*ast.Identifier)

	var result object.Object
	switch node.Operator {
	case "++":
		result = evalIncrementExpression(right)
	case "--":
		result = evalDecrementExpression(right)
	}

	// We are dealing with digits only..
	if !ok {
		return result
	}

	ident := astIdent.Token
	env.Set(ident.Literal, result)

	return result
}
