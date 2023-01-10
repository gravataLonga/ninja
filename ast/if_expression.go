package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/token"
	"github.com/gravataLonga/ninja/visitor"
)

type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

func (ie *IfExpression) Accept(visitor visitor.ExprVisitor) (object object.Object) {
	return visitor.VisitIfExpr(ie)
}
