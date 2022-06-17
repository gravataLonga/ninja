package evaluator

import (
	"ninja/object"
	"testing"
)

func TestEvalArrayExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"[1];", object.Array{Elements: []object.Object{&object.Integer{Value: 1}}}},
		{"[1+1];", object.Array{Elements: []object.Object{&object.Integer{Value: 2}}}},
		{"[\"ola\"];", object.Array{Elements: []object.Object{&object.String{Value: "ola"}}}},
		{"[\"ola\" + \" mundo\"];", object.Array{Elements: []object.Object{&object.String{Value: "ola mundo"}}}},
		{"[\"hello\", 1, 2.2, true, function() {return \"fn\";}][0]", "hello"},
		{"[\"hello\", 1, 2.2, true, function() {return \"fn\";}][1]", 1},
		{"[\"hello\", 1, 2.2, true, function() {return \"fn\";}][2]", 2.2},
		{"[\"hello\", 1, 2.2, true, function() {return \"fn\";}][3]", true},
		{"[\"hello\", 1, 2.2, true, function() {return \"fn\";}][4]()", "fn"},
		{"[] == []", false},
		{"[] != []", false},
		{"[] && []", true},
		{"[] || []", true},
		{"[] && false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testObjectLiteral(t, evaluated, tt.expected)
	}
}

func TestErrorArrayHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"-[];",
			"unknown operator: -ARRAY",
		},
		{
			"[] + [];",
			"unknown operator: ARRAY + ARRAY",
		},
		{
			"[] < [];",
			"unknown operator: ARRAY < ARRAY",
		},
		{
			"[] > [];",
			"unknown operator: ARRAY > ARRAY",
		},
		{
			"[] <= [];",
			"unknown operator: ARRAY <= ARRAY",
		},
		{
			"[] >= [];",
			"unknown operator: ARRAY >= ARRAY",
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
