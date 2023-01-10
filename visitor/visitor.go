package visitor

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
)

type Expression interface {
	ast.Node
	expressionNode()
	Accept(visitor ExprVisitor) (object object.Object)
}

type ExprVisitor interface {
	VisitArrayExpr(v *ast.ArrayLiteral) (result object.Object)
	VisitBooleanExpr(v *ast.Boolean) (result object.Object)
	VisitCallExpr(v *ast.CallExpression) (result object.Object)
	VisitDotExpr(v *ast.Dot) (result object.Object)
	VisitFloatExpr(v *ast.FloatLiteral) (result object.Object)
	VisitFuncExpr(v *ast.FunctionLiteral) (result object.Object)
	VisitHashExpr(v *ast.HashLiteral) (result object.Object)
	VisitIdentExpr(v *ast.Identifier) (result object.Object)
	VisitIfExpr(v *ast.IfExpression) (result object.Object)
	VisitScopeOperatorExpression(v *ast.ScopeOperatorExpression) (result object.Object)
	VisitImportExpr(v *ast.Import) (result object.Object)
	VisitIndexExpr(v *ast.IndexExpression) (result object.Object)
	VisitIntegerExpr(v *ast.IntegerLiteral) (result object.Object)
	VisitPostfixExpr(v *ast.PostfixExpression) (result object.Object)
	VisitPrefixExpr(v *ast.PrefixExpression) (result object.Object)
	VisitStringExpr(v *ast.StringLiteral) (result object.Object)
	VisitTernaryOperator(v *ast.TernaryOperatorExpression) (result object.Object)
	VisitElvisOperator(v *ast.ElvisOperatorExpression) (result object.Object)
	VisitFor(v *ast.ForStatement) (result object.Object)
	VisitInfix(v *ast.InfixExpression) (result object.Object)
}

type StmtVisitor interface {
	VisitProgram(v *ast.Program) (result object.Object)
	VisitBlock(v *ast.BlockStatement) (result object.Object)
	VisitBreak(v *ast.BreakStatement) (result object.Object)
	VisitDelete(v *ast.DeleteStatement) (result object.Object)
	VisitEnum(v *ast.EnumStatement) (result object.Object)
	VisitExprStmt(v *ast.ExpressionStatement) (result object.Object)
	VisitReturn(v *ast.ReturnStatement) (result object.Object)
	VisitVarStmt(v *ast.VarStatement) (result object.Object)
	VisitAssignStmt(v *ast.AssignStatement) (result object.Object)
}
