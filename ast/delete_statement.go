package ast

import (
	"bytes"
	"ninja/token"
)

type DeleteStatement struct {
	Token token.Token // the delete token
	Left  Expression
	Index Expression
}

func (de *DeleteStatement) statementNode()       {}
func (de *DeleteStatement) TokenLiteral() string { return string(de.Token.Literal) }
func (de *DeleteStatement) String() string {
	var out bytes.Buffer

	out.WriteString(de.TokenLiteral() + " ")
	out.WriteString(de.Left.String())
	out.WriteString("[")
	out.WriteString(de.Index.String())
	out.WriteString("]")
	return out.String()
}
