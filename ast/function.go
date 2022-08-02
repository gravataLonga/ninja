package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/token"
	"strings"
)

type Function struct {
	Token      token.Token // The 'function' token
	Name       *Identifier
	Parameters []Expression
	Body       *BlockStatement
}

func (fl *Function) expressionNode()      {}
func (fl *Function) TokenLiteral() string { return fl.Token.Literal }
func (fl *Function) String() string {
	var out bytes.Buffer
	params := make([]string, len(fl.Parameters))
	for i, p := range fl.Parameters {
		params[i] = p.String()
	}
	out.WriteString(fl.TokenLiteral() + " ")
	out.WriteString(fl.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}
