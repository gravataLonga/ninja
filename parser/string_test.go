package parser

import (
	"ninja/ast"
	"ninja/lexer"
	"testing"
)

func TestStringExpression(t *testing.T) {
	tests := []struct {
		input          string
		expectedString string
	}{
		{"\"Testing\"", "Testing"},
		{"\"\"", ""},
		{"\" \"", " "},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		str, ok := stmt.Expression.(*ast.StringLiteral)
		if !ok {
			t.Fatalf("exp not *ast.String. got=%T", stmt.Expression)
		}

		if str.Value != tt.expectedString {
			t.Errorf("boolean.Value not %s. got=%s", tt.expectedString, str.Value)
		}
	}
}
