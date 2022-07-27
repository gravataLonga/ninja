package evaluator

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalProgram(node.Statements, env)

		// ExpressionStatement
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.DeleteStatement:
		return evalDelete(node.Left, Eval(node.Index, env), env)
	case *ast.Import:
		return evalImport(node, env)

	case *ast.Identifier:
		return evalIdentifier(node, env)

		// VarStatement
	case *ast.VarStatement:
		val := Eval(node.Value, env)
		if object.IsError(val) {
			return val
		}
		env.Set(node.Name.Value, val)

	case *ast.AssignStatement:
		return evalAssignStatement(node, env)

		// PrefixExpression
	case *ast.PrefixExpression:
		return evalPrefixExpression(node, env)

		// InfixExpression
	case *ast.InfixExpression:
		if node.Left == nil {
			return object.NewErrorFormat("InfixExpression.Left is nil.")
		}

		if node.Right == nil {
			return object.NewErrorFormat("InfixExpression.Right is nil.")
		}

		left := Eval(node.Left, env)
		if object.IsError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if object.IsError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)

	case *ast.PostfixExpression:
		return evalPostfixExpression(node, env)

	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.TernaryOperatorExpression:
		return evalTernaryOperatorExpression(node, env)
	case *ast.ElvisOperatorExpression:
		return evalElvisOperatorExpression(node, env)

		// FunctionsLiteral
	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &object.FunctionLiteral{Parameters: params, Env: env, Body: body}

	case *ast.Function:
		params := node.Parameters
		body := node.Body
		env.Set(node.Name.Value, &object.Function{Parameters: params, Env: env, Body: body})
		return &object.Function{Parameters: params, Env: env, Body: body}

		// CallFunctionNode
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if object.IsError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && object.IsError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)

		// BlockStatement
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

		// ReturnStatement
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if object.IsError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

	case *ast.BreakStatement:
		return &object.Break{Value: nil}

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
		if len(elements) == 1 && object.IsError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}

		// IndexExpression for Array and Object
	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if object.IsError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if object.IsError(index) {
			return index
		}
		return evalIndexExpression(left, index)

		// Hash
	case *ast.HashLiteral:
		return evalHashLiteral(node, env)
	case *ast.ForStatement:
		return evalForStatement(node, env)
	case *ast.ObjectCall:
		return evalObjectCallExpression(node, env)
	case *ast.EnumStatement:
		return evalEnumStatement(node, env)
	case *ast.ScopeOperatorExpression:
		return evalScopeOperatorExpression(node, env)
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
		case *object.Break:
			return object.NewErrorFormat("'break' not in the 'loop' context")
		case *object.Error:
			return result
		}
	}
	return result
}

func evalExpressions(
	exps []ast.Expression,
	env *object.Environment,
) []object.Object {
	var result []object.Object = []object.Object{}

	for _, e := range exps {
		evaluated := Eval(e, env)
		if object.IsError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}
