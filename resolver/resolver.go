package resolver

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/interpreter"
	"github.com/gravataLonga/ninja/object"
)

type Resolver struct {
	i interpreter.Interpreter
}

func New(i interpreter.Interpreter) *Resolver {
	return &Resolver{i: i}
}

func (r *Resolver) Resolve(node ast.Node) {
	r.resolve(node)
}

func (r *Resolver) resolve(node ast.Node) {
	switch node := node.(type) {

	case *ast.Program:
		node.Accept(r)
		break
	case *ast.BlockStatement:
		node.Accept(r)
		break
	case *ast.ExpressionStatement:
		node.Accept(r)
		break
	case *ast.ReturnStatement:
		node.Accept(r)
		break
	case *ast.BreakStatement:
		node.Accept(r)
		break
	case *ast.VarStatement:
		node.Accept(r)
		break
	case *ast.AssignStatement:
		node.Accept(r)
		break
	}
}

func (r *Resolver) VisitArrayExpr(v *ast.ArrayLiteral) (result object.Object) {
	for _, item := range v.Elements {
		item.Accept(r)
	}
	return nil
}

func (r *Resolver) VisitBooleanExpr(v *ast.Boolean) (result object.Object) {
	return nil
}

func (r *Resolver) VisitCallExpr(v *ast.CallExpression) (result object.Object) {
	for _, arg := range v.Arguments {
		arg.Accept(r)
	}
	v.Function.Accept(r)
	return nil
}

func (r *Resolver) VisitDotExpr(v *ast.Dot) (result object.Object) {
	v.Object.Accept(r)
	v.Right.Accept(r)
	return nil
}

func (r *Resolver) VisitFloatExpr(v *ast.FloatLiteral) (result object.Object) {
	return nil
}

func (r *Resolver) VisitFuncExpr(v *ast.FunctionLiteral) (result object.Object) {
	for _, params := range v.Parameters {
		params.Accept(r)
	}
	v.Body.Accept(r)
	return nil
}

func (r *Resolver) VisitHashExpr(v *ast.HashLiteral) (result object.Object) {
	for _, exp := range v.Pairs {
		exp.Accept(r)
	}
	return nil
}

func (r *Resolver) VisitIdentExpr(v *ast.Identifier) (result object.Object) {
	return nil
}

func (r *Resolver) VisitIfExpr(v *ast.IfExpression) (result object.Object) {
	v.Condition.Accept(r)
	v.Consequence.Accept(r)
	if v.Alternative != nil {
		v.Alternative.Accept(r)
	}
	return nil
}

func (r *Resolver) VisitScopeOperatorExpression(v *ast.ScopeOperatorExpression) (result object.Object) {
	v.PropertyIdentifier.Accept(r)
	v.AccessIdentifier.Accept(r)
	return nil
}

func (r *Resolver) VisitImportExpr(v *ast.Import) (result object.Object) {
	v.Filename.Accept(r)
	return nil
}

func (r *Resolver) VisitIndexExpr(v *ast.IndexExpression) (result object.Object) {
	v.Left.Accept(r)
	v.Index.Accept(r)
	return nil
}

func (r *Resolver) VisitIntegerExpr(v *ast.IntegerLiteral) (result object.Object) {
	return nil
}

func (r *Resolver) VisitPostfixExpr(v *ast.PostfixExpression) (result object.Object) {
	v.Left.Accept(r)
	return nil
}

func (r *Resolver) VisitPrefixExpr(v *ast.PrefixExpression) (result object.Object) {
	v.Right.Accept(r)
	return nil
}

func (r *Resolver) VisitStringExpr(v *ast.StringLiteral) (result object.Object) {
	return nil
}

func (r *Resolver) VisitTernaryOperator(v *ast.TernaryOperatorExpression) (result object.Object) {
	v.Condition.Accept(r)
	v.Consequence.Accept(r)
	v.Alternative.Accept(r)
	return nil
}

func (r *Resolver) VisitElvisOperator(v *ast.ElvisOperatorExpression) (result object.Object) {
	v.Left.Accept(r)
	v.Right.Accept(r)
	return nil
}

func (r *Resolver) VisitFor(v *ast.ForStatement) (result object.Object) {
	if v.InitialCondition != nil {
		v.InitialCondition.Accept(r)
	}

	if v.Condition != nil {
		v.Condition.Accept(r)
	}

	if v.Iteration != nil {
		v.Iteration.Accept(r)
	}

	if v.Body != nil {
		v.Body.Accept(r)
	}
	return nil
}

func (r *Resolver) VisitInfix(v *ast.InfixExpression) (result object.Object) {
	v.Left.Accept(r)
	v.Right.Accept(r)
	return nil
}

func (r *Resolver) VisitProgram(v *ast.Program) (result object.Object) {
	for _, stmt := range v.Statements {
		result = stmt.Accept(r)
	}
	return
}

func (r *Resolver) VisitBlock(v *ast.BlockStatement) (result object.Object) {
	for _, stmt := range v.Statements {
		stmt.Accept(r)
	}
	return
}

func (r *Resolver) VisitBreak(v *ast.BreakStatement) (result object.Object) {
	return nil
}

func (r *Resolver) VisitDelete(v *ast.DeleteStatement) (result object.Object) {
	v.Left.Accept(r)
	v.Index.Accept(r)
	return nil
}

func (r *Resolver) VisitEnum(v *ast.EnumStatement) (result object.Object) {
	v.Identifier.Accept(r)
	for _, a := range v.Branches {
		a.Accept(r)
	}
	return nil
}

func (r *Resolver) VisitExprStmt(v *ast.ExpressionStatement) (result object.Object) {
	return v.Accept(r)
}

func (r *Resolver) VisitReturn(v *ast.ReturnStatement) (result object.Object) {
	v.ReturnValue.Accept(r)
	return nil
}

func (r *Resolver) VisitVarStmt(v *ast.VarStatement) (result object.Object) {
	return v.Value.Accept(r)
}

func (r *Resolver) VisitAssignStmt(v *ast.AssignStatement) (result object.Object) {
	v.Right.Accept(r)
	return nil
}
