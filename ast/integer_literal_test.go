package ast

import (
	"ninja/token"
	"testing"
)

func TestIntegerLiteral_String(t *testing.T) {
	tests := []struct {
		intValue int64
		expected string
	}{
		{
			10,
			"10",
		},
		{
			2,
			"2",
		},
		{
			1000000,
			"1000000",
		},
	}

	for _, tt := range tests {
		literal := &IntegerLiteral{Token: token.Token{Type: token.FLOAT, Literal: []byte(tt.expected)}, Value: tt.intValue}

		if literal.String() != tt.expected {
			t.Errorf("IntegerLiteral.String() not match to %s. Got: %s", tt.expected, literal.String())
		}
	}
}
