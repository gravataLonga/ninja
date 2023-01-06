package ast

import (
	"github.com/gravataLonga/ninja/token"
)

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (il *FloatLiteral) expressionNode()      {}
func (il *FloatLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *FloatLiteral) String() string {
	return il.Token.Literal
}

func (il *FloatLiteral) Accept(visitor ExprVisitor) (object interface{}) {
	return visitor.VisitFloatExpr(il)
}
