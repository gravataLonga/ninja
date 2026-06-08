package parser

import (
	"strings"
	"testing"

	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
)

func TestDotCallIntegerExpression(t *testing.T) {
	input := "1.type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	callExp, ok := stmt.Expression.(*ast.CallExpression)

	if !ok {
		t.Fatalf("exp not *ast.CallExpression. got=%T", stmt.Expression)
	}

	dotCall, ok := callExp.Function.(*ast.Dot)

	if !ok {
		t.Fatalf("callExp.Function is not *ast.Dot. got=%T", callExp.Function)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, dotCall.Right, "type")
	testIntegerLiteral(t, dotCall.Object, 1)
}

func TestDotCallBooleanExpression(t *testing.T) {
	input := "true.type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	callExp, ok := stmt.Expression.(*ast.CallExpression)

	if !ok {
		t.Fatalf("exp not *ast.CallExpression. got=%T", stmt.Expression)
	}

	dotCall, ok := callExp.Function.(*ast.Dot)

	if !ok {
		t.Fatalf("callExp.Function is not *ast.Dot. got=%T", callExp.Function)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, dotCall.Right, "type")
	testBooleanLiteral(t, dotCall.Object, true)
}

func TestDotCallStringExpression(t *testing.T) {
	input := "\"hello\".type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	callExp, ok := stmt.Expression.(*ast.CallExpression)

	if !ok {
		t.Fatalf("exp not *ast.CallExpression. got=%T", stmt.Expression)
	}

	dotCall, ok := callExp.Function.(*ast.Dot)

	if !ok {
		t.Fatalf("callExp.Function is not *ast.Dot. got=%T", callExp.Function)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, dotCall.Right, "type")
	testStringLiteral(t, dotCall.Object, "hello")
}

func TestDotCallArrayExpression(t *testing.T) {
	input := "[].type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	callExp, ok := stmt.Expression.(*ast.CallExpression)

	if !ok {
		t.Fatalf("exp not *ast.CallExpression. got=%T", stmt.Expression)
	}

	dotCall, ok := callExp.Function.(*ast.Dot)

	if !ok {
		t.Fatalf("callExp.Function is not *ast.Dot. got=%T", callExp.Function)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, dotCall.Right, "type")
	testArrayLiteral(t, dotCall.Object, "[]")
}

func TestDotCallHashExpression(t *testing.T) {
	input := "{}.type()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	callExp, ok := stmt.Expression.(*ast.CallExpression)

	if !ok {
		t.Fatalf("exp not *ast.CallExpression. got=%T", stmt.Expression)
	}

	dotCall, ok := callExp.Function.(*ast.Dot)

	if !ok {
		t.Fatalf("callExp.Function is not *ast.Dot. got=%T", callExp.Function)
	}

	if len(callExp.Arguments) > 0 {
		t.Fatalf("Arguments is not empty")
	}

	testIdentifier(t, dotCall.Right, "type")
	testHashLiteral(t, dotCall.Object, "{}")
}

func TestDotPropertyAccessExpression(t *testing.T) {
	input := "h.name"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	dot, ok := stmt.Expression.(*ast.Dot)
	if !ok {
		t.Fatalf("exp not *ast.Dot. got=%T", stmt.Expression)
	}

	testIdentifier(t, dot.Object, "h")

	if dot.Right.Value != "name" {
		t.Errorf("ident.Value not %q. got=%q", "name", dot.Right.Value)
	}
}

func TestDotMethodCallExpressionShape(t *testing.T) {
	input := "h.name()"

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	// `.` binds tighter than a call: the Dot is resolved first (the property/method
	// is looked up on the object), then the resulting value is invoked.
	call, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("exp not *ast.CallExpression. got=%T", stmt.Expression)
	}

	dot, ok := call.Function.(*ast.Dot)
	if !ok {
		t.Fatalf("call.Function is not *ast.Dot. got=%T", call.Function)
	}

	testIdentifier(t, dot.Object, "h")

	if dot.Right.Value != "name" {
		t.Errorf("ident.Value not %q. got=%q", "name", dot.Right.Value)
	}
}

func TestDotPropertyAssignmentExpression(t *testing.T) {
	input := `h.name = "x";`

	l := lexer.New(strings.NewReader(input))
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.AssignStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.AssignStatement. got=%T", program.Statements[0])
	}

	exprStmt, ok := stmt.Left.(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt.Left is not *ast.ExpressionStatement. got=%T", stmt.Left)
	}

	dot, ok := exprStmt.Expression.(*ast.Dot)
	if !ok {
		t.Fatalf("exprStmt.Expression is not *ast.Dot. got=%T", exprStmt.Expression)
	}

	testIdentifier(t, dot.Object, "h")

	if dot.Right.Value != "name" {
		t.Errorf("ident.Value not %q. got=%q", "name", dot.Right.Value)
	}

	testStringLiteral(t, stmt.Right, "x")
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
			[]string{"a", "b", "(c.type)()"},
		},
	}

	for _, tt := range tests {
		l := lexer.New(strings.NewReader(tt.input))
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		callExp, ok := stmt.Expression.(*ast.CallExpression)

		if !ok {
			t.Fatalf("exp not *ast.CallExpression. got=%T", stmt.Expression)
		}

		dotCall, ok := callExp.Function.(*ast.Dot)

		if !ok {
			t.Fatalf("callExp.Function is not *ast.Dot. got=%T", callExp.Function)
		}

		testLiteralExpression(t, dotCall.Object, tt.expectedIdentifier)
		testIdentifier(t, dotCall.Right, "type")

		for i, arg := range tt.expectedArgs {
			if callExp.Arguments[i].String() != arg {
				t.Errorf("argument %d wrong. want=%q, got=%q", i, arg, callExp.Arguments[i].String())
			}
		}
	}
}
