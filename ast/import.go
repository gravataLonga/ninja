package ast

import (
	"bytes"
	"fmt"
	"ninja/token"
)

type Import struct {
	Token    token.Token
	Filename Expression
}

func (i *Import) expressionNode()      {}
func (i *Import) TokenLiteral() string { return i.Token.Literal }
func (i *Import) String() string {
	var out bytes.Buffer

	out.WriteString(i.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(fmt.Sprintf("\"%s\"", i.Filename))

	return out.String()
}
