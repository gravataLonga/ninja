package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/token"
)

type ForStatement struct {
	Token            token.Token // The 'for' token
	InitialCondition *VarStatement
	Condition        Expression
	Iteration        Statement
	Body             *BlockStatement
}

func (fs *ForStatement) expressionNode()      {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out bytes.Buffer

	out.WriteString(string(fs.TokenLiteral()) + " ")

	out.WriteString("(")
	if fs.InitialCondition != nil {
		out.WriteString(fs.InitialCondition.String())
	} else {
		out.WriteString(";")
	}

	if fs.Condition != nil {
		out.WriteString(fs.Condition.String() + ";")
	} else {
		out.WriteString(";")
	}

	if fs.Iteration != nil {
		out.WriteString(fs.Iteration.String())
	}

	out.WriteString(") ")
	out.WriteString("{")
	out.WriteString(fs.Body.String())
	out.WriteString("}")

	return out.String()
}
