package ast

import (
	"testing"
)

func TestFloatLiteral_String(t *testing.T) {
	tests := []struct {
		float    float64
		expected string
	}{
		{
			10.3,
			"10.3",
		},
		{
			2.0,
			"2",
		},
		{
			1000000.20,
			"1000000.2",
		},
	}

	for _, tt := range tests {
		literal := createFloatLiteral(tt.float)

		if literal.String() != tt.expected {
			t.Errorf("FloatLiteral.String() not match to %s. Got: %s", tt.expected, literal.String())
		}
	}
}
