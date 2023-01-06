package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/token"
)

type DeleteStatement struct {
	Token token.Token // the delete token
	Left  Expression
	Index Expression
}

func (de *DeleteStatement) statementNode()       {}
func (de *DeleteStatement) TokenLiteral() string { return de.Token.Literal }
func (de *DeleteStatement) String() string {
	var out bytes.Buffer

	out.WriteString(de.TokenLiteral() + " ")
	out.WriteString(de.Left.String())
	out.WriteString("[")
	out.WriteString(de.Index.String())
	out.WriteString("]")
	return out.String()
}

func (de *DeleteStatement) Accept(visitor StmtVisitor) (object interface{}) {
	return visitor.VisitDelete(de)
}
