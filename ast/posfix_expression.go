package ast

import (
	"bytes"
	"ninja/token"
)

type PostfixExpression struct {
	Token    token.Token
	Operator string
}

func (pe *PostfixExpression) expressionNode()      {}
func (pe *PostfixExpression) TokenLiteral() string { return string(pe.Token.Literal) }

func (pe *PostfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.TokenLiteral())
	out.WriteString(pe.Operator)
	out.WriteString(")")
	return out.String()
}
