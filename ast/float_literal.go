package ast

import (
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/token"
	"github.com/gravataLonga/ninja/visitor"
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

func (il *FloatLiteral) Accept(visitor visitor.ExprVisitor) (object object.Object) {
	return visitor.VisitFloatExpr(il)
}
