package ast

import "ninja/token"

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return string(i.Token.Literal) }
func (i *Identifier) String() string {
	return i.Value
}
