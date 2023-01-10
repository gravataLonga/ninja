package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/token"
	"github.com/gravataLonga/ninja/visitor"
)

type Dot struct {
	Token  token.Token
	Object Expression
	Right  Expression
}

func (oc *Dot) expressionNode()      {}
func (oc *Dot) TokenLiteral() string { return oc.Token.Literal }
func (oc *Dot) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oc.Object.String())
	out.WriteString(".")
	out.WriteString(oc.Right.String())
	out.WriteString(")")

	return out.String()
}

func (oc *Dot) Accept(visitor visitor.ExprVisitor) (object object.Object) {
	return visitor.VisitDotExpr(oc)
}
