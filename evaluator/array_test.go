package evaluator

import (
	"fmt"
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

func TestArrayLiteralsAssign(t *testing.T) {
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

func TestArrayMethod(t *testing.T) {

	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`[].type()`,
			"ARRAY",
		},
		{
			`[].length()`,
			0,
		},
		{
			`[1].length()`,
			1,
		},
		{
			`[1,2,3].length()`,
			3,
		},
		{
			`[].join(",")`,
			`[]`,
		},
		{
			`[1, true, 1.1, "hello", function() {return 1;}()].join(",")`,
			`[1,true,1.1,hello,1]`,
		},
		{
			`var a = [1, 2]; a.join(";");`,
			`[1;2]`,
		},
		{
			`var a = [1, 2]; a.join(";"); a`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 1}, &object.Integer{Value: 2}}},
		},
		{
			`var a = [1]; a.push(2); a`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 1}, &object.Integer{Value: 2}}},
		},
		{
			`var a = [1]; a.push(2, 3); a`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 1}, &object.Integer{Value: 2}, &object.Integer{Value: 3}}},
		},
		{
			`var a = [1]; a.push(2, 3); a`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 1}, &object.Integer{Value: 2}, &object.Integer{Value: 3}}},
		},
		{
			`var a = [1, 2]; a.pop(); a;`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 1}}},
		},
		{
			`var a = [1, 2]; a.pop();`,
			2,
		},
		{
			`[1, 2].pop()`,
			2,
		},
		{
			`[].pop()`,
			nil,
		},
		{
			`[1, 2].shift()`,
			1,
		},
		{
			`var a = [1, 2]; a.shift()`,
			1,
		},
		{
			`var a = [1, 2]; a.shift(); a`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 2}}},
		},
		{
			`[].shift()`,
			nil,
		},
		{
			`[1, 2, 3].slice(1)`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 2}, &object.Integer{Value: 3}}},
		},
		{
			`var a = [1, 2, 3]; a.slice(1)`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 2}, &object.Integer{Value: 3}}},
		},
		{
			`[1, 2, 3].slice(4)`,
			object.Array{Elements: []object.Object{}},
		},
		{
			`var a = [1, 2, 3]; a.slice(4)`,
			object.Array{Elements: []object.Object{}},
		},
		{
			`[1, 2, 3].slice(1, 1)`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 2}}},
		},
		{
			`var a = [1, 2, 3]; a.slice(1, 1)`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 2}}},
		},
		{
			`[1, 2, 3].slice(1, 2)`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 2}, &object.Integer{Value: 3}}},
		},
		{
			`var a = [1, 2, 3]; a.slice(1, 2)`,
			object.Array{Elements: []object.Object{&object.Integer{Value: 2}, &object.Integer{Value: 3}}},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestArrayMethod[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			testObjectLiteral(t, evaluated, tt.expected)
		})

	}
}

func TestArrayMethodWrongUsage(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`[].type(1)`,
			"TypeError: array.type() takes exactly 0 argument (1 given)",
		},
		{
			`[].join([])`,
			"TypeError: array.join() expected argument #1 to be `STRING` got `ARRAY`",
		},
		{
			`[1].length(1)`,
			"TypeError: array.length() takes exactly 0 argument (1 given)",
		},
		{
			`[].join()`,
			"TypeError: array.join() takes exactly 1 argument (0 given)",
		},
		{
			`[].push()`,
			"TypeError: array.push() takes a minimum 1 arguments (0 given)",
		},
		{
			`[1].pop(1)`,
			"TypeError: array.pop() takes exactly 0 argument (1 given)",
		},
		{
			`[1].shift(1)`,
			"TypeError: array.shift() takes exactly 0 argument (1 given)",
		},
		{
			`[1].slice(1, 2, 3)`,
			`TypeError: array.push() takes at least 1 arguments at most 2 (3 given)`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestArrayMethodWrongUsage[%d]", i), func(t *testing.T) {
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
