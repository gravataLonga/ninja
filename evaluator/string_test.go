package evaluator

import (
	"fmt"
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
		{`"Hello"[0] != "World"[0]`, true},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestEvalStringExpression[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)
			testObjectLiteral(t, evaluated, tt.expected)
		})
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

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestErrorStringHandling[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedMessage {
				t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
			}
		})
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

	if errObj.Message != "TypeError: string.type() takes exactly 0 argument (1 given)" {
		t.Errorf("error message expected to be: \"%s\". got: \"%s\"", "TypeError: string.type() takes exactly 0 argument (1 given)", errObj.Message)
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
			`("a;b,c;d" + "other,nice").split(",")`,
			`[a;b, c;dother, nice]`,
		},
		{
			`"a\nb\nc".split("\n")`,
			`[a, b, c]`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodSplit[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			if evaluated.Inspect() != tt.expected {
				t.Errorf("string.split() expected %s. Got: %s", tt.expected, evaluated.Inspect())
			}
		})
	}

}

func TestStringMethodSplitWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".split()`,
			"TypeError: string.split() takes exactly 1 argument (0 given)",
		},
		{
			`"ola".split("hello", "abc")`,
			"TypeError: string.split() takes exactly 1 argument (2 given)",
		},
		{
			`"ola".split(true)`,
			"TypeError: string.split() expected argument #1 to be `STRING` got `BOOLEAN`",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodSplitWrongParameter[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedErrorMessage {
				t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
			}
		})

	}
}

func TestStringMethodLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{
			`"".length()`,
			0,
		},
		{
			`"hello".length()`,
			5,
		},
		{
			`"Acção".length()`,
			5,
		},
		{
			`"नमस्ते दुनिया".length()`,
			13,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodLength_%d", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			if !testIntegerObject(t, evaluated, tt.expected) {
				t.Errorf("string.length() expected %d. Got: %s", tt.expected, evaluated.Inspect())
			}
		})
	}
}

func TestStringMethodLengthWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".length(1)`,
			"TypeError: string.length() takes exactly 0 argument (1 given)",
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
			"TypeError: string.contain() takes exactly 1 argument (0 given)",
		},
		{
			`"ola".contain("hello", "abc")`,
			"TypeError: string.contain() takes exactly 1 argument (2 given)",
		},
		{
			`"ola".contain(true)`,
			"TypeError: string.contain() expected argument #1 to be `STRING` got `BOOLEAN`",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodContainWrongParameter[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedErrorMessage {
				t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
			}
		})
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

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodIndex[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			testIntegerObject(t, evaluated, tt.expected)
		})
	}
}

func TestStringMethodIndexWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".index()`,
			"TypeError: string.index() takes exactly 1 argument (0 given)",
		},
		{
			`"ola".index("hello", "abc")`,
			"TypeError: string.index() takes exactly 1 argument (2 given)",
		},
		{
			`"ola".index(true)`,
			"TypeError: string.index() expected argument #1 to be `STRING` got `BOOLEAN`",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodIndexWrongParameter[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedErrorMessage {
				t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
			}
		})
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

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodUpper[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			testStringObject(t, evaluated, tt.expected)
		})
	}
}

func TestStringMethodUpperWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".upper("hello", "abc")`,
			"TypeError: string.upper() takes exactly 0 argument (2 given)",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodUpperWrongParameter[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedErrorMessage {
				t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
			}
		})
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

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodLower[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			testStringObject(t, evaluated, tt.expected)
		})
	}
}

func TestStringMethodLowerWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".lower("hello", "abc")`,
			"TypeError: string.lower() takes exactly 0 argument (2 given)",
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

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodTrim[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			testStringObject(t, evaluated, tt.expected)
		})
	}
}

func TestStringMethodTrimWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"ola".trim("hello")`,
			"TypeError: string.trim() takes exactly 0 argument (1 given)",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodTrimWrongParameter[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedErrorMessage {
				t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
			}
		})

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

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodInteger[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			testIntegerObject(t, evaluated, tt.expected)
		})

	}
}

func TestStringMethodIntegerWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"100".int("hello")`,
			"TypeError: string.int() takes exactly 0 argument (1 given)",
		},
		{
			`"0x000".int()`,
			"TypeError: string.int() fail to convert to int. Got: strconv.ParseInt: parsing \"0x000\": invalid syntax",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodIntegerWrongParameter[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedErrorMessage {
				t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
			}
		})

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

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodFloat[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			testFloatObject(t, evaluated, tt.expected)
		})
	}
}

func TestStringMethodFloatWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"100.0".float("hello")`,
			"TypeError: string.float() takes exactly 0 argument (1 given)",
		},
		{
			`"0x000".float()`,
			"TypeError: string.float() fail to convert to float. Got: strconv.ParseFloat: parsing \"0x000\": invalid syntax",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodFloatWrongParameter[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedErrorMessage {
				t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
			}
		})

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

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodReplace[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			testStringObject(t, evaluated, tt.expected)
		})
	}
}

func TestStringMethodReplaceWrongParameter(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`"Hello World".replace()`,
			"TypeError: string.replace() takes exactly 2 argument (0 given)",
		},
		{
			`"Hello World".replace("this")`,
			"TypeError: string.replace() takes exactly 2 argument (1 given)",
		},
		{
			`"Hello World".replace("this", "that", "other")`,
			"TypeError: string.replace() takes exactly 2 argument (3 given)",
		},
		{
			`"Hello World".replace(1, "other")`,
			"TypeError: string.replace() expected argument #1 to be `STRING` got `INTEGER`",
		},
		{
			`"Hello World".replace("other", [])`,
			"TypeError: string.replace() expected argument #2 to be `STRING` got `ARRAY`",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStringMethodReplaceWrongParameter[%d]", i), func(t *testing.T) {
			evaluated := testEval(tt.input, t)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedErrorMessage {
				t.Errorf("error message expected to be: \"%s\". got: \"%s\"", tt.expectedErrorMessage, errObj.Message)
			}
		})

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
