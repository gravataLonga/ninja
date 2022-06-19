package ast

import (
	"math"
	"ninja/token"
)

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (il *FloatLiteral) expressionNode()      {}
func (il *FloatLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *FloatLiteral) String() string {
	return il.Token.Literal
}

// FloatSmall return always smallest decimal place
func FloatSmall(f float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(f*ratio) / ratio
}
