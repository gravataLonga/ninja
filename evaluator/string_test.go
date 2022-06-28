package evaluator

import (
	"ninja/object"
	"testing"
)

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`

	evaluated := testEval(input, t)
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

	evaluated := testEval(input, t)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

// @todo we can improve performance of comparisons.
func TestEvalStringExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{

		{`"Hello"`, "Hello"},
		{`"Hello" + "World"`, "HelloWorld"},
		{`"Hello" && "World"`, true},
		{`"Hello" && false`, false},
		{`"Hello" || false`, true},
		{`!"Hello"`, false},
		{`"Hello" == "Hello"`, true},
		{`"Hello" != "Hello"`, false},
		{`"Hello" == "World"`, false},
		{`"Hello" != "World"`, true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)
		testObjectLiteral(t, evaluated, tt.expected)
	}
}

func TestErrorStringHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"-\"hello\"",
			"unknown operator: -STRING",
		},
		{
			`"Hello" - "Nice"`,
			"unknown operator: STRING - STRING",
		},
		{
			`"Hello" * "Nice"`,
			"unknown operator: STRING * STRING",
		},
		{
			`"Hello" / "Nice"`,
			"unknown operator: STRING / STRING",
		},
		{
			`++"Nice"`,
			"unknown object type STRING for operator ++",
		},
		{
			`--"Nice"`,
			"unknown object type STRING for operator --",
		},
		{
			`"1" < "2"`,
			"unknown operator: STRING < STRING",
		},
		{
			`"1" > "2"`,
			"unknown operator: STRING > STRING",
		},
		{
			`"1" <= "2"`,
			"unknown operator: STRING <= STRING",
		},
		{
			`"1" >= "2"`,
			"unknown operator: STRING >= STRING",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestStringIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`"ola"[0]`,
			"o",
		},
		{
			`"ola"[2]`,
			"a",
		},
		{
			`"ola"[3]`,
			nil,
		},
		{
			`"ola"[-1]`,
			nil,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)
		str, ok := tt.expected.(string)
		if ok {
			testStringObject(t, evaluated, string(str))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestStringMethodType(t *testing.T) {
	evaluated := testEval(`"ola".type()`, t)
	testObjectLiteral(t, evaluated, object.STRING_OBJ)
}

func TestStringMethodTypeWrongParameter(t *testing.T) {
	evaluated := testEval(`"ola".type(1)`, t)

	errObj, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
	}

	if errObj.Message != "method type not accept any arguments. got: [1]" {
		t.Errorf("error message expected to be: \"%s\". got: \"%s\"", "method type not accept any arguments. Got: [1]", errObj.Message)
	}
}

func TestStringMethodSplit(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`"1,2,3".split(",")`,
			`[1, 2, 3]`,
		},
		{
			`"a;b;c".split(",")`,
			`[a;b;c]`,
		},
		{
			`"".split(",")`,
			`[]`,
		},
		{
			`"a;b,c;d".split(",")`,
			`[a;b, c;d]`,
		},
		{
			`"a;b,c;d" + "other,nice".split(",")`,
			`[a;b, c;dother, nice]`,
		},
		{
			`"a\nb\nc".split("\n")`,
			`[a, b, c]`,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		if evaluated.Inspect() != tt.expected {
			t.Errorf("string.split() expected %s. Got: %s", tt.expected, evaluated.Inspect())
		}
	}

}

func TestStringMethodSplitWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".split()`,
			"split method expect exactly 1 argument",
		},
		{
			`"ola".split("hello", "abc")`,
			"split method expect exactly 1 argument",
		},
		{
			`"ola".split(true)`,
			"first argument must be string, got: BOOLEAN",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
		}

		if errObj.Message != tt.expectedErrorMessage {
			t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
		}
	}
}

func TestStringMethodContain(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			`"1".contain("1")`,
			true,
		},
		{
			`"".contain("1")`,
			false,
		},
		{
			`"".contain("")`,
			true,
		},
		{
			`"1, 2 Hello World".contain("Hello World")`,
			true,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestStringMethodContainWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".contain()`,
			"contain method expect exactly 1 argument",
		},
		{
			`"ola".contain("hello", "abc")`,
			"contain method expect exactly 1 argument",
		},
		{
			`"ola".contain(true)`,
			"first argument must be string, got: BOOLEAN",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
		}

		if errObj.Message != tt.expectedErrorMessage {
			t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
		}
	}
}

func TestStringMethodIndex(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{
			`"1".index("1")`,
			0,
		},
		{
			`"1".index("0")`,
			-1,
		},
		{
			`"Hello World".index("W")`,
			6,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestStringMethodIndexWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".index()`,
			"index method expect exactly 1 argument",
		},
		{
			`"ola".index("hello", "abc")`,
			"index method expect exactly 1 argument",
		},
		{
			`"ola".index(true)`,
			"first argument must be string, got: BOOLEAN",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
		}

		if errObj.Message != tt.expectedErrorMessage {
			t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
		}
	}
}

func TestStringMethodUpper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`"hello".upper()`,
			"HELLO",
		},
		{
			`"hEllO".upper()`,
			"HELLO",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		testStringObject(t, evaluated, tt.expected)
	}
}

func TestStringMethodUpperWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".upper("hello", "abc")`,
			"upper method expect exactly 0 argument",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
		}

		if errObj.Message != tt.expectedErrorMessage {
			t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
		}
	}
}

func TestStringMethodLower(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`"HELLO".lower()`,
			"hello",
		},
		{
			`"hEllO".lower()`,
			"hello",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		testStringObject(t, evaluated, tt.expected)
	}
}

func TestStringMethodLowerWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".lower("hello", "abc")`,
			"lower method expect exactly 0 argument",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
		}

		if errObj.Message != tt.expectedErrorMessage {
			t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
		}
	}
}

func TestStringMethodTrim(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`"  hello  ".trim()`,
			"hello",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		testStringObject(t, evaluated, tt.expected)
	}
}

func TestStringMethodTrimWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".trim("hello")`,
			"trim method expect exactly 0 argument",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
		}

		if errObj.Message != tt.expectedErrorMessage {
			t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
		}
	}
}

func TestStringMethodInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{
			`"1".int()`,
			1,
		},
		{
			`"1000".int()`,
			1000,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestStringMethodIntegerWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"100".int("hello")`,
			"int method expect exactly 0 argument",
		},
		{
			`"0x000".int()`,
			"string.int() fail to convert to int. Got: strconv.ParseInt: parsing \"0x000\": invalid syntax",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
		}

		if errObj.Message != tt.expectedErrorMessage {
			t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
		}
	}
}

func TestStringMethodFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{
			`"1.2".float()`,
			1.2,
		},
		{
			`"0.0".float()`,
			0.0,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		testFloatObject(t, evaluated, tt.expected)
	}
}

func TestStringMethodFloatWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"100.0".float("hello")`,
			"float method expect exactly 0 argument",
		},
		{
			`"0x000".float()`,
			"string.float() fail to convert to float. Got: strconv.ParseFloat: parsing \"0x000\": invalid syntax",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
		}

		if errObj.Message != tt.expectedErrorMessage {
			t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
		}
	}
}

func TestStringMethodReplace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`"Hello World".replace("World", "Jonathan")`,
			"Hello Jonathan",
		},
		{
			`"Hello World".replace("Nothing", "Jonathan")`,
			"Hello World",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		testStringObject(t, evaluated, tt.expected)
	}
}

func TestStringMethodReplaceWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"Hello World".replace()`,
			"replace method expect exactly 2 argument",
		},
		{
			`"Hello World".replace("this")`,
			"replace method expect exactly 2 argument",
		},
		{
			`"Hello World".replace("this", "that", "other")`,
			"replace method expect exactly 2 argument",
		},
		{
			`"Hello World".replace(1, "other")`,
			"first argument must be string, got: INTEGER",
		},
		{
			`"Hello World".replace("other", [])`,
			"second argument must be string, got: ARRAY",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input, t)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
		}

		if errObj.Message != tt.expectedErrorMessage {
			t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
		}
	}
}

func TestStringMethodNotFound(t *testing.T) {
	input := `"a".ups()`
	expected := ""

	evaluated := testEval(input, t)

	errObj, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
	}

	if errObj.Message != "method ups not exists on string object." {
		t.Errorf("error message expected to be %s. got: %s", expected, errObj)
	}
}
