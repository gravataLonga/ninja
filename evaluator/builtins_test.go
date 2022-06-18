package evaluator

import (
	"ninja/object"
	"testing"
)

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
		nullable bool
	}{
		{`len("")`, 0, false},
		{`len("four")`, 4, false},
		{`len("hello world")`, 11, false},
		{`len(1)`, "argument to `len` not supported, got INTEGER", false},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1", false},
		{`len([1, 2, 3])`, 3, false},
		{`len([])`, 0, false},
		{`puts("hello", "world!")`, nil, true},
		{`first([1, 2, 3])`, 1, false},
		{`first([])`, nil, false},
		{`first(1)`, "argument to `first` must be ARRAY, got INTEGER", false},
		{`last([1, 2, 3])`, 3, false},
		{`last([])`, nil, false},
		{`last(1)`, "argument to `last` must be ARRAY, got INTEGER", false},
		{`rest([1, 2, 3])`, []int{2, 3}, false},
		{`rest([])`, nil, false},
		{`push([], 1)`, []int{1}, false},
		{`push(1, 1)`, "argument to `push` must be ARRAY, got INTEGER", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case nil:
			if tt.nullable {
				if evaluated != nil {
					t.Errorf("Test must return nil. Got: %T", evaluated)
				}

				continue
			}
			testNullObject(t, evaluated)

		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)",
					evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q",
					expected, errObj.Message)
			}
		case []int:
			array, ok := evaluated.(*object.Array)
			if !ok {
				t.Errorf("obj not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(array.Elements) != len(expected) {
				t.Errorf("wrong num of elements. want=%d, got=%d",
					len(expected), len(array.Elements))
				continue
			}

			for i, expectedElem := range expected {
				testIntegerObject(t, array.Elements[i], int64(expectedElem))
			}

		}

	}
}
