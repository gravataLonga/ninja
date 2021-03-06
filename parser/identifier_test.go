package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"strings"
	"testing"
)

func TestIdentifierExpression(t *testing.T) {
	input := `
foobar;
testing_other;
with_number_123;
Abc_123;
`

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 4 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		identifier string
	}{
		{"foobar"},
		{"testing_other"},
		{"with_number_123"},
		{"Abc_123"},
	}

	for i, tt := range tests {
		stmt, ok := program.Statements[i].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[%d] is not ast.ExpressionStatement. got=%T", i, program.Statements[i])
		}

		ident, ok := stmt.Expression.(*ast.Identifier)
		if !ok {
			t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
		}

		if ident.Value != tt.identifier {
			t.Errorf("ident.Value not %s. got=%s", tt.identifier, ident.Value)
		}

		if ident.TokenLiteral() != tt.identifier {
			t.Errorf("ident.TokenLiteral not %s. got=%s", tt.identifier, ident.TokenLiteral())
		}
	}
}
