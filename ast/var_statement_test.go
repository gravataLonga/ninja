package ast

import (
	"testing"
)

func TestVarStatement_String(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected string
	}{
		{
			"age", 31, "var age = 31;",
		},
		{
			"pi", 31415, "var pi = 31415;",
		},
		{
			"year", 2022, "var year = 2022;",
		},
	}

	for _, tt := range tests {

		fn := createVarStatement(tt.name, createIntegerLiteral(tt.value))

		if fn.String() != tt.expected {
			t.Fatalf("Function.String() isn't expected %s. Got %s", tt.expected, fn.String())
		}
	}
}
