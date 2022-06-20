package evaluator

import (
	"ninja/object"
	"testing"
)

func TestFunctionLiteralObject(t *testing.T) {
	input := "function(x) { x + 2; };"

	evaluated := testEval(input, t)
	fn, ok := evaluated.(*object.FunctionLiteral)
	if !ok {
		t.Fatalf("object is not FunctionLiteral. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionObject(t *testing.T) {
	input := "function(x) { x + 2; };"

	evaluated := testEval(input, t)
	fn, ok := evaluated.(*object.FunctionLiteral)
	if !ok {
		t.Fatalf("object is not FunctionLiteral. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
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
			"function add(a, b) { function test(a, b) { return a + b }; return test(a + b, 10); } add(10, 20);",
			40,
		},
		{
			"function add(a, b) { return function test(x, y) { return a + b + x + y }; } add(10, 10)(10, 10);",
			40,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.expression, t)

		if !testObjectLiteral(t, evaluated, tt.rs) {
			t.Errorf("TestCallFunction unable to test")
		}
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

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input, t), tt.expected)
	}
}
