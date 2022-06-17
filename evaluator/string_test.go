package evaluator

import (
	"ninja/object"
	"testing"
)

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestEvalStringExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{

		{`"Hello"`, "Hello"},
		{`"Hello" + "World"`, "HelloWorld"},
		{`"Hello" && "World"`, true},
		{`"Hello" && false`, false},
		{`"Hello" || false`, true},
		{`!"Hello"`, false},
		{`"Hello" == "Hello"`, true},
		{`"Hello" != "Hello"`, false},
		{`"Hello" == "World"`, false},
		{`"Hello" != "World"`, true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testObjectLiteral(t, evaluated, tt.expected)
	}
}

func TestErrorStringHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"-\"hello\"",
			"unknown operator: -STRING",
		},
		{
			`"Hello" - "Nice"`,
			"unknown operator: STRING - STRING",
		},
		{
			`"Hello" * "Nice"`,
			"unknown operator: STRING * STRING",
		},
		{
			`"Hello" / "Nice"`,
			"unknown operator: STRING / STRING",
		},
		{
			`++"Nice"`,
			"unknown object type STRING for operator ++",
		},
		{
			`--"Nice"`,
			"unknown object type STRING for operator --",
		},
		{
			`"1" < "2"`,
			"unknown operator: STRING < STRING",
		},
		{
			`"1" > "2"`,
			"unknown operator: STRING > STRING",
		},
		{
			`"1" <= "2"`,
			"unknown operator: STRING <= STRING",
		},
		{
			`"1" >= "2"`,
			"unknown operator: STRING >= STRING",
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
