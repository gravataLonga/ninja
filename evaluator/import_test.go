package evaluator

import (
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
		v := testEval(tt.input)
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
			"IO Error: error reading file 'non-exists-file': open non-exists-file: no such file or directory",
		},
		{
			`import "../fixtures/stub-with-error.nj"`,
			"../fixtures/stub-with-error.nj: expected next token to be (, got EOF instead. [line: 1, character: 14]",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}
