package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/object"
)

type Node interface {
	// TokenLiteral is used only for testing and debugging
	// return literal associated with
	TokenLiteral() string

	// String is handy for testing...
	String() string
}

type Statement interface {
	Node
	statementNode()
	Accept(visitor StmtVisitor) (object object.Object)
}

type Expression interface {
	Node
	expressionNode()
	Accept(visitor ExprVisitor) (object object.Object)
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (p *Program) Accept(visitor StmtVisitor) (object object.Object) {
	return visitor.VisitProgram(p)
}
