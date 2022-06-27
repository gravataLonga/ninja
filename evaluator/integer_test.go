package evaluator

import (
	"ninja/object"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{

		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"++4", 5},
		{"--6", 5},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"[1, 2][0] + 1", 2},
		{`{"key":1}["key"] + 1`, 2},
		{`function () { return 1 }() + 1`, 2},
		{`var add = function() {return 1;}; add() + 1`, 2},
		{`var add = function() {return 1;}; [add(), add()][0] + 1`, 2},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorIntegerHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"[1, 2] + 5",
			"type mismatch: ARRAY + INTEGER",
		},
		{
			"5 + [1, 2]",
			"type mismatch: INTEGER + ARRAY",
		},
		{
			"{} + 5",
			"type mismatch: HASH + INTEGER",
		},
		{
			"10 + {}",
			"type mismatch: INTEGER + HASH",
		},
		{
			"function () {} + 5",
			"type mismatch: FUNCTION + INTEGER",
		},
		{
			"5 + function () {}",
			"type mismatch: INTEGER + FUNCTION",
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

func TestIntegerMethod(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`1.string()`,
			"1",
		},
		{
			`1.type()`,
			"INTEGER",
		},
		{
			`1.float()`,
			1.0,
		},
		{
			`-1.abs()`,
			1,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		testObjectLiteral(t, evaluated, tt.expected)
	}
}

func TestIntegerMethodWrongUsage(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`1.type(1)`,
			"method type not accept any arguments. got: [1]",
		},
		{
			`1.string(1)`,
			"method string not accept any arguments. got: [1]",
		},
		{
			`1.float(1)`,
			"method float not accept any arguments. got: [1]",
		},
		{
			`1.abs(1)`,
			"method abs not accept any arguments. got: [1]",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
		}

		if errObj.Message != tt.expectedErrorMessage {
			t.Errorf("erro expected \"%s\". Got: %s", tt.expectedErrorMessage, errObj.Message)
		}
	}
}
