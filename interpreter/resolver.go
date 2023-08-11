package interpreter

import (
	"github.com/gravataLonga/ninja/analysis"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

type Resolver struct {
	i     *Interpreter
	stack *analysis.Stack
}

func NewResolver(i *Interpreter) *Resolver {
	return &Resolver{i: i, stack: analysis.NewStack()}
}

func (resolver *Resolver) BeginScope() {
	resolver.stack.Push(analysis.NewScope())
}

func (resolver *Resolver) EndScope() {
	resolver.stack.Pop()
}

func (resolver *Resolver) Resolve(node ast.Node) {
	resolver.resolve(node)
}

func (resolver *Resolver) resolve(node ast.Node) {
	switch node := node.(type) {

	case *ast.Program:
		node.Accept(resolver)
		break
	case *ast.BlockStatement:
		node.Accept(resolver)
		break
	case *ast.ExpressionStatement:
		node.Accept(resolver)
		break
	case *ast.ReturnStatement:
		node.Accept(resolver)
		break
	case *ast.BreakStatement:
		node.Accept(resolver)
		break
	case *ast.VarStatement:
		node.Accept(resolver)
		break
	case *ast.AssignStatement:
		node.Accept(resolver)
		break
	}
}

func (resolver *Resolver) VisitArrayExpr(v *ast.ArrayLiteral) (result object.Object) {
	for _, item := range v.Elements {
		item.Accept(resolver)
	}
	return nil
}

func (resolver *Resolver) VisitBooleanExpr(v *ast.Boolean) (result object.Object) {
	return nil
}

func (resolver *Resolver) VisitCallExpr(v *ast.CallExpression) (result object.Object) {
	for _, arg := range v.Arguments {
		arg.Accept(resolver)
	}
	v.Function.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitDotExpr(v *ast.Dot) (result object.Object) {
	v.Object.Accept(resolver)
	v.Right.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitFloatExpr(v *ast.FloatLiteral) (result object.Object) {
	return nil
}

func (resolver *Resolver) VisitFuncExpr(v *ast.FunctionLiteral) (result object.Object) {
	for _, params := range v.Parameters {
		params.Accept(resolver)
	}
	v.Body.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitHashExpr(v *ast.HashLiteral) (result object.Object) {
	for _, exp := range v.Pairs {
		exp.Accept(resolver)
	}
	return nil
}

func (resolver *Resolver) VisitIdentExpr(v *ast.Identifier) (result object.Object) {
	return nil
}

func (resolver *Resolver) VisitIfExpr(v *ast.IfExpression) (result object.Object) {
	if v.Condition != nil {
		v.Condition.Accept(resolver)
	}

	if v.Consequence != nil {
		v.Consequence.Accept(resolver)
	}

	if v.Alternative != nil {
		v.Alternative.Accept(resolver)
	}
	return nil
}

func (resolver *Resolver) VisitScopeOperatorExpression(v *ast.ScopeOperatorExpression) (result object.Object) {
	v.PropertyIdentifier.Accept(resolver)
	v.AccessIdentifier.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitImportExpr(v *ast.Import) (result object.Object) {
	v.Filename.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitIndexExpr(v *ast.IndexExpression) (result object.Object) {
	v.Left.Accept(resolver)
	v.Index.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitIntegerExpr(v *ast.IntegerLiteral) (result object.Object) {
	return nil
}

func (resolver *Resolver) VisitPostfixExpr(v *ast.PostfixExpression) (result object.Object) {
	v.Left.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitPrefixExpr(v *ast.PrefixExpression) (result object.Object) {
	v.Right.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitStringExpr(v *ast.StringLiteral) (result object.Object) {
	return nil
}

func (resolver *Resolver) VisitTernaryOperator(v *ast.TernaryOperatorExpression) (result object.Object) {
	v.Condition.Accept(resolver)
	v.Consequence.Accept(resolver)
	v.Alternative.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitElvisOperator(v *ast.ElvisOperatorExpression) (result object.Object) {
	v.Left.Accept(resolver)
	v.Right.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitFor(v *ast.ForStatement) (result object.Object) {
	if v.InitialCondition != nil {
		v.InitialCondition.Accept(resolver)
	}

	if v.Condition != nil {
		v.Condition.Accept(resolver)
	}

	if v.Iteration != nil {
		v.Iteration.Accept(resolver)
	}

	if v.Body != nil {
		v.Body.Accept(resolver)
	}
	return nil
}

func (resolver *Resolver) VisitInfix(v *ast.InfixExpression) (result object.Object) {
	v.Left.Accept(resolver)
	v.Right.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitProgram(v *ast.Program) (result object.Object) {
	for _, stmt := range v.Statements {
		result = stmt.Accept(resolver)
	}
	return
}

func (resolver *Resolver) VisitBlock(v *ast.BlockStatement) (result object.Object) {
	resolver.BeginScope()
	for _, stmt := range v.Statements {
		stmt.Accept(resolver)
	}
	resolver.EndScope()
	return
}

func (resolver *Resolver) VisitBreak(v *ast.BreakStatement) (result object.Object) {
	return nil
}

func (resolver *Resolver) VisitDelete(v *ast.DeleteStatement) (result object.Object) {
	v.Left.Accept(resolver)
	v.Index.Accept(resolver)
	return nil
}

func (resolver *Resolver) VisitEnum(v *ast.EnumStatement) (result object.Object) {
	v.Identifier.Accept(resolver)
	for _, a := range v.Branches {
		a.Accept(resolver)
	}
	return nil
}

func (resolver *Resolver) VisitExprStmt(v *ast.ExpressionStatement) (result object.Object) {
	return v.Expression.Accept(resolver)
}

func (resolver *Resolver) VisitReturn(v *ast.ReturnStatement) (result object.Object) {
	if v.ReturnValue != nil {
		return v.ReturnValue.Accept(resolver)
	}
	return nil
}

func (resolver *Resolver) VisitVarStmt(v *ast.VarStatement) (result object.Object) {
	if resolver.stack.Size() == 0 {
		return v.Value.Accept(resolver)
	}
	resolver.stack.Peek().Put(v.Name.Value, false)
	result = v.Value.Accept(resolver)
	resolver.stack.Peek().Put(v.Name.Value, true)
	// inform *Interpreter about this AST where is located.

	return
}

func (resolver *Resolver) VisitAssignStmt(v *ast.AssignStatement) (result object.Object) {
	// Check r.s.Peek().Exists() == false then give and error.
	// Check r.s.Peek().Get() == false wasn't initilized yet, give and error.

	v.Right.Accept(resolver)
	return nil
}
