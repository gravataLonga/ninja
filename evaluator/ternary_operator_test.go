package evaluator

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
			evaluated := testEval(tt.input, t)
			testObjectLiteral(t, evaluated, tt.expected)
		})

	}
}
