package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/token"
)

type TernaryOperatorExpression struct {
	Token       token.Token // The '?' token
	Condition   Expression
	Consequence Expression
	Alternative Expression
}

func (to *TernaryOperatorExpression) expressionNode()      {}
func (to *TernaryOperatorExpression) TokenLiteral() string { return to.Token.Literal }
func (to *TernaryOperatorExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(to.Condition.String())
	out.WriteString("?")
	out.WriteString(to.Consequence.String())
	out.WriteString(":")
	out.WriteString(to.Alternative.String())
	out.WriteString(")")
	return out.String()
}

type ElvisOperatorExpression struct {
	Token token.Token // The '?:' token
	Left  Expression
	Right Expression
}

func (eo *ElvisOperatorExpression) expressionNode()      {}
func (eo *ElvisOperatorExpression) TokenLiteral() string { return eo.Token.Literal }
func (eo *ElvisOperatorExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(eo.Left.String())
	out.WriteString("?:")
	out.WriteString(eo.Right.String())
	out.WriteString(")")
	return out.String()
}
