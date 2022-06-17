package evaluator

import (
	"ninja/lexer"
	"ninja/object"
	"ninja/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{

		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"++4", 5},
		{"--6", 5},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5.2", 5.2},
		{"10.0", 10.0},
		{"10.000000000123", 10.0000000001},
		{"-5.2", -5.2},
		{"-10.0", -10.0},
		{"++1.0", 2.0},
		{"--1.0", 0.0},
		{"5.0 + 5.0 + 5.5 + 5.5 - 10", 11},
		{"2.2 * 2.2 * 2.2 * 2.2 * 2.2", 51.5363200000},
		{"-50.50 + 100.50 + -50.50", -0.5},
		{"5.5 * 2.5 + 10.5", 24.25},
		{"5.5 + 2.5 * 10.5", 31.75},
		{"20 + 2.0 * -10", 0.0},
		{"50.10 / 2.20 * 2.20 + 10.2", 60.2999999999},
		{"2 * (5.2 + 10.2)", 30.8},
		{"3 * 3 * 3 + 10.5", 37.5},
		{"3 * (3 * 3.5) + 10", 41.5},
		{"(5 + 10 * 2 + 15 / 3) * 2.2 + -10", 56.0},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"true && true", true},
		{"true && false", false},
		{"true || false", true},
		{"false || true", true},
		{"false || false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"1 && 1", true},
		{"1 && 0", true},
		{"0 && 1", true},
		{"1 || false", true},
		{"false || false", false},
		{"false && 1", false},
		{"10.0 == 10.0", true},
		{"10.5 >= 10.0", true},
		{"10.5 && 1", true},
		{"(20.5 > 5) == true", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"(1 <= 1) == true", true},
		{"(1 >= 1) == true", true},
		{"(1 >= 1) && true", true},
		{"(1 || false) || false", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"!10.0", false},
		{"!!10.0", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 <= 2) { 10 }", 10},
		{"if (1 >= 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 >= 2) { 10 } else { 20 }", 20},
		{"if (1 <= 2) { 10 } else { 20 }", 10},
		{"if (1.0) { 10 }", 10},
		{"if (1.0 < 2.0) { 10 }", 10},
		{"if (1.0 > 2.0) { 10 }", nil},
		{"if (1.0 <= 2.0) { 10 }", 10},
		{"if (1.0 >= 2.0) { 10 }", nil},
		{"if (1.0 > 2.0) { 10 } else { 20 }", 20},
		{"if (1.0 < 2.0) { 10 } else { 20 }", 10},
		{"if (1.0 >= 2.0) { 10 } else { 20 }", 20},
		{"if (1.0 <= 2.0) { 10 } else { 20 }", 10},
		{"if (1) { 20.50 }", 20.50},
		{"if (1.0 < 2.0) { 50.20 }", 50.20},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		float, okFloat := tt.expected.(float64)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else if okFloat {
			testFloatObject(t, evaluated, float64(float))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"return 20;", 20},
		{"return 20.50;", 20.50},
		{"return 30; 9;", 30},
		{"return 8 * 5; 9;", 40},
		{"9; return 12 * 5; 9;", 60},
		{"if (10 > 1) { return 10; }", 10},
		{
			`
if (10 > 1) {
  if (10 > 1) {
    return 90;
  }

  return 1;
}
`,
			90,
		},
		{
			`
		var f = function(x) {
		  return x;
		  x + 10;
		};
		f(10);`,
			10,
		},
		{
			`
		var f = function(x) {
		   var result = x + 10;
		   return result;
		   return 10;
		};
		f(10);`,
			20,
		},
		{
			`
		function f(x) {
		  return x;
		  x + 10;
		};
		f(10);`,
			10,
		},
		{
			`
		function f(x) {
		   var result = x + 10;
		   return result;
		   return 10;
		};
		f(10);`,
			20,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		integer, ok := tt.expected.(int)
		float, okFloat := tt.expected.(float64)

		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else if okFloat {
			testFloatObject(t, evaluated, float64(float))
		}

	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"50.50 + true;",
			"type mismatch: FLOAT + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5.0 + true; 5.3;",
			"type mismatch: FLOAT + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"true + false + true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
if (10 > 1) {
  if (10 > 1) {
    return true + false;
  }

  return 1;
}
`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			`"Hello" - "World"`,
			"unknown operator: STRING - STRING",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

// @todo reassign values we need to check for same type values?
func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"var a = 5; a;", 5},
		{"var a = 5 * 5; a;", 25},
		{"var a = 5; var b = a; b;", 5},
		{"var a = 5; var b = a; var c = a + b + 5; c;", 15},
		{"var a = 5; ++a;", 6},
		{"var a = 5; --a;", 4},
		{"var a = 5; a = a + 5; a;", 10},
		{"var a = 5; var b = 10; a = a + b * a; a", 55},
		// {"var a = 5; a++;", 5},
		// {"var a = 5; a--;", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionLiteralObject(t *testing.T) {
	input := "function(x) { x + 2; };"

	evaluated := testEval(input)
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

	evaluated := testEval(input)
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
		evaluated := testEval(tt.expression)

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
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
		{`len([1, 2, 3])`, 3},
		{`len([])`, 0},
		{`puts("hello", "world!")`, nil},
		{`first([1, 2, 3])`, 1},
		{`first([])`, nil},
		{`first(1)`, "argument to `first` must be ARRAY, got INTEGER"},
		{`last([1, 2, 3])`, 3},
		{`last([])`, nil},
		{`last(1)`, "argument to `last` must be ARRAY, got INTEGER"},
		{`rest([1, 2, 3])`, []int{2, 3}},
		{`rest([])`, nil},
		{`push([], 1)`, []int{1}},
		{`push(1, 1)`, "argument to `push` must be ARRAY, got INTEGER"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case nil:
			testNullObject(t, evaluated)
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)",
					evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q",
					expected, errObj.Message)
			}
		case []int:
			array, ok := evaluated.(*object.Array)
			if !ok {
				t.Errorf("obj not Array. got=%T (%+v)", evaluated, evaluated)
				continue
			}

			if len(array.Elements) != len(expected) {
				t.Errorf("wrong num of elements. want=%d, got=%d",
					len(expected), len(array.Elements))
				continue
			}

			for i, expectedElem := range expected {
				testIntegerObject(t, array.Elements[i], int64(expectedElem))
			}

		}

	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d",
			len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"var i = 0; [1][i];",
			1,
		},
		{
			"[1, 2, 3][1 + 1];",
			3,
		},
		{
			"var myArray = [1, 2, 3]; myArray[2];",
			3,
		},
		{
			"var myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];",
			6,
		},
		{
			"var myArray = [1, 2, 3]; var i = myArray[0]; myArray[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestHashLiterals(t *testing.T) {
	input := `var two = "two";
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 6 / 2,
		4: 4,
		true: 5,
		false: 6
	}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		object.TRUE.HashKey():                      5,
		object.FALSE.HashKey():                     6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`var key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{true: 5}[true]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
		{
			`{1.0: 35}[1.0]`,
			35,
		},
		{
			`{1.000000: 35}[1.0]`,
			35,
		},
		{
			`{1.000001: 35}[1.0]`,
			nil,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestForStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`for(var i = 0; i <= 1; i = i + 1) { i; }`,
			2,
		},
		{
			`for(var i = 0; i > 100; i = i + 1) { i; }`,
			nil,
		},
		{
			`var i = 0; for(; i <= 1; i = i + 1) { i; }`,
			2,
		},
		{
			`for(var i = 0; i <= 1; ) { var i = i + 1; i; }`,
			2,
		},
		{
			`var i = 0; var a = 0; for(;i <= 10 && a <= 10;) { var i = i + 2; var a = a + 2; i; }`,
			12,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else if evaluated != nil {
			t.Errorf("result isnt nil. Got %v", evaluated)
		}
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
		return false
	}
	return true
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%s, want=%s", result.Value, expected)
		return false
	}

	return true
}

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%.30f, want=%.30f", result.Value, expected)
		return false
	}

	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != object.NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func testObjectLiteral(
	t *testing.T,
	objectResult object.Object,
	expected interface{},
) bool {

	switch expected.(type) {
	case int:
		expected := expected.(int)
		return testIntegerObject(t, objectResult, int64(expected))
	case int64:
		expected := expected.(int64)
		return testIntegerObject(t, objectResult, expected)
	case bool:
		expected := expected.(bool)
		return testBooleanObject(t, objectResult, expected)
	case float64:
		expected := expected.(float64)
		return testFloatObject(t, objectResult, expected)
	case string:
		expected := expected.(string)
		return testStringObject(t, objectResult, expected)
	}

	t.Errorf("type of exp not handled. got=%T", expected)
	return false
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}
