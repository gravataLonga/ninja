package ast

import (
	"bytes"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/token"
)

type IndexExpression struct {
	Token token.Token // The [ token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}

func (ie *IndexExpression) Accept(visitor ExprVisitor) (object object.Object) {
	return visitor.VisitIndexExpr(ie)
}
