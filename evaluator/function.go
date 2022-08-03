package evaluator

import (
	"errors"
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func applyFunction(fn object.Object, args []object.Object) object.Object {

	switch fn := fn.(type) {
	case *object.FunctionLiteral:
		if err := argumentsIsValid(args, fn.Parameters); err != nil {
			return object.NewErrorFormat(err.Error()+" at %s", fn.Body.Token)
		}
		extendedEnv := extendFunctionEnv(fn.Env, fn.Parameters, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Function:
		if err := argumentsIsValid(args, fn.Parameters); err != nil {
			return object.NewErrorFormat(err.Error()+" at %s", fn.Body.Token)
		}
		extendedEnv := extendFunctionEnv(fn.Env, fn.Parameters, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return object.NewErrorFormat("not a function: %s", fn.Type())
	}

}

func extendFunctionEnv(
	fnEnv *object.Environment,
	fnArguments []ast.Expression,
	parameters []object.Object,
) *object.Environment {

	maxLen := len(parameters)

	env := object.NewEnclosedEnvironment(fnEnv)

	for argumentIndex, argument := range fnArguments {
		var value object.Object
		var identifier string

		switch argument.(type) {
		case *ast.Identifier:
			ident, _ := argument.(*ast.Identifier)
			value = parameters[argumentIndex]
			identifier = ident.Value
			break
		case *ast.InfixExpression:
			infix, _ := argument.(*ast.InfixExpression)
			ident, _ := infix.Left.(*ast.Identifier)
			identifier = ident.Value
			value = Eval(infix.Right, env)
			if maxLen > argumentIndex {
				value = parameters[argumentIndex]
			}
		}

		env.Set(identifier, value)
	}

	return env
}

// argumentsIsValid check if parameters passed to function is expected by arguments
func argumentsIsValid(parameters []object.Object, arguments []ast.Expression) error {
	if len(parameters) == len(arguments) {
		return nil
	}

	if len(parameters) > len(arguments) {
		return errors.New(fmt.Sprintf("Function expected %d arguments, got %d", len(arguments), len(parameters)))
	}

	// all arguments are infix expression, which mean, they have a default value
	total := 0
	for _, arg := range arguments {
		if _, ok := arg.(*ast.InfixExpression); ok {
			total++
		}
	}
	// all arguments have default value
	if total == len(arguments) {
		return nil
	}

	// a, b = 1
	if total+len(parameters) == len(arguments) {
		return nil
	}

	return errors.New(fmt.Sprintf("Function expected %d arguments, got %d", len(arguments), len(parameters)))
}
