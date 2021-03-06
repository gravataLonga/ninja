package evaluator

import (
	"fmt"
	"github.com/gravataLonga/ninja/object"
	"testing"
)

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5.2", 5.2},
		{"10.0", 10.0},
		{"10.000000000123", 10.000000000123},
		{"-5.2", -5.2},
		{"-10.0", -10.0},
		// Prefix
		{"++1.0", 2.0},
		{"var a = 1.0; ++a; a;", 2.0},
		{"var a = 1.0; ++a;", 2.0},
		{"--1.0", 0.0},
		{"var a = 1.0; --a; a;", 0.0},
		{"var a = 1.0; --a;", 0.0},

		// Postfix
		{"1.0++", 2.0},
		{"var a = 1.0; a++; a;", 2.0},
		{"var a = 1.0; a++;", 1.0},
		{"1.0--", 0.0},
		{"var a = 1.0; a--; a;", 0.0},
		{"var a = 1.0; a--;", 1.0},

		{"--1.0", 0.0},
		{"3.0 ** 0", 1.0},
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

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestEvalFloatExpression[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)
			testFloatObject(t, evaluated, tt.expected)
		})
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

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestErrorFloatHandling[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedMessage {
				t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
			}
		})

	}
}

func TestFloatMethod(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`1.1.string()`,
			"1.1",
		},
		{
			`1.1.type()`,
			"FLOAT",
		},
		{
			`var a = -1.1; a.abs()`,
			1.1,
		},
		{
			`0.8.round()`,
			1.0,
		},
		{
			`0.4.round()`,
			0.0,
		},
		{
			`0.05.round()`,
			0.0,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestFloatMethod[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			testObjectLiteral(t, evaluated, tt.expected)
		})
	}
}

func TestFloatMethodWrongUsage(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`1.1.type(1)`,
			"TypeError: float.type() takes exactly 0 argument (1 given)",
		},
		{
			`1.1.string(1)`,
			"TypeError: float.string() takes exactly 0 argument (1 given)",
		},
		{
			`1.1.abs(1)`,
			"TypeError: float.abs() takes exactly 0 argument (1 given)",
		},
		{
			`1.1.round(1)`,
			"TypeError: float.round() takes exactly 0 argument (1 given)",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestFloatMethodWrongUsage[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedErrorMessage {
				t.Errorf("erro expected \"%s\". Got: %s", tt.expectedErrorMessage, errObj.Message)
			}
		})

	}
}
