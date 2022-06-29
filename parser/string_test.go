package parser

import (
	"ninja/ast"
	"ninja/lexer"
	"strings"
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
		l := lexer.New(strings.NewReader(tt.input))
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

func TestStringIndexExpression(t *testing.T) {
	tests := []struct {
		input          string
		expectedString string
		index          int64
	}{
		{"\"\"[0]", "", 0},
		{"\"1\"[0]", "1", 0},
		{"\"Hello World\"[2]", "Hello World", 2},
	}

	for _, tt := range tests {
		l := lexer.New(strings.NewReader(tt.input))
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("TestStringIndexExpression program has not enough statements. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("TestStringIndexExpression program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		index, ok := stmt.Expression.(*ast.IndexExpression)
		if !ok {
			t.Fatalf("TestStringIndexExpression is not *ast.IndexExpression. got=%T", stmt.Expression)
		}

		str, ok := index.Left.(*ast.StringLiteral)
		if !ok {
			t.Fatalf("TestStringIndexExpression left expression is not *ast.StringLiteral. got=%T", stmt.Expression)
		}

		if str.Value != tt.expectedString {
			t.Errorf("boolean.Value not %s. got=%s", tt.expectedString, str.Value)
		}

		ix, ok := index.Index.(*ast.IntegerLiteral)
		if !ok {
			t.Fatalf("TestStringIndexExpression index expression is not *ast.IntegerLiteral. got=%T", stmt.Expression)
		}

		if ix.Value != tt.index {
			t.Fatalf("TestStringIndexExpression index expression is not %d. Got: %d", tt.index, ix.Value)
		}
	}
}
