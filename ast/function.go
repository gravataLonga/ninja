package ast

import (
	"bytes"
	"ninja/token"
	"strings"
)

type Function struct {
	Token      token.Token // The 'function' token
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *Function) expressionNode()      {}
func (fl *Function) TokenLiteral() string { return fl.Token.Literal }
func (fl *Function) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral() + " ")
	out.WriteString(fl.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}
