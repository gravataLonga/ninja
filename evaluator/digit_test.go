package evaluator

import (
	"fmt"
	"testing"
)

func TestEvalDigitExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`1`,
			1,
		},
		{
			`0.2`,
			0.2,
		},
		{
			`1e3`,
			1e3,
		},
		{
			`1e-3`,
			1e-3,
		},
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
		}, {
			`1 - 1`,
			0,
		},
		{
			`1.0 - 1.0`,
			0.0,
		},
		{
			`1 - 1.0`,
			0.0,
		},
		{
			`1.0 - 1`,
			0.0,
		},
		{
			`1.0 - 1`,
			0.0,
		},
		{
			`[1.0, 3.0][1] - 1`,
			2.0,
		},
		{
			`[1, 3.0][1] - 1.2`,
			1.8,
		},
		{
			`2 * 2`,
			4,
		},
		{
			`4 / 2`,
			2,
		},
		{
			`100 / 8`,
			12.5,
		},
		{
			`4 % 2`,
			0,
		},
		{
			`4.0 % 2.0`,
			0.0,
		},
		{
			`4 % 2.0`,
			0.0,
		},
		{
			`4.0 % 2`,
			0.0,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestEvalDigitExpression[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)
			testObjectLiteral(t, evaluated, tt.expected)
		})
	}
}
