package evaluator

import (
	"ninja/object"
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

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestErrorBooleanHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"true - false;",
			"unknown operator: BOOLEAN - BOOLEAN",
		},
		{
			"true < false;",
			"unknown operator: BOOLEAN < BOOLEAN",
		},
		{
			"true > false;",
			"unknown operator: BOOLEAN > BOOLEAN",
		},
		{
			"true >= false;",
			"unknown operator: BOOLEAN >= BOOLEAN",
		},
		{
			"true <= false;",
			"unknown operator: BOOLEAN <= BOOLEAN",
		},
		{
			"true + false + true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
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
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"[true, false][0] + true",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"{0: true, 1: false}[0] + true",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"[] + true",
			"type mismatch: ARRAY + BOOLEAN",
		},
		{
			"{} + true",
			"type mismatch: HASH + BOOLEAN",
		},
		{
			"function () {} + true",
			"type mismatch: FUNCTION + BOOLEAN",
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

func TestBooleanMethod(t *testing.T) {

	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`true.type()`,
			"BOOLEAN",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		testObjectLiteral(t, evaluated, tt.expected)
	}
}
