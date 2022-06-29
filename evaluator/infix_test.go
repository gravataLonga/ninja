package evaluator

import (
	"ninja/lexer"
	"ninja/object"
	"ninja/parser"
	"strings"
	"testing"
)

func TestErrorParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input         string
		expectedError string
	}{
		{"5 + * 5;", "no prefix parse function for * found"},
	}

	for _, tt := range infixTests {

		l := lexer.New(strings.NewReader(tt.input))
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()

		_ = Eval(program, env)

		if len(p.Errors()) <= 0 {
			t.Fatalf("Program wasn't returning any error.")
		}

		if p.Errors()[0] != tt.expectedError {
			t.Errorf("Expected: %s. Got: %s", tt.expectedError, p.Errors())
		}
	}
}
