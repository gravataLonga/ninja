package ast

import (
	"bytes"
	"ninja/token"
	"strings"
)

type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return string(al.Token.Literal) }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := make([]string, len(al.Elements))
	for i, el := range al.Elements {
		elements[i] = el.String()
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
