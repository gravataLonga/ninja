package ast

import "github.com/gravataLonga/ninja/token"

type Parameter struct {
	Token   token.Token // the identifier token
	Name    *Identifier
	Default Expression // nil if no default
}

func (p *Parameter) expressionNode()      {}
func (p *Parameter) TokenLiteral() string { return p.Token.Literal }
func (p *Parameter) String() string {
	if p.Default != nil {
		return "(" + p.Name.String() + " = " + p.Default.String() + ")"
	}
	return p.Name.String()
}
