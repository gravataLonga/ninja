package ast

import (
	"ninja/token"
)

type BreakStatement struct {
	Token token.Token // the 'return' token
}

func (rs *BreakStatement) statementNode()       {}
func (rs *BreakStatement) TokenLiteral() string { return string(rs.Token.Literal) }
func (rs *BreakStatement) String() string {
	return rs.TokenLiteral()
}
