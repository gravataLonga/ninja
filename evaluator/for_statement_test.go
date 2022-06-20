package evaluator

import (
	"ninja/object"
	"testing"
)

func TestForStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`for(var i = 0; i <= 1; i = i + 1) { i; }`,
			1,
		},
		{
			`for(;;) { return; }`,
			object.NULL,
		},
		{
			`for(var i = 0; i > 10; i = i + 1) { i; }`,
			nil,
		},
		{
			`var i = 0; for(; i <= 1; i = i + 1) { i; }`,
			1,
		},
		{
			`for(;;) { return 1; }`,
			1,
		},
		{
			`var total = 0; var arr = [1, 1]; for(var i = 0; i <= len(arr) -1; i = i + 1) { total = total + arr[i]; }; total;`,
			2,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else if evaluated != nil {
			t.Errorf("result isnt nil. Got %v", evaluated)
		}
	}
}
