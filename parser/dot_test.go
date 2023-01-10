package parser

import (
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"strings"
	"testing"
)

func TestDotCallIntegerExpression(t *testing.T) {
	input := "1.type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	dotCall, ok := stmt.Expression.(*ast.Dot)

	if !ok {
		t.Fatalf("exp not *ast.Dot. got=%T", stmt.Expression)
	}

	callExp, ok := dotCall.Right.(*ast.CallExpression)

	if !ok {
		t.Fatalf("dotCall.Right is not CallExpression. got=%T", dotCall.Right)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, callExp.Function, "type")
	testIntegerLiteral(t, dotCall.Object, 1)
}

func TestDotCallBooleanExpression(t *testing.T) {
	input := "true.type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	dotCall, ok := stmt.Expression.(*ast.Dot)

	if !ok {
		t.Fatalf("exp not *ast.Dot. got=%T", stmt.Expression)
	}

	callExp, ok := dotCall.Right.(*ast.CallExpression)

	if !ok {
		t.Fatalf("dotCall.Right is not CallExpression. got=%T", dotCall.Right)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, callExp.Function, "type")
	testBooleanLiteral(t, dotCall.Object, true)
}

func TestDotCallStringExpression(t *testing.T) {
	input := "\"hello\".type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	dotCall, ok := stmt.Expression.(*ast.Dot)

	if !ok {
		t.Fatalf("exp not *ast.Dot. got=%T", stmt.Expression)
	}

	callExp, ok := dotCall.Right.(*ast.CallExpression)

	if !ok {
		t.Fatalf("dotCall.Right is not CallExpression. got=%T", dotCall.Right)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, callExp.Function, "type")
	testStringLiteral(t, dotCall.Object, "hello")
}

func TestDotCallArrayExpression(t *testing.T) {
	input := "[].type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	dotCall, ok := stmt.Expression.(*ast.Dot)

	if !ok {
		t.Fatalf("exp not *ast.Dot. got=%T", stmt.Expression)
	}

	callExp, ok := dotCall.Right.(*ast.CallExpression)

	if !ok {
		t.Fatalf("dotCall.Right is not CallExpression. got=%T", dotCall.Right)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, callExp.Function, "type")
	testArrayLiteral(t, dotCall.Object, "[]")
}

func TestDotCallHashExpression(t *testing.T) {
	input := "{}.type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	dotCall, ok := stmt.Expression.(*ast.Dot)

	if !ok {
		t.Fatalf("exp not *ast.Dot. got=%T", stmt.Expression)
	}

	callExp, ok := dotCall.Right.(*ast.CallExpression)

	if !ok {
		t.Fatalf("dotCall.Right is not CallExpression. got=%T", dotCall.Right)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, callExp.Function, "type")
	testHashLiteral(t, dotCall.Object, "{}")
}

func TestDotCallExpression_Multiple(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier interface{}
		expectedArgs       []string
	}{
		{
			"1.type()",
			1,
			[]string{},
		},
		{
			"a.type()",
			"a",
			[]string{},
		},
		{
			"true.type()",
			true,
			[]string{},
		},
		{
			"a.type(a)",
			"a",
			[]string{"a"},
		},
		{
			"a.type(a, b)",
			"a",
			[]string{"a", "b"},
		},
		{
			"a.type(a, b, c.type())",
			"a",
			[]string{"a", "b", "(c.type())"},
		},
	}

	for _, tt := range tests {
		l := lexer.New(strings.NewReader(tt.input))
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		dotCall, ok := stmt.Expression.(*ast.Dot)

		if !ok {
			t.Fatalf("exp not *ast.Dot. got=%T", stmt.Expression)
		}

		testLiteralExpression(t, dotCall.Object, tt.expectedIdentifier)

		callExp, ok := dotCall.Right.(*ast.CallExpression)

		if !ok {
			t.Fatalf("dotCall.Right is not CallExpression. got=%T", dotCall.Right)
		}

		for i, arg := range tt.expectedArgs {
			if callExp.Arguments[i].String() != arg {
				t.Errorf("argument %d wrong. want=%q, got=%q", i, arg, callExp.Arguments[i].String())
			}
		}
	}
}
