package ast

import (
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/token"
)

type ContinueStatement struct {
	Token token.Token // the 'continue' token
}

func (rs *ContinueStatement) statementNode()       {}
func (rs *ContinueStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ContinueStatement) String() string {
	return rs.TokenLiteral()
}

func (rs *ContinueStatement) Accept(visitor StmtVisitor) (object object.Object) {
	return visitor.VisitContinue(rs)
}
