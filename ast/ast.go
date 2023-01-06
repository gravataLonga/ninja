package ast

import (
	"bytes"
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
	Accept(visitor StmtVisitor) (object interface{})
}

type Expression interface {
	Node
	expressionNode()
	Accept(visitor ExprVisitor) (object interface{})
}

type ExprVisitor interface {
	VisitArrayExpr(v *ArrayLiteral) (result interface{})
	VisitBooleanExpr(v *Boolean) (result interface{})
	VisitCallExpr(v *CallExpression) (result interface{})
	VisitFloatExpr(v *FloatLiteral) (result interface{})
	VisitFuncExpr(v *FunctionLiteral) (result interface{})
	VisitHashExpr(v *HashLiteral) (result interface{})
	VisitIdentExpr(v *Identifier) (result interface{})
	VisitIfExpr(v *IfExpression) (result interface{})
	VisitScopeOperatorExpression(v *ScopeOperatorExpression) (result interface{})
	VisitImportExpr(v *Import) (result interface{})
	VisitIndexExpr(v *IndexExpression) (result interface{})
	VisitIntegerExpr(v *IntegerLiteral) (result interface{})
	VisitObjectCallExpr(v *ObjectCall) (result interface{})
	VisitPostfixExpr(v *PostfixExpression) (result interface{})
	VisitPrefixExpr(v *PrefixExpression) (result interface{})
	VisitStringExpr(v *StringLiteral) (result interface{})
	VisitTernaryOperator(v *TernaryOperatorExpression) (result interface{})
	VisitElvisOperator(v *ElvisOperatorExpression) (result interface{})
	VisitFor(v *ForStatement) (result interface{})
	VisitInfix(v *InfixExpression) (result interface{})
}

type StmtVisitor interface {
	VisitProgram(v *Program) (result interface{})
	VisitBlock(v *BlockStatement) (result interface{})
	VisitBreak(v *BreakStatement) (result interface{})
	VisitDelete(v *DeleteStatement) (result interface{})
	VisitEnum(v *EnumStatement) (result interface{})
	VisitExprStmt(v *ExpressionStatement) (result interface{})
	VisitReturn(v *ReturnStatement) (result interface{})
	VisitVarStmt(v *VarStatement) (result interface{})
	VisitAssignStmt(v *AssignStatement) (result interface{})
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

func (p *Program) Accept(visitor StmtVisitor) (object interface{}) {
	return visitor.VisitProgram(p)
}
