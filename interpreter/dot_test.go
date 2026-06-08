package interpreter_test

import (
	"fmt"
	"github.com/gravataLonga/ninja/object"
	"testing"
)

func TestDotPropertyAccessOnHash(t *testing.T) {
	tests := []struct {
		input         string
		expected      string
		expectedIsNil bool
	}{
		{
			`var h = {"name": "ninja"}; h.name;`,
			"ninja",
			false,
		},
		{
			`var h = {"name": "ninja"}; h.missing;`,
			"",
			true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestDotPropertyAccessOnHash[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)

			if tt.expectedIsNil {
				testNullObject(t, evaluated)
				return
			}

			testStringObject(t, evaluated, tt.expected)
		})
	}
}

func TestDotPropertyAccessOnNonHash(t *testing.T) {
	input := `var a = [1, 2]; a.first;`

	evaluated := evalProgram(t, input)

	if _, ok := evaluated.(*object.Error); !ok {
		t.Fatalf("expected *object.Error for property access on non-hash. got=%T (%+v)", evaluated, evaluated)
	}
}
