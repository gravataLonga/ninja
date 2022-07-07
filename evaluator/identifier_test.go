package evaluator

import (
	"fmt"
	"ninja/object"
	"testing"
)

func TestVarStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"var a = 5; a;", 5},
		{"var a = \"Hello World\"; a;", "Hello World"},
		{"var a = 5.5; a;", 5.5},
		{"var a = true; a;", true},
		{"var a = false; a;", false},
		{"var a = !false; a;", true},
		{"var a = 1 + 1 * 3; a;", 4},
		{"var a = 1; a = a + 1; a;", 2},
		{"var a = 1; var b = 2; a = a + b + 1; a;", 4},
		{"var a = function () { return 1;}; a();", 1},
		{"var a = 5 * 5; a;", 25},
		{"var a = 5; var b = a; b;", 5},
		{"var a = 5; var b = a; var c = a + b + 5; c;", 15},
		{"var a = 5; ++a;", 6},
		{"var a = 5; --a;", 4},
		{"var a = 5; a = a + 5; a;", 10},
		{"var a = 5; var b = 10; a = a + b * a; a", 55},
		{"var a = 10; a = \"hello\"; a;", "hello"},
		{"var a = 23.3; a;", 23.3},
		{"var a = {}; a[0] = true; a[0];", true},
		{"var a = [false]; a[0] = true; a[0];", true},
		// {"var a = 5; a++;", 5},
		// {"var a = 5; a--;", 5},
	}

	for _, tt := range tests {
		testObjectLiteral(t, testEval(tt.input, t), tt.expected)
	}
}

func TestErrorIdentifierHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"foobar",
			"identifier not found: foobar IDENT at [Line: 1, Offset: 7]",
		},
		{
			"foobar = 1 + 1;",
			"identifier not found: foobar IDENT at [Line: 1, Offset: 7]",
		},
		{
			"var b = a + 1;",
			"identifier not found: a IDENT at [Line: 1, Offset: 10]",
		},
		{
			`"Hello" - "World"`,
			"unknown operator: STRING - STRING",
		},
		{
			`"Hello" > "World"`,
			"unknown operator: STRING > STRING",
		},
		{
			`"Hello" < "World"`,
			"unknown operator: STRING < STRING",
		},
		{
			`"Hello" >= "World"`,
			"unknown operator: STRING >= STRING",
		},
		{
			`"Hello" <= "World"`,
			"unknown operator: STRING <= STRING",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestErrorIdentifierHandling[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedMessage {
				t.Errorf("wrong error message. expected=%q, got=%q",
					tt.expectedMessage, errObj.Message)
			}
		})

	}
}
