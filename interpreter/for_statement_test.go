package interpreter_test

import (
	"fmt"
	"testing"

	"github.com/gravataLonga/ninja/object"
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
		{
			`var skipped = 0; for(var i = 0; i < 5; i = i + 1) { if(i == 2) { continue; } skipped = skipped + 1; }; skipped;`,
			4,
		},
		{
			`var count = 0; for(var i = 0; i < 5; i = i + 1) { if(i > 1) { continue; } count = count + 1; }; count;`,
			2,
		},
		{
			`var x = 0; for(var i = 0; i < 3; i = i + 1) { continue; x = x + 1; }; x;`,
			0,
		},
		{
			`var last = 0; for(var i = 0; i < 5; i = i + 1) { continue; last = i; }; last;`,
			0,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestForStatement[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)

			integer, ok := tt.expected.(int)
			if ok {
				testIntegerObject(t, evaluated, int64(integer))
			} else if evaluated.Type() != object.NULL_OBJ {
				t.Errorf("result isnt null. Got %v", evaluated)
			}
		})
	}
}

func TestLoopStopWhenFoundError(t *testing.T) {
	input := `for(var i = 0; i <= 2; i = i + 1) { i; if (i == 1) { 1 + "ola"; } }`
	expected := "unknown operator: INTEGER + STRING at [Line: 1, Offset: 56]"

	evaluated := evalProgram(t, input)

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

func TestBreakOutsideForLoop(t *testing.T) {
	input := `break`
	expected := "'break' not in the 'loop' context"

	evaluated := evalProgram(t, input)

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

func TestContinueOutsideForLoop(t *testing.T) {
	input := `continue`
	expected := "'continue' not in the 'loop' context"

	evaluated := evalProgram(t, input)

	if evaluated == nil {
		t.Fatalf("evaluated is empty")
	}

	err, ok := evaluated.(*object.Error)

	if !ok {
		t.Fatalf("expected error. Got: %T(%+v)", evaluated, evaluated)
	}

	if err.Message != expected {
		t.Fatalf("expected error message to be %s, got: %s", expected, err.Message)
	}
}
