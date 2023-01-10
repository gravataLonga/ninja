package ast

import (
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/token"
)

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
	Stack Stack
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) Accept(visitor ExprVisitor) (object object.Object) {
	return visitor.VisitIdentExpr(i)
}
