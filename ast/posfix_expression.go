package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/token"
)

type PostfixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
}

func (pe *PostfixExpression) expressionNode()      {}
func (pe *PostfixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PostfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Left.String())
	out.WriteString(pe.Operator)
	out.WriteString(")")
	return out.String()
}

func (pe *PostfixExpression) Accept(visitor ExprVisitor) (object interface{}) {
	return visitor.VisitPostfixExpr(pe)
}
