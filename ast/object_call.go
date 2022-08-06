package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/token"
)

type ObjectCall struct {
	Token  token.Token
	Object Expression
	Call   Expression
}

func (oc *ObjectCall) expressionNode()      {}
func (oc *ObjectCall) TokenLiteral() string { return oc.Token.Literal }
func (oc *ObjectCall) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oc.Object.String())
	out.WriteString(".")
	out.WriteString(oc.Call.String())
	out.WriteString(")")

	return out.String()
}

func (oc *ObjectCall) Accept(visitor ExprVisitor) (object object.Object) {
	return visitor.VisitObjectCallExpr(oc)
}
