package evaluator

import (
	"ninja/object"
	"testing"
)

func TestHashLiterals(t *testing.T) {
	input := `var two = "two";
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 6 / 2,
		4: 4,
		true: 5,
		false: 6
	}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		object.TRUE.HashKey():                      5,
		object.FALSE.HashKey():                     6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`var key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{true: 5}[true]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
		{
			`{1.0: 35}[1.0]`,
			35,
		},
		{
			`{1.000000: 35}[1.0]`,
			35,
		},
		{
			`{1.000001: 35}[1.0]`,
			nil,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

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
