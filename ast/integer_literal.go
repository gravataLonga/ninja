package ast

import (
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/token"
)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string {
	return il.TokenLiteral()
}

func (il *IntegerLiteral) Accept(visitor ExprVisitor) (object object.Object) {
	return visitor.VisitIntegerExpr(il)
}
