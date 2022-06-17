package evaluator

import (
	"ninja/object"
	"testing"
)

func TestEvalHashExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"{};", object.Hash{}},
		{"!{}", false},
		{"!!{}", true},
		{"{1: 1}[1]", 1},
		{"{1: \"ola ola\"}[1]", "ola ola"},
		{"{1: 2.5}[1]", 2.5},
		{"{1: true}[1]", true},
		{"{1: 1 + 1}[1]", 2},
		{"{1: !true}[1]", false},
		{"{1 + 1: 4}[2]", 4},
		{"{} == {}", false},
		{"{} == {1: 0}", false},
		{"{} != {}", false},
		{"{} != {1: 2}", false},
		{"{} && {}", true},
		{"{} || {}", true},
		{"{} && false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testObjectLiteral(t, evaluated, tt.expected)
	}
}

func TestErrorHashHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"-{};",
			"unknown operator: -HASH",
		},
		{
			"{} + {}",
			"unknown operator: HASH + HASH",
		},
		{
			"{} - {}",
			"unknown operator: HASH - HASH",
		},
		{
			"{} > {}",
			"unknown operator: HASH > HASH",
		},
		{
			"{} < {}",
			"unknown operator: HASH < HASH",
		},
		{
			"{} <= {}",
			"unknown operator: HASH <= HASH",
		},
		{
			"{} >= {}",
			"unknown operator: HASH >= HASH",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

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
