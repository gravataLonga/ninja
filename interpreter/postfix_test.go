package interpreter_test

import (
	"fmt"
	"github.com/gravataLonga/ninja/object"
	"testing"
)

func TestPostfixOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`1++`,
			2,
		},
		{
			`0++`,
			1,
		},
		{
			`0.0++`,
			1.0,
		},
		{
			`1.0++`,
			2.0,
		},

		{
			`1--`,
			0,
		},
		{
			`0--`,
			-1,
		},
		{
			`0.0--`,
			-1.0,
		},
		{
			`1.0--`,
			0.0,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestPostfixOperator[%d]", i), func(t *testing.T) {

			v := evalProgram(t, tt.input)

			if v == nil {
				t.Fatalf("Interpreter return nil as result")
			}

			if _, ok := v.(*object.Error); ok {
				t.Fatalf("Interpreter return error. %s", v.Inspect())
			}

			if !testLiteralObject(t, v, tt.expected) {
				t.Fatalf("testLiteralObject got false, expected true.")
			}
		})
	}
}

func TestPostfixOperatorErrors(t *testing.T) {
	tests := []struct {
		input string
	}{
		{`var s = "hello"; s++;`},
		{`var b = true; b++;`},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestPostfixOperatorErrors[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
			if _, ok := evaluated.(*object.Error); !ok {
				t.Fatalf("expected *object.Error, got=%T(%+v)", evaluated, evaluated)
			}
		})
	}
}
