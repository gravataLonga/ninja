package ast

import (
	"github.com/gravataLonga/ninja/token"
	"testing"
)

func TestArrayLiteral_String(t *testing.T) {
	elements := []Expression{
		&IntegerLiteral{Token: token.Token{Type: token.LBRACKET, Literal: "1"}, Value: 1},
		&FloatLiteral{Token: token.Token{Type: token.INT, Literal: "2.2"}, Value: 2.2},
		&StringLiteral{Token: token.Token{Type: token.STRING, Literal: "3"}, Value: "3"},
		&Boolean{Token: token.Token{Type: token.TRUE, Literal: "True"}, Value: true},
	}
	arrLiteral := &ArrayLiteral{Token: token.Token{Type: token.LBRACKET, Literal: "["}, Elements: elements}

	if arrLiteral.String() != "[1, 2.2, 3, True]" {
		t.Errorf("ArrayLiteral isnt equal to %s. Got: %s", "[1, 2, 3, True]", arrLiteral.String())
	}

	if arrLiteral.TokenLiteral() != "[" {
		t.Errorf("ArrayLiteral.TokenLiteral() isnt equal to %s. Got: %s", "[", arrLiteral.TokenLiteral())
	}
}

func BenchmarkArrayLiteral_String(b *testing.B) {
	elements := []Expression{
		&IntegerLiteral{Token: token.Token{Type: token.LBRACKET, Literal: "1"}, Value: 1},
		&FloatLiteral{Token: token.Token{Type: token.INT, Literal: "2.2"}, Value: 2.2},
		&StringLiteral{Token: token.Token{Type: token.STRING, Literal: "3"}, Value: "3"},
		&Boolean{Token: token.Token{Type: token.TRUE, Literal: "True"}, Value: true},
	}
	arrLiteral := &ArrayLiteral{Token: token.Token{Type: token.LBRACKET, Literal: "["}, Elements: elements}

	for i := 0; i < b.N; i++ {
		arrLiteral.String()
	}
}

func BenchmarkArrayLiteral_StringWithHash(b *testing.B) {
	elements := []Expression{
		&IntegerLiteral{Token: token.Token{Type: token.LBRACKET, Literal: "1"}, Value: 1},
		&FloatLiteral{Token: token.Token{Type: token.INT, Literal: "2.2"}, Value: 2.2},
		&StringLiteral{Token: token.Token{Type: token.STRING, Literal: "3"}, Value: "3"},
		&Boolean{Token: token.Token{Type: token.TRUE, Literal: "True"}, Value: true},
		&HashLiteral{Pairs: map[Expression]Expression{
			&IntegerLiteral{Token: token.Token{Type: token.LBRACKET, Literal: "1"}, Value: 1}: &Boolean{Token: token.Token{Type: token.TRUE, Literal: "True"}, Value: true},
		}},
	}
	arrLiteral := &ArrayLiteral{Token: token.Token{Type: token.LBRACKET, Literal: "["}, Elements: elements}

	for i := 0; i < b.N; i++ {
		arrLiteral.String()
	}
}
