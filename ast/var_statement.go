package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/token"
)

type VarStatement struct {
	Token token.Token // the token.VAR token
	Name  *Identifier
	Value Expression
}

func (ls *VarStatement) statementNode()       {}
func (ls *VarStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *VarStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

func (ls *VarStatement) Accept(visitor StmtVisitor) (object object.Object) {
	return visitor.VisitVarStmt(ls)
}

type AssignStatement struct {
	Token token.Token // the token.VAR token
	Left  Node
	Right Expression // Any valid expression
}

func (ls *AssignStatement) statementNode()       {}
func (ls *AssignStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *AssignStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.Left.String())
	out.WriteString(" = ")

	if ls.Right != nil {
		out.WriteString(ls.Right.String())
	}
	out.WriteString(";")

	return out.String()
}

func (ls *AssignStatement) Accept(visitor StmtVisitor) (object object.Object) {
	return visitor.VisitAssignStmt(ls)
}
