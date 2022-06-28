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
		{
			`var i = 0; for(;;) { if( i > 3) { break; } i = i + 1; }; i;`,
			4,
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

func TestBreakOutsideForLoop(t *testing.T) {
	input := `break`
	expected := "'break' not in the 'loop' context"

	evaluated := testEval(input, t)

	if evaluated == nil {
		t.Fatalf("evaluated is empty")
	}

	err, ok := evaluated.(*object.Error)

	if !ok {
		t.Fatalf("expected error. Got: %s", evaluated.Inspect())
	}

	if err.Message != expected {
		t.Fatalf("expected error message to be %s, got: %s", expected, err.Message)
	}

}
