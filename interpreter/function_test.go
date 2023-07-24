package interpreter

import (
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/object"
	"os"
	"testing"
)

func TestFunctionLiteral(t *testing.T) {
	p := createParser(t, `function add() { 1 }`)
	i := New(os.Stdout, object.NewEnvironment())
	i.Interpreter(p)

	_, ok := i.env.Get("add")
	if !ok {
		t.Fatalf("Expected add identifier on env")
	}
}

func TestFunctionLiteralObject(t *testing.T) {
	input := "function(x) { x + 2; };"

	evaluated := interpreter(t, input)
	fn, ok := evaluated.(*object.FunctionLiteral)
	if !ok {
		t.Fatalf("object is not FunctionLiteral. got=%T (%+v)", evaluated, evaluated)
	}

	parameters, ok := fn.Parameters.([]ast.Expression)
	if !ok {
		t.Fatalf("object.parameters is not []ast.Expression. got=%T (%+v)", fn.Parameters, fn.Parameters)
	}

	body, ok := fn.Body.(*ast.BlockStatement)
	if !ok {
		t.Fatalf("object.body is not ast.BlockStatement. got=%T (%+v)", fn.Body, fn.Body)
	}

	if len(parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", parameters)
	}

	if parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", parameters[0])
	}

	expectedBody := "(x + 2)"

	if body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, body.String())
	}
}

func TestFunctionObject(t *testing.T) {
	input := "function(x) { x + 2; };"

	evaluated := interpreter(t, input)
	fn, ok := evaluated.(*object.FunctionLiteral)
	if !ok {
		t.Fatalf("object is not FunctionLiteral. got=%T (%+v)", evaluated, evaluated)
	}

	parameters, ok := fn.Parameters.([]ast.Expression)
	if !ok {
		t.Fatalf("object.parameters is not []ast.Expression. got=%T (%+v)", fn.Parameters, fn.Parameters)
	}

	body, ok := fn.Body.(*ast.BlockStatement)
	if !ok {
		t.Fatalf("object.body is not ast.BlockStatement. got=%T (%+v)", fn.Body, fn.Body)
	}

	if len(parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}

	if parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", parameters[0])
	}

	expectedBody := "(x + 2)"

	if body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, body.String())
	}
}

func TestFunctionWithDefaultArguments(t *testing.T) {
	tests := []struct {
		input  string
		result interface{}
	}{
		{
			`function (a = 0) { return a;}();`,
			0,
		},
		{
			`function (a = 0) { return a;}(2);`,
			2,
		},
		{
			`function add (a = 0) { return a;}; add();`,
			0,
		},
		{
			`function add (a = 0) { return a;}; add(2);`,
			2,
		},
		{
			`function (a, b = 1) { return a + b;}(1);`,
			2,
		},
		{
			`function (a, b = 1) { return a + b;}(1, 2);`,
			3,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestFunctionWithDefaultArguments[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)

			if !testLiteralObject(t, evaluated, tt.result) {
				t.Errorf("TestCallFunction unable to test")
			}
		})
	}
}

func TestCallFunction(t *testing.T) {
	tests := []struct {
		expression string
		rs         interface{}
	}{

		{
			"function(x) { x + 2; }(10);",
			12,
		},
		{
			"function add (x) { x + 2; }; add(2)",
			4,
		},
		{
			"var add = function (x) { return x + 2; }; add(2);",
			4,
		},
		{
			"var say = function (x) { return \"Hello \" + x; }; say(\"Dog\");",
			"Hello Dog",
		},
		{
			"function t(a) { return a + 20.0; } t(10.5);",
			30.5,
		},
		{
			"function t(a) { return !a; } t(1);",
			false,
		},
		{
			"function t(a) { if (a > 0) { return true; } else { return false }; } t(1);",
			true,
		},
		{
			"function t(a) { return a > 0; } t(1);",
			true,
		},
		{
			"function t(a) { return !(a > 0); } t(1);",
			false,
		},
		{
			"function add(a, b) { return a + b; } add(5, add(5, 5));",
			15,
		},
		{
			"function add(a, b) { function test(e, f) { return e + f }; return test(a + b, 10); } add(10, 20);",
			40,
		},
		/*{
			"function add(a, b) { return function test(x, y) { return a + b + x + y }; } add(10, 10)(10, 10);",
			40,
		},
		{
			"var a = 0; function add() { return function increment() { a++; return a; }}; var b = add()(); add()();",
			2,
		},*/
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestCallFunction[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.expression)

			if !testLiteralObject(t, evaluated, tt.rs) {
				t.Errorf("TestCallFunction unable to test")
			}
		})

	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"var identity = function(x) { x; }; identity(5);", 5},
		{"function identity(x) { x; }; identity(5);", 5},
		{"var identity = function(x) { return x; }; identity(5);", 5},
		{"function identity(x) { return x; }; identity(5);", 5},
		{"var double = function(x) { x * 2; }; double(5);", 10},
		{"var add = function(x, y) { x + y; }; add(5, 5);", 10},
		{"var add = function(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"function(x) { x; }(5)", 5},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test[%d]", i), func(t *testing.T) {
			testIntegerObject(t, interpreter(t, tt.input), tt.expected)
		})
	}
}

func TestCallWrongParameters(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{"function (x) {}();", "Function expected 1 arguments, got 0 at ( at [Line: 1, Offset: 16]"},
		{"function () {}(0);", "Function expected 0 arguments, got 1 at ( at [Line: 1, Offset: 15]"},
		{"function () { return add(); }();", "identifier not found: add IDENT at [Line: 1, Offset: 25]"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestCallWrongParameters[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedErrorMessage {
				t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedErrorMessage, errObj.Message)
			}
		})

	}
}
