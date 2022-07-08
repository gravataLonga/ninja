package ast

import "ninja/token"

type Boolean struct {
	Token token.Token
	Value bool
}

func (il *Boolean) expressionNode()      {}
func (il *Boolean) TokenLiteral() string { return il.Token.Literal }
func (il *Boolean) String() string       { return il.Token.Literal }
