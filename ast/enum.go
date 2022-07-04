package ast

import (
	"ninja/token"
	"strings"
)

type EnumStatement struct {
	Token      token.Token
	Branches   map[string]Expression
	Identifier Expression
}

func (e *EnumStatement) TokenLiteral() string { return e.Token.Literal }
func (e *EnumStatement) statementNode()       {}
func (e *EnumStatement) String() string {
	out := strings.Builder{}
	out.WriteString("enum")
	out.WriteString(e.Identifier.String())
	branches := make([]string, len(e.Branches))
	i := 0
	for name, v := range e.Branches {
		br := strings.Builder{}
		br.WriteString(name)
		br.WriteString(":")
		br.WriteString(v.String())
		branches[i] = br.String()
		i++
	}
	out.WriteString("{")
	out.WriteString(strings.Join(branches, ";"))
	out.WriteString("}")
	return out.String()
}

type ScopeOperatorExpression struct {
	Token              token.Token // "::"
	AccessIdentifier   Expression
	PropertyIdentifier Expression
}

func (so *ScopeOperatorExpression) TokenLiteral() string { return so.Token.Literal }
func (so *ScopeOperatorExpression) expressionNode()      {}
func (so *ScopeOperatorExpression) String() string {
	out := strings.Builder{}
	out.WriteString(so.AccessIdentifier.String())
	out.WriteString(so.TokenLiteral())
	out.WriteString(so.PropertyIdentifier.String())

	return out.String()
}
