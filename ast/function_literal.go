package ast

import (
	"bytes"
	"ninja/token"
	"strings"
)

type FunctionLiteral struct {
	Token      token.Token // The 'function' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return string(fl.Token.Literal) }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := make([]string, len(fl.Parameters))
	for i, p := range fl.Parameters {
		params[i] = p.String()
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}
