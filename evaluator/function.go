package evaluator

import (
	"ninja/ast"
	"ninja/object"
)

func applyFunction(fn object.Object, args []object.Object) object.Object {

	switch fn := fn.(type) {
	case *object.FunctionLiteral:
		extendedEnv := extendFunctionEnv(fn.Env, fn.Parameters, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn.Env, fn.Parameters, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError("not a function: %s", fn.Type())
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
