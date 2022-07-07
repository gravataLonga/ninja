package parser

import (
	"fmt"
	"ninja/ast"
	"ninja/lexer"
	"ninja/token"
	"strings"
	"testing"
)

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	return 10.20;
	return foobar;
	return;
	return import "ola";
	return 3 <= 5;
	return "ola" == "ola";
   `

	l := lexer.New(strings.NewReader(input))
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 9 {
		t.Fatalf("program.Statements does not contain 6 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStatemnt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}

		if returnStatemnt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStatemnt.TokenLiteral())
		}
	}
}

func TestReturnStatementErrors(t *testing.T) {
	input := `
return var;
return return;
return --;
return ++;
return <=;
return >=;
`

	l := lexer.New(strings.NewReader(input))
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(p.errors) != 6 {
		t.Errorf("Expected 6 errors, got: %d", len(p.errors))
	}

	tests := []struct {
		expectedError string
	}{
		{fmt.Sprintf("Next token expected to be nil or expression. Got: %s at [Line: 2, Offset: 11].", token.VAR)},
		{fmt.Sprintf("Next token expected to be nil or expression. Got: %s at [Line: 3, Offset: 14].", token.RETURN)},
		{fmt.Sprintf("Next token expected to be nil or expression. Got: %s at [Line: 4, Offset: 9].", token.DECRE)},
	}

	errors := p.Errors()

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestReturnStatementErrors_%d", i), func(t *testing.T) {
			err := errors[i]
			if err != tt.expectedError {
				t.Errorf("Error \"%s\" got=%s", tt.expectedError, err)
			}
		})

	}
}
