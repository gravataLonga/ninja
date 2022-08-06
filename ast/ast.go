package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/object"
)

type Node interface {
	// TokenLiteral is used only for testing and debugging
	// return literal associated with
	TokenLiteral() string

	// String is handy for testing...
	String() string
}

type Statement interface {
	Node
	statementNode()
	Accept(visitor StmtVisitor) (object object.Object)
}

type Expression interface {
	Node
	expressionNode()
	Accept(visitor ExprVisitor) (object object.Object)
}

type ExprVisitor interface {
	VisitArrayExpr(v *ArrayLiteral) (result object.Object)
	VisitBooleanExpr(v *Boolean) (result object.Object)
	VisitCallExpr(v *CallExpression) (result object.Object)
	VisitFloatExpr(v *FloatLiteral) (result object.Object)
	VisitFuncExpr(v *FunctionLiteral) (result object.Object)
	VisitHashExpr(v *HashLiteral) (result object.Object)
	VisitIdentExpr(v *Identifier) (result object.Object)
	VisitIfExpr(v *IfExpression) (result object.Object)
	VisitScopeOperatorExpression(v *ScopeOperatorExpression) (result object.Object)
	VisitImportExpr(v *Import) (result object.Object)
	VisitIndexExpr(v *IndexExpression) (result object.Object)
	VisitIntegerExpr(v *IntegerLiteral) (result object.Object)
	VisitObjectCallExpr(v *ObjectCall) (result object.Object)
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
	VisitBlock(v *BlockStatement) (object object.Object)
	VisitBreak(v *BreakStatement) (object object.Object)
	VisitDelete(v *DeleteStatement) (object object.Object)
	VisitEnum(v *EnumStatement) (object object.Object)
	VisitExprStmt(v *ExpressionStatement) (object object.Object)
	VisitReturn(v *ReturnStatement) (object object.Object)
	VisitVarStmt(v *VarStatement) (object object.Object)
	VisitAssignStmt(v *AssignStatement) (object object.Object)
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (p *Program) Accept(visitor StmtVisitor) (object object.Object) {
	return visitor.VisitProgram(p)
}
