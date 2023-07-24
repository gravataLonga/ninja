package interpreter

import (
	"fmt"
	"github.com/gravataLonga/ninja/object"
	"testing"
)

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"true && true", true},
		{"true && false", false},
		{"true || false", true},
		{"false || true", true},
		{"false || false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"1 && 1", true},
		{"1 && 0", true},
		{"0 && 1", true},
		{"1 || false", true},
		{"false || false", false},
		{"false && 1", false},
		{"10.0 == 10.0", true},
		{"10.5 >= 10.0", true},
		{"10.5 && 1", true},
		{"(20.5 > 5) == true", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"(1 <= 1) == true", true},
		{"(1 >= 1) == true", true},
		{"(1 >= 1) && true", true},
		{"(1 || false) || false", true},
		{"function () { return true; }() && true", true},
		{"function () { return true; }() || false", true},
		{"function () { return true; }() && false", false},
		{"[1, 0][0] && true", true},
		{"[1, 0][0] || false", true},
		{"[1, 0][0] && false", false},
		{"{\"0\": 1}[\"0\"] && true", true},
		{"{\"0\": 1}[\"0\"] || false", true},
		{"{\"0\": 1}[\"0\"] && false", false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestEvalBooleanExpression[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)

			testBooleanObject(t, evaluated, tt.expected)
		})
	}
}

func TestErrorBooleanHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"-true",
			"unknown operator: -BOOLEAN - at [Line: 1, Offset: 1]",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN + at [Line: 1, Offset: 6]",
		},
		{
			"true - false;",
			"unknown operator: BOOLEAN - BOOLEAN - at [Line: 1, Offset: 6]",
		},
		{
			"true < false;",
			"unknown operator: BOOLEAN < BOOLEAN < at [Line: 1, Offset: 6]",
		},
		{
			"true > false;",
			"unknown operator: BOOLEAN > BOOLEAN > at [Line: 1, Offset: 6]",
		},
		{
			"true >= false;",
			"unknown operator: BOOLEAN >= BOOLEAN >= at [Line: 1, Offset: 7]",
		},
		{
			"true <= false;",
			"unknown operator: BOOLEAN <= BOOLEAN <= at [Line: 1, Offset: 7]",
		},
		{
			"true + false + true + false;",
			"unknown operator: BOOLEAN + BOOLEAN + at [Line: 1, Offset: 6]",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN + at [Line: 1, Offset: 9]",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN + at [Line: 1, Offset: 20]",
		},
		{
			`
if (10 > 1) {
  if (10 > 1) {
    return true + false;
  }
  return 1;
}
`,
			"unknown operator: BOOLEAN + BOOLEAN + at [Line: 4, Offset: 17]",
		},
		{
			"[true, false][0] + true",
			"unknown operator: BOOLEAN + BOOLEAN + at [Line: 1, Offset: 18]",
		},
		{
			"{0: true, 1: false}[0] + true",
			"unknown operator: BOOLEAN + BOOLEAN + at [Line: 1, Offset: 24]",
		},
		{
			"[] + true",
			"unknown operator: ARRAY + BOOLEAN + at [Line: 1, Offset: 4]",
		},
		{
			"{} + true",
			"unknown operator: HASH + BOOLEAN + at [Line: 1, Offset: 4]",
		},
		{
			"function () {} + true",
			"unknown operator: FUNCTION + BOOLEAN + at [Line: 1, Offset: 16]",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestErrorBooleanHandling[%d]", i), func(t *testing.T) {

			evaluated := interpreter(t, tt.input)

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

func TestBooleanMethod(t *testing.T) {

	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`true.type()`,
			"BOOLEAN",
		},
		{
			`var a = false; a.type()`,
			"BOOLEAN",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestBooleanMethod[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)

			testObjectLiteral(t, evaluated, tt.expected)
		})
	}
}

func TestBooleanWrongMethod(t *testing.T) {

	tests := []struct {
		input                string
		expectedErrorMessage interface{}
	}{
		{
			`true.type(1)`,
			"TypeError: bool.type() takes exactly 0 argument (1 given)",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestBooleanMethod[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)

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
