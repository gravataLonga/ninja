package evaluator

import (
	"ninja/object"
	"testing"
)

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5.2", 5.2},
		{"10.0", 10.0},
		{"10.000000000123", 10.0000000001},
		{"-5.2", -5.2},
		{"-10.0", -10.0},
		{"++1.0", 2.0},
		{"--1.0", 0.0},
		{"5.0 + 5.0 + 5.5 + 5.5 - 10", 11},
		{"2.2 * 2.2 * 2.2 * 2.2 * 2.2", 51.53632},
		{"-50.50 + 100.50 + -50.50", -0.5},
		{"5.5 * 2.5 + 10.5", 24.25},
		{"5.5 + 2.5 * 10.5", 31.75},
		{"20 + 2.0 * -10", 0.0},
		{"50.10 / 2.20 * 2.20 + 10.2", 60.2999999999},
		{"2 * (5.2 + 10.2)", 30.8},
		{"3 * 3 * 3 + 10.5", 37.5},
		{"3 * (3 * 3.5) + 10", 41.5},
		{"(5 + 10 * 2 + 15 / 3) * 2.2 + -10", 56.0},
		{"[1.2, 4.2][0] + 1.3", 2.5},
		{`{"key":1.2}["key"] + 1.3`, 2.5},
		{`function () { return 1.2 }() + 1.3`, 2.5},
		{`var add = function() {return 1.2;}; add() + 1.3`, 2.5},
		{`var add = function() {return 1.2;}; [add(), add()][0] + 1.3`, 2.5},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)
		testFloatObject(t, evaluated, tt.expected)
	}
}

func TestErrorFloatHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"50.50 + true;",
			"type mismatch: FLOAT + BOOLEAN",
		},
		{
			"5.0 + true; 5.3;",
			"type mismatch: FLOAT + BOOLEAN",
		},
		{
			"[] + 10.3",
			"type mismatch: ARRAY + FLOAT",
		},
		{
			"10.3 + []",
			"type mismatch: FLOAT + ARRAY",
		},
		{
			"10.3 + {}",
			"type mismatch: FLOAT + HASH",
		},
		{
			"{} + 10.3",
			"type mismatch: HASH + FLOAT",
		},
		{
			"function () {} + 10.3",
			"type mismatch: FUNCTION + FLOAT",
		},
		{
			"10.3 + function () {}",
			"type mismatch: FLOAT + FUNCTION",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}
