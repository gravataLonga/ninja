package ast

import (
	"ninja/token"
)

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (il *FloatLiteral) expressionNode()      {}
func (il *FloatLiteral) TokenLiteral() string { return string(il.Token.Literal) }
func (il *FloatLiteral) String() string {
	return string(il.Token.Literal)
}

// ToBeDeleted return always smallest decimal place
func ToBeDeleted(f float64, precision uint) float64 {
	return f
}
