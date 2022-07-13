package parser

import (
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"strings"
	"testing"
)

func TestEnumStatement(t *testing.T) {

	input := `enum math {
	case PI: 3.1415;
	case EPSILON: 0.000000001;
};`

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0]
	enum, ok := stmt.(*ast.EnumStatement)
	if !ok {
		t.Fatalf("exp is not ast.EnumStatement. got=%T", stmt)
	}

	if _, ok := enum.Branches["PI"]; !ok {
		t.Fatalf("Branch PI not found")
	}

	if _, ok := enum.Branches["EPSILON"]; !ok {
		t.Fatalf("Branch PI not found")
	}

	pi, _ := enum.Branches["PI"]
	epsilon, _ := enum.Branches["EPSILON"]

	testIdentifier(t, enum.Identifier, "math")
	testFloatLiteral(t, pi, 3.1415)
	testFloatLiteral(t, epsilon, 0.000000001)
}

func TestEnumAccessorValue(t *testing.T) {
	input := `math::PI`

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[%d] is not ast.ExpressionStatement. got=%T", 0, program.Statements[0])
	}

	expression, ok := stmt.Expression.(*ast.ScopeOperatorExpression)
	if !ok {
		t.Fatalf("exp not *ast.ScopeOperatorExpression. got=%T", stmt.Expression)
	}

	testIdentifier(t, expression.AccessIdentifier, "math")
	testIdentifier(t, expression.PropertyIdentifier, "PI")
}

func TestEnumAccessorValueWrong(t *testing.T) {
	input := `math::::`

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	p.ParseProgram()

	if len(p.Errors()) <= 0 {
		t.Fatalf("expected got at least 1 error. Got: 0")
	}

	if p.Errors()[0] != "expected next token to be IDENT, got :: at [Line: 1, Offset: 8] instead." {
		t.Fatalf("Expected error to be %s. Got: %s", "expected next token to be IDENT, got :: (::) at [Line: 1, Offset: 8] instead.", p.Errors()[0])
	}
}

func TestEnumStatementWrong(t *testing.T) {
	tests := []struct {
		input         string
		expectedError string
	}{
		{
			`enum t { case T: 1; case T: 1 }`,
			`Fatal error: Cannot redefine identifier T`,
		},
		{
			`enum { case T: 1; case T: 1 }`,
			`expected next token to be IDENT, got { at [Line: 1, Offset: 6] instead.`,
		},
		{
			`enum T case T: 1; case T: 1 }`,
			`expected next token to be {, got CASE at [Line: 1, Offset: 12] instead.`,
		},
		{
			`enum T { T: 1; case T: 1 }`,
			`expected next token to be CASE, got IDENT at [Line: 1, Offset: 11] instead.`,
		},
		{
			`enum T { case T 1; case T: 1 }`,
			`expected next token to be :, got INT at [Line: 1, Offset: 18] instead.`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestEnumStatementWrong[%d]", i), func(t *testing.T) {
			l := lexer.New(strings.NewReader(tt.input))
			p := New(l)
			p.ParseProgram()

			if len(p.Errors()) <= 0 {
				t.Fatalf("expected at least 1 error. Got: 0")
			}

			if p.Errors()[0] != tt.expectedError {
				t.Errorf("error message expected to be %s. Got: %s", tt.expectedError, p.Errors()[0])
			}

		})
	}
}
