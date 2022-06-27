package ast

import (
	"bytes"
	"ninja/token"
)

type ObjectCall struct {
	Token  token.Token
	Object Expression
	Call   Expression
}

func (oc *ObjectCall) expressionNode()      {}
func (oc *ObjectCall) TokenLiteral() string { return string(oc.Token.Literal) }
func (oc *ObjectCall) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oc.Object.String())
	out.WriteString(".")
	out.WriteString(oc.Call.String())
	out.WriteString(")")

	return out.String()
}
