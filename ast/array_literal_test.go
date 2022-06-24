package ast

import (
	"ninja/token"
	"testing"
)

func TestArrayLiteral_String(t *testing.T) {
	elements := []Expression{
		&IntegerLiteral{Token: token.Token{Type: token.LBRACKET, Literal: []byte("1")}, Value: 1},
		&FloatLiteral{Token: token.Token{Type: token.INT, Literal: []byte("2.2")}, Value: 2.2},
		&StringLiteral{Token: token.Token{Type: token.STRING, Literal: []byte("3")}, Value: "3"},
		&Boolean{Token: token.Token{Type: token.TRUE, Literal: []byte("True")}, Value: true},
	}
	arrLiteral := &ArrayLiteral{Token: token.Token{Type: token.LBRACKET, Literal: []byte("[")}, Elements: elements}

	if arrLiteral.String() != "[1, 2.2, 3, True]" {
		t.Errorf("ArrayLiteral isnt equal to %s. Got: %s", "[1, 2, 3, True]", arrLiteral.String())
	}

	if arrLiteral.TokenLiteral() != "[" {
		t.Errorf("ArrayLiteral.TokenLiteral() isnt equal to %s. Got: %s", "[", arrLiteral.TokenLiteral())
	}

}
