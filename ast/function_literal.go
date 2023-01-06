package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/token"
	"strings"
)

type FunctionLiteral struct {
	Token      token.Token // The 'function' token
	Parameters []Expression
	Body       *BlockStatement
	Name       *Identifier
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := make([]string, len(fl.Parameters))
	for i, p := range fl.Parameters {
		params[i] = p.String()
	}
	out.WriteString(fl.TokenLiteral())
	if fl.Name != nil {
		out.WriteString(" ")
		out.WriteString(fl.Name.String())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {")
	out.WriteString(fl.Body.String())
	out.WriteString("}")
	return out.String()
}

func (fl *FunctionLiteral) Accept(visitor ExprVisitor) (object interface{}) {
	return visitor.VisitFuncExpr(fl)
}
