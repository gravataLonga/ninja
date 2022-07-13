package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"strings"
	"testing"
)

func TestBreakStatement(t *testing.T) {
	input := `break;`

	l := lexer.New(strings.NewReader(input))
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 0 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		breakStatement, ok := stmt.(*ast.BreakStatement)
		if !ok {
			t.Errorf("stmt not *ast.BreakStatement. got=%T", stmt)
			continue
		}

		if breakStatement.TokenLiteral() != "break" {
			t.Errorf("breakStatement.TokenLiteral not 'break', got %q", breakStatement.TokenLiteral())
		}
	}
}
