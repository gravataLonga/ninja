package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/token"
	"github.com/gravataLonga/ninja/visitor"
)

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) Accept(visitor visitor.StmtVisitor) (object object.Object) {
	return visitor.VisitReturn(rs)
}
