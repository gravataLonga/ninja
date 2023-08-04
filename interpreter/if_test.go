package interpreter

import (
	"fmt"
	"testing"
)

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (0) { 10 }", 10},
		// everthing below this line, is desnecessary they are been checked by boolean_test.go
		// but we keep them for sanity check.
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 <= 2) { 10 }", 10},
		{"if (1 >= 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 >= 2) { 10 } else { 20 }", 20},
		{"if (1 <= 2) { 10 } else { 20 }", 10},
		{"if (1.0) { 10 }", 10},
		{"if (1.0 < 2.0) { 10 }", 10},
		{"if (1.0 > 2.0) { 10 }", nil},
		{"if (1.0 <= 2.0) { 10 }", 10},
		{"if (1.0 >= 2.0) { 10 }", nil},
		{"if (1.0 > 2.0) { 10 } else { 20 }", 20},
		{"if (1.0 < 2.0) { 10 } else { 20 }", 10},
		{"if (1.0 >= 2.0) { 10 } else { 20 }", 20},
		{"if (1.0 <= 2.0) { 10 } else { 20 }", 10},
		{"if (1) { 20.50 }", 20.50},
		{"if (1.0 < 2.0) { 50.20 }", 50.20},
		{"if (2.0 != 2.0) { 50.20 }", nil},
		{"if (2.0 == 2.0) { 1 }", 1},
		{"if (true && true) { 1 }", 1},
		// {"if (\"a\" == \"a\") { 1 }", 1},
	}

	for o, tt := range tests {
		t.Run(fmt.Sprintf("TestIfElseExpressions[%d]", o), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)
			testObjectLiteral(t, evaluated, tt.expected)
		})

	}
}
