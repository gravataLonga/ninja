package interpreter

import (
	"errors"
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
	"io"
)

type Interpreter struct {
	env     *object.Environment
	globals *object.Environment
	locals  map[ast.Expression]int
	output  io.Writer
}

func New(w io.Writer) *Interpreter {
	return &Interpreter{
		env:     object.NewEnvironment(),
		globals: object.NewEnvironment(),
		locals:  make(map[ast.Expression]int),
	}
}

func (i *Interpreter) evaluate(node ast.Expression) object.Object {
	result := node.Accept(i)
	if _, ok := result.(object.Object); !ok {
		return nil
	}
	return result.(object.Object)
}

func (i *Interpreter) Interpreter(node ast.Node) object.Object {
	return i.execute(node)
}

func (i *Interpreter) execute(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		result := node.Accept(i)
		if result == nil {
			return nil
		}
		return result.(object.Object)
	case *ast.BlockStatement:
		result := node.Accept(i).(object.Object)
		if result == nil {
			return nil
		}
		return result.(object.Object)
	case *ast.ExpressionStatement:
		result := node.Accept(i).(object.Object)
		if result == nil {
			return nil
		}
		return result.(object.Object)
	}
	return nil
}

// Statements

func (i *Interpreter) VisitProgram(v *ast.Program) (result interface{}) {
	for _, stmt := range v.Statements {
		result = stmt.Accept(i)
		if result != nil {
			return
		}
	}
	return nil
}

func (i *Interpreter) VisitBlock(v *ast.BlockStatement) (result interface{}) {
	for _, stmt := range v.Statements {
		result = i.execute(stmt)
	}
	return
}

func (i *Interpreter) VisitBreak(v *ast.BreakStatement) (result interface{}) {
	return nil
}

func (i *Interpreter) VisitDelete(v *ast.DeleteStatement) (result interface{}) {
	return nil
}

func (i *Interpreter) VisitEnum(v *ast.EnumStatement) (result interface{}) {
	return nil
}

func (i *Interpreter) VisitExprStmt(v *ast.ExpressionStatement) (result interface{}) {
	return i.evaluate(v.Expression)
}

func (i *Interpreter) VisitReturn(v *ast.ReturnStatement) (result interface{}) {
	return nil
}

func (i *Interpreter) VisitVarStmt(v *ast.VarStatement) (result interface{}) {
	i.env.Set(v.Name.Value, i.evaluate(v.Value))
	return nil
}

func (i *Interpreter) VisitAssignStmt(v *ast.AssignStatement) (result interface{}) {
	ident, ok := v.Name.(*ast.Identifier)
	if !ok {
		return nil
	}

	left := ident.Value
	i.env.Set(left, i.evaluate(v.Value))
	return nil
}

// Expresions

func (i *Interpreter) VisitArrayExpr(v *ast.ArrayLiteral) (result interface{}) {
	elements := i.evaluateExpressions(v.Elements)
	if len(elements) == 1 && object.IsError(elements[0]) {
		return elements[0]
	}
	return &object.Array{Elements: elements}
}

func (i *Interpreter) VisitBooleanExpr(v *ast.Boolean) (result interface{}) {
	return &object.Boolean{Value: v.Value}
}

func (i *Interpreter) VisitCallExpr(v *ast.CallExpression) (result interface{}) {
	function := i.evaluate(v.Function)

	if object.IsError(function) {
		return function
	}

	args := i.evaluateExpressions(v.Arguments)

	if len(args) == 1 && object.IsError(args[0]) {
		return args[0]
	}

	return applyFunction(function, args)
}

func (i *Interpreter) VisitFloatExpr(v *ast.FloatLiteral) (result interface{}) {
	return &object.Float{Value: v.Value}
}

func (i *Interpreter) VisitFuncExpr(v *ast.FunctionLiteral) (result interface{}) {
	fn := &object.FunctionLiteral{Parameters: v.Parameters, Body: v.Body}
	if v.Name != nil {
		i.env.Set(v.Name.Value, fn)
	}
	return fn
}

func (i *Interpreter) VisitHashExpr(v *ast.HashLiteral) (result interface{}) {
	pairs := make(map[object.HashKey]object.HashPair)

	for keyNode, valueNode := range v.Pairs {
		key := i.evaluate(keyNode)
		if object.IsError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return object.NewErrorFormat("unusable as hash key: %s", key.Type())
		}

		value := i.evaluate(valueNode)
		if object.IsError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}

func (i *Interpreter) VisitIdentExpr(v *ast.Identifier) (result interface{}) {
	value, ok := i.env.Get(v.Value)
	if !ok {
		return object.NULL
	}
	return value
}

func (i *Interpreter) VisitIfExpr(v *ast.IfExpression) (result interface{}) {
	if object.IsTruthy(i.evaluate(v.Condition)) {
		return i.execute(v.Consequence)
	}
	if v.Alternative != nil {
		return i.execute(v.Alternative)
	}
	return object.NULL
}

func (i *Interpreter) VisitScopeOperatorExpression(v *ast.ScopeOperatorExpression) (result interface{}) {
	return nil
}

func (i *Interpreter) VisitImportExpr(v *ast.Import) (result interface{}) {
	return nil
}

func (i *Interpreter) VisitIndexExpr(v *ast.IndexExpression) (result interface{}) {
	left := i.evaluate(v.Left)
	index := i.evaluate(v.Index)
	return indexExpression(left, index)
}

func (i *Interpreter) VisitIntegerExpr(v *ast.IntegerLiteral) (result interface{}) {
	return &object.Integer{Value: v.Value}
}

func (i *Interpreter) VisitObjectCallExpr(v *ast.ObjectCall) (result interface{}) {
	return nil
}

func (i *Interpreter) VisitPostfixExpr(v *ast.PostfixExpression) (result interface{}) {
	return postfixExpression(v, i.evaluate(v.Left))
}

func (i *Interpreter) VisitPrefixExpr(v *ast.PrefixExpression) (result interface{}) {
	return prefixExpression(v, i.evaluate(v.Right))
}

func (i *Interpreter) VisitStringExpr(v *ast.StringLiteral) (result interface{}) {
	return &object.String{Value: v.Value}
}

func (i *Interpreter) VisitTernaryOperator(v *ast.TernaryOperatorExpression) (result interface{}) {
	condition := i.evaluate(v.Condition)
	if object.IsTruthy(condition) {
		return i.evaluate(v.Consequence)
	}

	return i.evaluate(v.Alternative)
}

func (i *Interpreter) VisitElvisOperator(v *ast.ElvisOperatorExpression) (result interface{}) {
	left := i.evaluate(v.Left)
	if object.IsTruthy(left) {
		return left
	}
	return i.evaluate(v.Right)
}

func (i *Interpreter) VisitFor(v *ast.ForStatement) (result interface{}) {
	return nil
}

func (i *Interpreter) VisitInfix(v *ast.InfixExpression) (result interface{}) {
	return infixExpression(v, v.Operator, i.evaluate(v.Left), i.evaluate(v.Right))
}

func (i *Interpreter) evaluateExpressions(exprs []ast.Expression) []object.Object {
	var result []object.Object

	for _, e := range exprs {
		evaluated := i.evaluate(e)
		if object.IsError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func (i *Interpreter) applyFunction(fn object.Object, args []object.Object) object.Object {

	switch fn := fn.(type) {
	case *object.FunctionLiteral:
		if err := argumentsIsValid(args, fn.Parameters); err != nil {
			return object.NewErrorFormat(err.Error()+" at %s", fn.Body.Token)
		}
		// extendedEnv := extendFunctionEnv(fn.Env, fn.Parameters, args)
		_ = i.execute(fn.Body)
		return nil
		// return unwrapReturnValue(evaluated)
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
