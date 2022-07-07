package evaluator

import (
	"fmt"
	"ninja/object"
	"testing"
)

func TestImportStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`import "../fixtures/stub.nj"; add(1, 1);`, 2},
		{`var a = import "../fixtures/stub.nj"; a;`, nil},
		{`var a = import "../fixtures/stub_return.nj"; a;`, 2},
	}

	for _, tt := range tests {
		v := testEval(tt.input, t)
		testObjectLiteral(t, v, tt.expected)
	}
}

func TestErrorImportHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			`import "non-exists-file"`,
			"IO Error: error reading file 'non-exists-file': open non-exists-file: no such file or directory IMPORT at [Line: 1, Offset: 7]",
		},
		{
			`import "../fixtures/stub-with-error.nj"`,
			"../fixtures/stub-with-error.nj: expected next token to be (, got EOF at [Line: 1, Offset: 13] instead.",
		},
		{
			`import "../fixtures/stub-with-error-in-function.nj"`,
			"../fixtures/stub-with-error-in-function.nj: Function expected 2 arguments, got 3 at { at [Line: 16, Offset: 38]",
		},
	}

	for i, tt := range tests {

		t.Run(fmt.Sprintf("TestErrorImportHandling_%d", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

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
