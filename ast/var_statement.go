package ast

import (
	"bytes"
	"ninja/token"
)

type VarStatement struct {
	Token token.Token // the token.VAR token
	Name  *Identifier
	Value Expression
}

func (ls *VarStatement) statementNode()       {}
func (ls *VarStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *VarStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

type AssignStatement struct {
	Token token.Token // the token.VAR token
	Name  Expression  // it can be var a = a + 1; or a = a + 1; or a[0] = 1;
	Value Expression  // Any valid expression
}

func (ls *AssignStatement) expressionNode()      {}
func (ls *AssignStatement) statementNode()       {}
func (ls *AssignStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *AssignStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}
