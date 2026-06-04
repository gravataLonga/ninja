package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"strings"
	"testing"
)

func TestContinueStatement(t *testing.T) {
	input := `continue;`

	l := lexer.New(strings.NewReader(input))
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 0 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		continueStatement, ok := stmt.(*ast.ContinueStatement)
		if !ok {
			t.Errorf("stmt not *ast.ContinueStatement. got=%T", stmt)
			continue
		}

		if continueStatement.TokenLiteral() != "continue" {
			t.Errorf("continueStatement.TokenLiteral not 'continue', got %q", continueStatement.TokenLiteral())
		}
	}
}
