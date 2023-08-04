package interpreter

import (
	"fmt"
	"testing"
)

func TestTernaryOperatorExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"true ? 10 : 0", 10},
		{"false ? 10 : 0", 0},
		{"1 ? 10 : 0", 10},
		{"1 < 2 ? 10 : 0", 10},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestTernaryOperatorExpressions[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)
			testObjectLiteral(t, evaluated, tt.expected)
		})

	}
}

func TestElvisOperatorExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"true?:false", true},
		{`"hello"?:"world"`, "hello"},
		{"false?:20", 20},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestElvisOperatorExpressions[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)
			testObjectLiteral(t, evaluated, tt.expected)
		})

	}
}
