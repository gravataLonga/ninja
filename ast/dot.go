package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/token"
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

func (oc *ObjectCall) Accept(visitor ExprVisitor) (object interface{}) {
	return visitor.VisitObjectCallExpr(oc)
}
