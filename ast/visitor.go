package ast

import (
	"github.com/gravataLonga/ninja/object"
)

type ExprVisitor interface {
	VisitArrayExpr(v *ArrayLiteral) (result object.Object)
	VisitBooleanExpr(v *Boolean) (result object.Object)
	VisitCallExpr(v *CallExpression) (result object.Object)
	VisitDotExpr(v *Dot) (result object.Object)
	VisitFloatExpr(v *FloatLiteral) (result object.Object)
	VisitFuncExpr(v *FunctionLiteral) (result object.Object)
	VisitHashExpr(v *HashLiteral) (result object.Object)
	VisitIdentExpr(v *Identifier) (result object.Object)
	VisitIfExpr(v *IfExpression) (result object.Object)
	VisitScopeOperatorExpression(v *ScopeOperatorExpression) (result object.Object)
	VisitImportExpr(v *Import) (result object.Object)
	VisitIndexExpr(v *IndexExpression) (result object.Object)
	VisitIntegerExpr(v *IntegerLiteral) (result object.Object)
	VisitPostfixExpr(v *PostfixExpression) (result object.Object)
	VisitPrefixExpr(v *PrefixExpression) (result object.Object)
	VisitStringExpr(v *StringLiteral) (result object.Object)
	VisitTernaryOperator(v *TernaryOperatorExpression) (result object.Object)
	VisitElvisOperator(v *ElvisOperatorExpression) (result object.Object)
	VisitFor(v *ForStatement) (result object.Object)
	VisitInfix(v *InfixExpression) (result object.Object)
}

type StmtVisitor interface {
	VisitProgram(v *Program) (result object.Object)
	VisitBlock(v *BlockStatement) (result object.Object)
	VisitBreak(v *BreakStatement) (result object.Object)
	VisitDelete(v *DeleteStatement) (result object.Object)
	VisitEnum(v *EnumStatement) (result object.Object)
	VisitExprStmt(v *ExpressionStatement) (result object.Object)
	VisitReturn(v *ReturnStatement) (result object.Object)
	VisitVarStmt(v *VarStatement) (result object.Object)
	VisitAssignStmt(v *AssignStatement) (result object.Object)
}
