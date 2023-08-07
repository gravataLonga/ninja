package interpreter

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/stdlib"
	"io"
)

type Interpreter struct {
	env       *object.Environment
	innerLoop int
	output    io.Writer
}

func New(w io.Writer, env *object.Environment) *Interpreter {
	env.Set("len", object.NewBuiltin(stdlib.Len))
	env.Set("first", object.NewBuiltin(stdlib.First))
	env.Set("puts", object.NewBuiltin(stdlib.Puts))
	env.Set("last", object.NewBuiltin(stdlib.Last))
	env.Set("rest", object.NewBuiltin(stdlib.Rest))
	env.Set("push", object.NewBuiltin(stdlib.Push))
	env.Set("time", object.NewBuiltin(stdlib.Time))
	env.Set("rand", object.NewBuiltin(stdlib.Rand))
	env.Set("args", object.NewBuiltin(stdlib.Args))
	env.Set("plugin", object.NewBuiltin(stdlib.Plugin))

	return &Interpreter{
		env:    env,
		output: w,
	}
}

func (i *Interpreter) EnterLoop() {
	i.innerLoop++
}

func (i *Interpreter) ExitLoop() {
	i.innerLoop--
}

func (i *Interpreter) InLoop() bool {
	return i.innerLoop > 0
}

func (i *Interpreter) Interpreter(node ast.Node) object.Object {
	return i.execute(node)
}

func (i *Interpreter) evaluate(node ast.Expression) object.Object {
	return node.Accept(i)
}

func (i *Interpreter) execute(node ast.Node) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		result := node.Accept(i)
		return result
	case *ast.BlockStatement:
		result := node.Accept(i)
		return result
	case *ast.ExpressionStatement:
		result := node.Accept(i)
		return result
	case *ast.ReturnStatement:
		result := node.Accept(i)
		return result
	case *ast.BreakStatement:
		result := node.Accept(i)
		return result
	case *ast.VarStatement:
		result := node.Accept(i)
		return result
	case *ast.AssignStatement:
		result := node.Accept(i)
		return result
	}
	return nil
}

// Statements

func (i *Interpreter) VisitProgram(v *ast.Program) (result object.Object) {
	for _, stmt := range v.Statements {
		result = stmt.Accept(i)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Break:
			return object.NewErrorFormat("'break' not in the 'loop' context")
		case *object.Error:
			return result
		}
	}
	return
}

func (i *Interpreter) VisitBlock(v *ast.BlockStatement) (result object.Object) {
	for _, stmt := range v.Statements {
		result = i.execute(stmt)

		if result == nil {
			continue
		}

		if object.IsError(result) {
			return
		}

		// @todo test it better
		if i.InLoop() {
			if object.IsReturn(result) {
				return
			}

			if result.Type() == object.BREAK_VALUE_OBJ {
				return
			}
		}

		if object.IsReturn(result) {
			return
		}

	}
	return
}

func (i *Interpreter) VisitBreak(v *ast.BreakStatement) (result object.Object) {
	return &object.Break{Value: object.NULL}
}

func (i *Interpreter) VisitExprStmt(v *ast.ExpressionStatement) (result object.Object) {
	return i.evaluate(v.Expression)
}

// Expresions

func (i *Interpreter) VisitIndexExpr(v *ast.IndexExpression) (result object.Object) {
	left := i.evaluate(v.Left)
	index := i.evaluate(v.Index)
	return indexExpression(left, index)
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
