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
			`var i = 0; for(;;) { i = i + 1; break; }; i;`,
			1,
		},
		{
			`var i = 0; for(;;) { break; }; i;`,
			0,
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
