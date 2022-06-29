package parser

import (
	"ninja/ast"
	"ninja/lexer"
	"strconv"
	"strings"
	"testing"
)

func TestIntegerLiteralExpression(t *testing.T) {
	input := `
5;
100000;
100000000000;
10000000000000000;
`

	l := lexer.New(strings.NewReader(input))
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 4 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		identifier int64
	}{
		{5},
		{100000},
		{100000000000},
		{10000000000000000},
	}

	for i, tt := range tests {
		stmt, ok := program.Statements[i].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[%d] is not ast.ExpressionStatement. got=%T", i, program.Statements[i])
		}

		literal, ok := stmt.Expression.(*ast.IntegerLiteral)
		if !ok {
			t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
		}

		if literal.Value != tt.identifier {
			t.Errorf("literal.Value not %d. got=%d", tt.identifier, literal.Value)
		}

		if literal.TokenLiteral() != strconv.FormatInt(tt.identifier, 10) {
			t.Errorf("literal.TokenLiteral not %d. got=%s", tt.identifier, literal.TokenLiteral())
		}
	}
}

func TestIntegerLiteralExpressionOverflowError(t *testing.T) {
	input := `
10000000000000000000;
`

	l := lexer.New(strings.NewReader(input))
	p := New(l)

	p.ParseProgram()

	if len(p.Errors()) <= 0 {
		t.Fatalf("program not caught overflow int.")
	}

	errors := p.Errors()

	if errors[0] != "could not parse \"10000000000000000000\" as integer" {
		t.Errorf("Error \"%s\" got=%q", "could not parse \"10000000000000000000\" as integer", errors[0])
	}
}
