package evaluator

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func evalObjectCallExpression(node *ast.ObjectCall, env *object.Environment) object.Object {
	// @todo check if is object.Object and if isnt error.
	obj := Eval(node.Object, env)

	callExpression, ok := node.Call.(*ast.CallExpression)
	if !ok {
		return object.NewErrorFormat("object.call is not call expression. Got: %s", callExpression)
	}

	method, ok := callExpression.Function.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("object.call.function isn't a identifier. Got: %s", callExpression.Function)
	}

	callable, ok := obj.(object.CallableMethod)

	if !ok {
		return object.NewErrorFormat("object.call.function isn't callable. Got: %T", obj)
	}

	args := evalExpressions(callExpression.Arguments, env)
	if len(args) == 1 && object.IsError(args[0]) {
		return args[0]
	}

	return callable.Call(method.Value, args...)
}
