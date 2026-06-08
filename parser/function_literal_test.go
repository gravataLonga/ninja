package parser

import (
	"strings"
	"testing"

	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
)

func TestFunctionLiteralParsing(t *testing.T) {
	input := `function (x, y) { return x + y; }`

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.FunctionLiteral. got=%T", stmt.Expression)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n", len(function.Parameters))
	}

	testIdentifier(t, function.Parameters[0].Name, "x")
	if function.Parameters[0].Default != nil {
		t.Errorf("function.Parameters[0].Default should be nil, got=%d\n", function.Parameters[0].Default)
	}
	testIdentifier(t, function.Parameters[1].Name, "y")
	if function.Parameters[1].Default != nil {
		t.Errorf("function.Parameters[1].Default should be nil, got=%d\n", function.Parameters[0].Default)
	}

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements. got=%d\n", len(function.Body.Statements))
	}

	_, ok = function.Body.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ExpressionStatement. got=%T", function.Body.Statements[0])
	}

	// testInfixExpression(t, bodyStmt.ReturnValue.(*ast.InfixExpression), "x", "+", "y")
}
