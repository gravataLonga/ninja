package interpreter

import (
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
	return node.Accept(i)
}

func (i *Interpreter) Interpreter(node ast.Node) object.Object {
	return i.execute(node)
}

func (i *Interpreter) execute(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return node.Accept(i)
	}
	return nil
}

// Statements

func (i *Interpreter) VisitProgram(v *ast.Program) (result object.Object) {
	for _, stmt := range v.Statements {
		err := stmt.Accept(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) VisitBlock(v *ast.BlockStatement) (object object.Object) {
	return nil
}

func (i *Interpreter) VisitBreak(v *ast.BreakStatement) (object object.Object) {
	return nil
}

func (i *Interpreter) VisitDelete(v *ast.DeleteStatement) (object object.Object) {
	return nil
}

func (i *Interpreter) VisitEnum(v *ast.EnumStatement) (object object.Object) {
	return nil
}

func (i *Interpreter) VisitExprStmt(v *ast.ExpressionStatement) (object object.Object) {
	return i.evaluate(v.Expression)
}

func (i *Interpreter) VisitReturn(v *ast.ReturnStatement) (object object.Object) {
	return nil
}

func (i *Interpreter) VisitVarStmt(v *ast.VarStatement) (object object.Object) {
	return nil
}

func (i *Interpreter) VisitAssignStmt(v *ast.AssignStatement) (object object.Object) {
	return nil
}

// Expresions

func (i *Interpreter) VisitArrayExpr(v *ast.ArrayLiteral) (result object.Object) {
	elements := i.evaluateExpressions(v.Elements)
	if len(elements) == 1 && object.IsError(elements[0]) {
		return elements[0]
	}
	return &object.Array{Elements: elements}
}

func (i *Interpreter) VisitBooleanExpr(v *ast.Boolean) (result object.Object) {
	return &object.Boolean{Value: v.Value}
}

func (i *Interpreter) VisitCallExpr(v *ast.CallExpression) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitFloatExpr(v *ast.FloatLiteral) (result object.Object) {
	return &object.Float{Value: v.Value}
}

func (i *Interpreter) VisitFuncExpr(v *ast.FunctionLiteral) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitHashExpr(v *ast.HashLiteral) (result object.Object) {
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

func (i *Interpreter) VisitIdentExpr(v *ast.Identifier) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitIfExpr(v *ast.IfExpression) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitScopeOperatorExpression(v *ast.ScopeOperatorExpression) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitImportExpr(v *ast.Import) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitIndexExpr(v *ast.IndexExpression) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitIntegerExpr(v *ast.IntegerLiteral) (result object.Object) {
	return &object.Integer{Value: v.Value}
}

func (i *Interpreter) VisitObjectCallExpr(v *ast.ObjectCall) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitPostfixExpr(v *ast.PostfixExpression) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitPrefixExpr(v *ast.PrefixExpression) (result object.Object) {
	return prefixExpression(v, i.evaluate(v.Right))
}

func (i *Interpreter) VisitStringExpr(v *ast.StringLiteral) (result object.Object) {
	return &object.String{Value: v.Value}
}

func (i *Interpreter) VisitTernaryOperator(v *ast.TernaryOperatorExpression) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitElvisOperator(v *ast.ElvisOperatorExpression) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitFor(v *ast.ForStatement) (result object.Object) {
	return nil
}

func (i *Interpreter) VisitInfix(v *ast.InfixExpression) (result object.Object) {
	left := i.evaluate(v.Left)
	right := i.evaluate(v.Right)

	switch v.Operator {
	case "+":
		vLeft, ok := left.(*object.Integer)
		if !ok {
		}

		vRight, ok := right.(*object.Integer)
		if !ok {
		}

		result := &object.Integer{Value: vLeft.Value + vRight.Value}
		return result
	}
	return nil
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
