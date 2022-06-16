package ast

import (
	"ninja/token"
	"testing"
)

func TestBoolean_String(t *testing.T) {
	test := Boolean{Token: token.Token{Type: token.TRUE, Literal: "true"}, Value: true}

	if test.String() != "true" {
		t.Errorf("ArrayLiteral isnt equal to %s. Got: %s", "true", test.String())
	}

	if test.TokenLiteral() != "true" {
		t.Errorf("ArrayLiteral.TokenLiteral() isnt equal to %s. Got: %s", "true", test.TokenLiteral())
	}

}
