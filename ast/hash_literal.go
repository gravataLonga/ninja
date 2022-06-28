package ast

import (
	"bytes"
	"ninja/token"
	"strings"
)

type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return string(hl.Token.Literal) }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer
	pairs := make([]string, len(hl.Pairs))
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
