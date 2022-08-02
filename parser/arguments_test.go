package parser

import (
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"strings"
	"testing"
)

func TestFunctionLiteralParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "function () {};", expectedParams: []string{}},
		{input: "function (x) {};", expectedParams: []string{"x"}},
		{input: "function (x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		l := lexer.New(strings.NewReader(tt.input))
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		function := stmt.Expression.(*ast.FunctionLiteral)

		if len(function.Parameters) != len(tt.expectedParams) {
			t.Errorf("length parameters wrong. want %d, got=%d\n", len(tt.expectedParams), len(function.Parameters))
		}

		for i, ident := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
		ident          string
	}{
		{input: "function add() {};", expectedParams: []string{}, ident: "add"},
		{input: "function sub(x) {};", expectedParams: []string{"x"}, ident: "sub"},
		{input: "function avg(x, y, z) {};", expectedParams: []string{"x", "y", "z"}, ident: "avg"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestFunctionParameter %s", tt.input), func(t *testing.T) {
			l := lexer.New(strings.NewReader(tt.input))
			p := New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			stmt := program.Statements[0].(*ast.ExpressionStatement)
			function := stmt.Expression.(*ast.Function)

			if len(function.Parameters) != len(tt.expectedParams) {
				t.Errorf("length parameters wrong. want %d, got=%d\n", len(tt.expectedParams), len(function.Parameters))
			}

			for i, ident := range tt.expectedParams {
				testLiteralExpression(t, function.Parameters[i], ident)
			}

			if function.Name.String() != tt.ident {
				t.Errorf("Identitier of function not match. want %s, got=%s\n", tt.ident, function.Name.String())
			}
		})
	}
}

func TestFunctionParameterOptionalParsing(t *testing.T) {
	l := lexer.New(strings.NewReader(`function (x, y = 0, z = "hello", k = true, a = b) {}`))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program.String() != "function(x, (y = 0), (z = hello), (k = true), (a = b)) " {
		t.Fatalf("program didn't produce expected string, got: %q", program.String())
	}

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	fn, ok := stmt.Expression.(*ast.FunctionLiteral)

	if !ok {
		t.Fatalf("Expression isn't FunctionLiteral. Got: %T", stmt.Expression)
	}

	if len(fn.Parameters) != 5 {
		t.Fatalf("Arguments of function isn't equal 4. Got: %d", len(fn.Parameters))
	}

	testInfixExpression(t, fn.Parameters[1], "y", "=", 0)
	testInfixExpression(t, fn.Parameters[2], "z", "=", "hello")
	testInfixExpression(t, fn.Parameters[3], "k", "=", true)
	testInfixExpression(t, fn.Parameters[4], "a", "=", "b")
}
