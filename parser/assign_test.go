package parser

import (
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/token"
	"strings"
	"testing"
)

func TestVarStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"var x = 5;", "x", 5},
		{"var y = true;", "y", true},
		{"var foobar = y;", "foobar", "y"},
		{"var barfoo = 12.3;", "barfoo", 12.3},
	}

	for _, tt := range tests {
		l := lexer.New(strings.NewReader(tt.input))
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.VarStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

func TestAssignStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"x = y;", "x", "y"},
		{"x = 1;", "x", 1},
		{"x = true;", "x", true},
		{"x = false;", "x", false},
		{"x = 13.3;", "x", 13.3},
		// {"x[0] = 10;", "x", 10},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestAssignStatements[%d]", i), func(t *testing.T) {
			l := lexer.New(strings.NewReader(tt.input))
			p := New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			if len(program.Statements) != 1 {
				t.Fatalf("program.Statements does not contain 1 statements. got=%d",
					len(program.Statements))
			}

			stmt := program.Statements[0]
			if !testAssingStatement(t, stmt, tt.expectedIdentifier) {
				return
			}

			val := stmt.(*ast.AssignStatement).Right
			if !testLiteralExpression(t, val, tt.expectedValue) {
				return
			}
		})

	}
}

func TestAssignIndexStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      string
	}{
		{"x[0] = 1;", "(x[0])", "1"},
		{"x[\"key\"] = 1;", "(x[key])", "1"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestAssignIndexStatements[%d]", i), func(t *testing.T) {
			l := lexer.New(strings.NewReader(tt.input))
			p := New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			if len(program.Statements) != 1 {
				t.Fatalf("program.Statements does not contain 1 statements. got=%d",
					len(program.Statements))
			}

			stmt := program.Statements[0]
			assign, ok := stmt.(*ast.AssignStatement)
			if !ok {
				t.Fatalf("program.Statements[0] does not equal to AssignStatement. got=%T", stmt)
			}

			exprStmt, ok := assign.Left.(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("program.Statements[0].Left does not equal to ExpressionStatement. got=%T", assign.Left)
			}

			indexStmt, ok := exprStmt.Expression.(*ast.IndexExpression)
			if !ok {
				t.Fatalf("program.Statements[0].Left.Expression does not equal to IndexExpression. got=%T", exprStmt.Expression)
			}

			if indexStmt.String() != tt.expectedIdentifier {
				t.Fatalf("indexExpr must be %s. Got: %s", tt.expectedIdentifier, indexStmt.String())
			}

			if assign.Right.String() != tt.expectedValue {
				t.Fatalf("assign.Right must be %s. Got: %s", tt.expectedValue, assign.Right.String())
			}
		})

	}
}

func TestAssignExpression(t *testing.T) {
	l := lexer.New(strings.NewReader(`var a = a + 1; a = a + 1; a;`))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	stmt := program.Statements[1]
	if !testAssingStatement(t, stmt, "a") {
		return
	}

}

func TestVarStatementErrors(t *testing.T) {
	input := `
var x true;
var = "123";
var var;
var = =;
i = i = 1;
`

	l := lexer.New(strings.NewReader(input))
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(p.errors) != 6 {
		t.Fatalf("Expected 6 errors, got: %d", len(p.errors))
	}

	tests := []struct {
		expectedError string
	}{
		{fmt.Sprintf("expected next token to be %s, got %s at [Line: 2, Offset: 11] instead.", token.ASSIGN, token.TRUE)},
		{fmt.Sprintf("expected next token to be %s, got %s at [Line: 3, Offset: 5] instead.", token.IDENT, token.ASSIGN)},
		{fmt.Sprintf("expected next token to be %s, got %s at [Line: 4, Offset: 8] instead.", token.IDENT, token.VAR)},
		{fmt.Sprintf("expected next token to be %s, got %s at [Line: 5, Offset: 5] instead.", token.IDENT, token.ASSIGN)},
		{fmt.Sprintf("expected next token to be %s, got %s at [Line: 6, Offset: 7] instead.", token.IDENT, token.ASSIGN)},
	}

	errors := p.Errors()

	for i, tt := range tests {
		err := errors[i]
		if err != tt.expectedError {
			t.Errorf("Error \"%s\" got=%q", tt.expectedError, err)
		}
	}
}

/*
@todo we need to validate when we are assign invalid objects
func TestIllegalAssignmentsErrors(t *testing.T) {
	tests := []struct {
		input         string
		expectedError string
	}{
		{
			`"ola" = "ola"`,
			`illegal "ola" assignment to "ola"`,
		},
		{
			`function() {} = "ola"`,
			`illegal "ola" assignment to "function"`,
		},
		{
			`1 = function() {}`,
			`illegal "function" assignment to "1"`,
		},
		{
			`{} = 1`,
			`illegal "1" assignment to "{"`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestIllegalAssignmentsErrors[%d]", i), func(t *testing.T) {
			l := lexer.New(strings.NewReader(tt.input))
			p := New(l)

			program := p.ParseProgram()

			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}

			errors := p.Errors()

			if len(errors) <= 0 {
				t.Fatalf("Program don't produce any error %s", tt.input)

			}

			if tt.expectedError != errors[0] {
				t.Fatalf("TestIllegalAssignmentsErrors expected \"%s\" error. Got: \"%s\"", tt.expectedError, errors[0])
			}
		})

	}
}
*/
