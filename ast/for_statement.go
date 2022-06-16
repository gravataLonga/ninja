package ast

import (
	"bytes"
	"ninja/token"
)

type ForStatement struct {
	Token            token.Token // The 'for' token
	InitialCondition *VarStatement
	Condition        Expression
	Iteration        *VarStatement
	Body             *BlockStatement
}

func (fs *ForStatement) expressionNode()      {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out bytes.Buffer

	out.WriteString(fs.TokenLiteral() + " ")

	out.WriteString("(")
	out.WriteString(fs.InitialCondition.String() + ";")
	out.WriteString(fs.Condition.String() + ";")
	out.WriteString(fs.Iteration.String())
	out.WriteString(") ")
	out.WriteString(fs.Body.String())

	return out.String()
}
