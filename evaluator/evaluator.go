package evaluator

import (
	"math"
	"ninja/ast"
	"ninja/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// Ast Program Eval
	case *ast.Program:
		return evalProgram(node.Statements, env)

		// ExpressionStatement
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.Identifier:
		return evalIdentifier(node, env)

		// VarStatement
	case *ast.VarStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)

		// PrefixExpression
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

		// InfixExpression
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)

		// IfExpression
	case *ast.IfExpression:
		return evalIfExpression(node, env)

		// FunctionsLiteral
	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &object.FunctionLiteral{Parameters: params, Env: env, Body: body}

	case *ast.Function:
		params := node.Parameters
		body := node.Body
		env.Set(node.Name.Value, &object.Function{Parameters: params, Env: env, Body: body})

		// CallFunctionNode
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)

		// BlockStatement
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

		// ReturnStatement
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

		// IntegerLiteral
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

		// FloatLiteral
	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}

		// Boolean
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

		// String
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

		// ArrayLiteral
	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}

		// IndexExpression for Array and Object
	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)

		// Hash
	case *ast.HashLiteral:
		return evalHashLiteral(node, env)
	case *ast.ForStatement:
		return evalForStatement(node, env)
	}
	return nil
}

func evalProgram(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement, env)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) object.Object {

	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("identifier not found: " + node.Value)
}

func evalExpressions(
	exps []ast.Expression,
	env *object.Environment,
) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

// @todo refactor
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
