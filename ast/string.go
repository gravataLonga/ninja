package ast

import (
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/token"
	"github.com/gravataLonga/ninja/visitor"
)

type StringLiteral struct {
	Token token.Token
	Value string
}

func (il *StringLiteral) expressionNode()      {}
func (il *StringLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *StringLiteral) String() string {
	return il.TokenLiteral()
}

func (il *StringLiteral) Accept(visitor visitor.ExprVisitor) (object object.Object) {
	return object
}
