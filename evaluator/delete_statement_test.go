package evaluator

import (
	"ninja/object"
	"testing"
)

func TestDeleteStatementArray(t *testing.T) {
	tests := []struct {
		input        string
		expectedInts []int64
	}{
		{
			`var a = [0, 1]; delete a[0]; a;`,
			[]int64{1},
		},
		{
			`var a = [0, 1, 2, 3]; delete a[3]; a;`,
			[]int64{0, 1, 2},
		},
		{
			`var a = [0, 1, 2, 3]; delete a[1+1]; a;`,
			[]int64{0, 1, 3},
		},
		{
			`var b = 0; var a = [0, 1, 2, 3]; delete a[b]; a;`,
			[]int64{1, 2, 3},
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		result, ok := evaluated.(*object.Array)
		if !ok {
			t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
		}

		for index, i := range tt.expectedInts {
			testIntegerObject(t, result.Elements[index], i)
		}

	}
}

func TestWrongDeleteStatementArray(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			`var a = ""; delete a[0];`,
			"DeleteStatement.left only work with array or hash object. Got: *object.String",
		},
		{
			`var a = []; delete a[{}];`,
			"DeleteStatement.index must be a Integer. Got: *object.Hash",
		},
		{
			`var a = ""; delete b[0];`,
			"DeleteStatement.left b identifier not found.",
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

func TestDeleteStatementHash(t *testing.T) {
	tests := []struct {
		input string
	}{
		{
			`var a = {"key":1, "key2":2}; delete a["key"]; a;`,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		result, ok := evaluated.(*object.Hash)
		if !ok {
			t.Fatalf("object is not Hash. got=%T (%+v)", evaluated, evaluated)
		}

		if len(result.Pairs) != 1 {
			t.Fatalf("object unable to delete key. Got: %+v", result.Pairs)
		}

	}
}
