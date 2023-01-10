package ast

import (
	"github.com/gravataLonga/ninja/object"
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

func (rs *BreakStatement) Accept(visitor StmtVisitor) (object object.Object) {
	return visitor.VisitBreak(rs)
}
