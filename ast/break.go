package ast

import (
	"github.com/gravataLonga/ninja/token"
)

type BreakStatement struct {
	Token token.Token // the 'return' token
}

func (rs *BreakStatement) statementNode()       {}
func (rs *BreakStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *BreakStatement) String() string {
	return rs.TokenLiteral()
}
