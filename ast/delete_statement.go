package ast

import (
	"bytes"
	"ninja/token"
)

type DeleteStatement struct {
	Token      token.Token // the delete token
	Identifier Expression
}

func (de *DeleteStatement) statementNode()       {}
func (de *DeleteStatement) TokenLiteral() string { return de.Token.Literal }
func (de *DeleteStatement) String() string {
	var out bytes.Buffer

	out.WriteString(de.TokenLiteral())
	out.WriteString(" (")
	out.WriteString(de.Identifier.String())
	out.WriteString(") ")
	return out.String()
}
