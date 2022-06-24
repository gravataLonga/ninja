package ast

import "ninja/token"

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return string(il.Token.Literal) }
func (il *IntegerLiteral) String() string {
	return il.TokenLiteral()
}
