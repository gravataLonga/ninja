package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func (i *Interpreter) VisitFuncExpr(v *ast.FunctionLiteral) (result object.Object) {
	fn := &object.FunctionLiteral{Parameters: v.Parameters, Body: v.Body}
	if v.Name != nil {
		i.env.Set(v.Name.Value, fn)
	}
	return fn
}

// VisitCallExpr
// @todo arguments must be it's on AST structure.
func (i *Interpreter) VisitCallExpr(v *ast.CallExpression) (result object.Object) {
	obj := i.evaluate(v.Function)

	if object.IsError(obj) {
		return obj
	}

	switch obj.Type() {
	case object.FUNCTION_OBJ:
		return i.applyFunction(obj, v)
	case object.BUILTIN_OBJ:
		var args []object.Object

		for _, e := range v.Arguments {
			evaluated := i.evaluate(e)
			if object.IsError(evaluated) {
				// return []object.Object{evaluated}
			}
			args = append(args, evaluated)
		}

		return obj.(*object.Builtin).Fn(args...)
	default:
		return object.NewErrorFormat("Not implement yet VisitCallExpr")
	}
}

func (i *Interpreter) applyFunction(obj object.Object, v *ast.CallExpression) (result object.Object) {
	fn, _ := obj.(*object.FunctionLiteral)
	parameters := fn.Parameters.([]ast.Expression)

	mParameter := len(v.Arguments)
	err := i.validateArguments(v, parameters)

	if object.IsError(err) {
		return err
	}

	envLocal := object.NewEnclosedEnvironment(i.env)

	for index, parameter := range parameters {
		ident, ok := parameter.(*ast.Identifier)
		if ok {
			envLocal.Set(ident.String(), i.evaluate(v.Arguments[index]))
			continue
		}

		infix, ok := parameter.(*ast.InfixExpression)
		ident, _ = infix.Left.(*ast.Identifier)

		var value object.Object
		if mParameter > index {
			argument := v.Arguments[index]
			value = i.evaluate(argument)
		} else {
			value = i.evaluate(infix.Right)
		}

		envLocal.Set(ident.String(), value)
	}

	env := i.env
	i.env = envLocal
	result = i.VisitBlock(fn.Body.(*ast.BlockStatement))
	i.env = env

	if object.IsReturn(result) {
		return result.(*object.ReturnValue).Value
	}
	return
}

func (i *Interpreter) validateArguments(v *ast.CallExpression, parameters []ast.Expression) object.Object {
	// cases:
	// 1. fn () {}(1);
	// 2. fn (x) {}();
	// 3. fn (x) {}(1, 2);
	// 4. fn (x, y = 1) {}();
	// 5. fn (x, y = 1) {}(1,2,3);

	/*
		A parameter is a variable in a function definition. It is a placeholder and hence does not have a concrete value.
		An argument is a value passed during function invocation.
	*/

	totalParameters := len(parameters)
	totalParametersDefault := 0
	totalParametersRequireBeforeDefault := 0
	totalArguments := len(v.Arguments)

	for _, p := range parameters {
		if _, ok := p.(*ast.InfixExpression); ok {
			totalParametersDefault++
		}
		if _, ok := p.(*ast.Identifier); ok {
			totalParametersRequireBeforeDefault++
		}
	}

	if totalParametersDefault > 0 {
		if totalParametersDefault != totalParameters-totalParametersRequireBeforeDefault {
			return object.NewErrorFormat("Function expected %d parameters, got %d at %s", totalParameters, totalArguments, v.Token)
		}
		if totalParametersDefault+totalParametersRequireBeforeDefault != totalParameters {
			return object.NewErrorFormat("Function expected %d parameters, got %d at %s", totalParametersDefault+totalParametersRequireBeforeDefault, totalArguments, v.Token)
		}

		return nil
	}

	if totalArguments != totalParameters {
		return object.NewErrorFormat("Function expected %d parameters, got %d at %s", totalParameters, totalArguments, v.Token)
	}

	return nil
}

func (i *Interpreter) VisitDotExpr(v *ast.Dot) (result object.Object) {
	obj := i.evaluate(v.Object)
	call, ok := v.Right.(*ast.CallExpression)
	if !ok {
		return object.NewErrorFormat("we expect to be a call on right of dot operation. Got: %t", v.Right)
	}

	objCallable, ok := obj.(object.CallableMethod)
	if !ok {
		return object.NewErrorFormat("object must implement callable.")
	}

	method, ok := call.Function.(*ast.Identifier)
	if !ok {
		return object.NewErrorFormat("method name isn't a identifier")
	}

	args := i.evaluateExpressions(call.Arguments)

	return objCallable.Call(method.Value, args...)

	/*
		function := Eval(node.Function, env)
		if object.IsError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && object.IsError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)

		return object.NewErrorFormat("Not implement yet VisitDotExpr")
	*/
	return object.NewErrorFormat("Not implement yet VisitDotExpr")
}
