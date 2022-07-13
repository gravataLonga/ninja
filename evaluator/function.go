package evaluator

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func applyFunction(fn object.Object, args []object.Object) object.Object {

	switch fn := fn.(type) {
	case *object.FunctionLiteral:
		if len(fn.Parameters) != len(args) {
			return object.NewErrorFormat("Function expected %d arguments, got %d at %s", len(fn.Parameters), len(args), fn.Body.Token)
		}
		extendedEnv := extendFunctionEnv(fn.Env, fn.Parameters, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Function:
		if len(fn.Parameters) != len(args) {
			return object.NewErrorFormat("Function expected %d arguments, got %d at %s", len(fn.Parameters), len(args), fn.Body.Token)
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
	fnArguments []*ast.Identifier,
	args []object.Object,
) *object.Environment {

	env := object.NewEnclosedEnvironment(fnEnv)

	for paramIdx, param := range fnArguments {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}
