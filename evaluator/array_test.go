package evaluator

import (
	"ninja/object"
	"testing"
)

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3, 34.4]"

	evaluated := testEval(input, t)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 4 {
		t.Fatalf("array has wrong num of elements. got=%d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
	testFloatObject(t, result.Elements[3], 34.4)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"var i = 0; [1][i];",
			1,
		},
		{
			"[1, 2, 3][1 + 1];",
			3,
		},
		{
			"var myArray = [1, 2, 3]; myArray[2];",
			3,
		},
		{
			"var myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];",
			6,
		},
		{
			"var myArray = [1, 2, 3]; var i = myArray[0]; myArray[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestArrayLiteralsAssing(t *testing.T) {
	tests := []struct {
		input            string
		expected         int64
		expectedElements int
	}{
		{
			`var a = []; a[0] = 1; a;`,
			1,
			1,
		},
		{
			`var a = [0]; a[0] = 2; a;`,
			2,
			1,
		},
		{
			`var a = [0]; a[1] = 2; a;`,
			0,
			2,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)
		result, ok := evaluated.(*object.Array)
		if !ok {
			t.Fatalf("Eval didn't return Array. got=%T (%+v)", evaluated, evaluated)
		}

		if len(result.Elements) != tt.expectedElements {
			t.Fatalf("Array has wrong num of elements, expected %d. got=%d", tt.expectedElements, len(result.Elements))
		}

		testIntegerObject(t, result.Elements[0], tt.expected)
	}

}

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
		evaluated := testEval(tt.input, t)
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
		{
			`var a = []; a[1] = 2;`,
			`index out of range, got 1 but array has only 0 elements`,
		},
		{
			`var a = []; a[-1] = 2;`,
			`index out of range, got -1 not positive index`,
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
