package evaluator

import "testing"

func TestForStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`for(var i = 0; i <= 1; i = i + 1) { i; }`,
			2,
		},
		{
			`for(var i = 0; i > 100; i = i + 1) { i; }`,
			nil,
		},
		{
			`var i = 0; for(; i <= 1; i = i + 1) { i; }`,
			2,
		},
		{
			`for(var i = 0; i <= 1; ) { var i = i + 1; i; }`,
			2,
		},
		{
			`var i = 0; var a = 0; for(;i <= 10 && a <= 10;) { var i = i + 2; var a = a + 2; i; }`,
			12,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else if evaluated != nil {
			t.Errorf("result isnt nil. Got %v", evaluated)
		}
	}
}
