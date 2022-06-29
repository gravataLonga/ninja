package parser

import (
	"ninja/ast"
	"ninja/lexer"
	"strings"
	"testing"
)

func TestObjectCallIntegerExpression(t *testing.T) {
	input := "1.type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	objectCall, ok := stmt.Expression.(*ast.ObjectCall)

	if !ok {
		t.Fatalf("exp not *ast.ObjectCall. got=%T", stmt.Expression)
	}

	callExp, ok := objectCall.Call.(*ast.CallExpression)

	if !ok {
		t.Fatalf("objectCall.Call is not CallExpression. got=%T", objectCall.Call)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, callExp.Function, "type")
	testIntegerLiteral(t, objectCall.Object, 1)
}

func TestObjectCallBooleanExpression(t *testing.T) {
	input := "true.type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	objectCall, ok := stmt.Expression.(*ast.ObjectCall)

	if !ok {
		t.Fatalf("exp not *ast.ObjectCall. got=%T", stmt.Expression)
	}

	callExp, ok := objectCall.Call.(*ast.CallExpression)

	if !ok {
		t.Fatalf("objectCall.Call is not CallExpression. got=%T", objectCall.Call)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, callExp.Function, "type")
	testBooleanLiteral(t, objectCall.Object, true)
}

func TestObjectCallStringExpression(t *testing.T) {
	input := "\"hello\".type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	objectCall, ok := stmt.Expression.(*ast.ObjectCall)

	if !ok {
		t.Fatalf("exp not *ast.ObjectCall. got=%T", stmt.Expression)
	}

	callExp, ok := objectCall.Call.(*ast.CallExpression)

	if !ok {
		t.Fatalf("objectCall.Call is not CallExpression. got=%T", objectCall.Call)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, callExp.Function, "type")
	testStringLiteral(t, objectCall.Object, "hello")
}

func TestObjectCallArrayExpression(t *testing.T) {
	input := "[].type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	objectCall, ok := stmt.Expression.(*ast.ObjectCall)

	if !ok {
		t.Fatalf("exp not *ast.ObjectCall. got=%T", stmt.Expression)
	}

	callExp, ok := objectCall.Call.(*ast.CallExpression)

	if !ok {
		t.Fatalf("objectCall.Call is not CallExpression. got=%T", objectCall.Call)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, callExp.Function, "type")
	testArrayLiteral(t, objectCall.Object, "[]")
}

func TestObjectCallHashExpression(t *testing.T) {
	input := "{}.type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	objectCall, ok := stmt.Expression.(*ast.ObjectCall)

	if !ok {
		t.Fatalf("exp not *ast.ObjectCall. got=%T", stmt.Expression)
	}

	callExp, ok := objectCall.Call.(*ast.CallExpression)

	if !ok {
		t.Fatalf("objectCall.Call is not CallExpression. got=%T", objectCall.Call)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, callExp.Function, "type")
	testHashLiteral(t, objectCall.Object, "{}")
}

func TestObjectCallExpression_Multiple(t *testing.T) {
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
		objectCall, ok := stmt.Expression.(*ast.ObjectCall)

		if !ok {
			t.Fatalf("exp not *ast.ObjectCall. got=%T", stmt.Expression)
		}

		testLiteralExpression(t, objectCall.Object, tt.expectedIdentifier)

		callExp, ok := objectCall.Call.(*ast.CallExpression)

		if !ok {
			t.Fatalf("objectCall.Call is not CallExpression. got=%T", objectCall.Call)
		}

		for i, arg := range tt.expectedArgs {
			if callExp.Arguments[i].String() != arg {
				t.Errorf("argument %d wrong. want=%q, got=%q", i, arg, callExp.Arguments[i].String())
			}
		}
	}
}
