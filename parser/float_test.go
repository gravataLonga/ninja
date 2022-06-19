package parser

import (
	"ninja/ast"
	"ninja/lexer"
	"testing"
)

// Todo test for float overflow or other errors
func TestFloatLiteralExpression(t *testing.T) {
	input := `
5.0;
0.20;
1.23;
1000.34;
20.00000000012;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 5 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		identifier   float64
		tokenLiteral string
	}{
		{5.0, "5.0"},
		{0.20, "0.20"},
		{1.23, "1.23"},
		{1000.34, "1000.34"},
		{20.0000000001, "20.00000000012"},
	}

	for i, tt := range tests {
		stmt, ok := program.Statements[i].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[%d] is not ast.ExpressionStatement. got=%T", i, program.Statements[i])
		}

		literal, ok := stmt.Expression.(*ast.FloatLiteral)
		if !ok {
			t.Fatalf("exp not *ast.FloatLiteral. got=%T", stmt.Expression)
		}

		if literal.Value != tt.identifier {
			t.Errorf("literal.Value isnt same %.100f. got=%.100f", tt.identifier, literal.Value)
		}

		if literal.TokenLiteral() != tt.tokenLiteral {
			t.Errorf("literal.TokenLiteral not %.2f. got=%s", tt.identifier, literal.TokenLiteral())
		}
	}
}
