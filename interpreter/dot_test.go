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

func TestDotErrorMasking(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{`(1 + "a").foo`, "unknown operator: INTEGER + STRING at [Line: 1, Offset: 4]"},
		{`(1 - "a").bar`, "unknown operator: INTEGER - STRING at [Line: 1, Offset: 4]"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestDotErrorMasking[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}
			if errObj.Message != tt.expectedMessage {
				t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
			}
		})
	}
}
