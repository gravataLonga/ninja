package ast

import (
	"github.com/gravataLonga/ninja/token"
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

func (il *StringLiteral) Accept(visitor ExprVisitor) (object interface{}) {
	return visitor.VisitStringExpr(il)
}
