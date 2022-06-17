package evaluator

import (
	"testing"
)

func TestEvalDigitExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`1 + 1`,
			2,
		},
		{
			`1.0 + 1.0`,
			2.0,
		},
		{
			`1 + 1.0`,
			2.0,
		},
		{
			`1.0 + 1`,
			2.0,
		},
		{
			`1.0 + 1`,
			2.0,
		},
		{
			`[1.0, 3.0][1] + 1`,
			4.0,
		},
		{
			`[1, 3.0][1] + 1.2`,
			4.2,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testObjectLiteral(t, evaluated, tt.expected)
	}
}
