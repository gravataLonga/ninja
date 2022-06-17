package evaluator

import "testing"

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"!10.0", false},
		{"!!10.0", true},
		{"![0, 1, 2][0]", false},
		{"![false, 1, 2][0]", true},
		{"!{0: 1}", false},
		{"!{0: false}[0]", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}
