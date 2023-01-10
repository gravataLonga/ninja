package interpreter

import (
	"fmt"
	"github.com/gravataLonga/ninja/ast"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/parser"
	"math"
	"os"
	"strings"
	"testing"
)

func TestLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`1`,
			1,
		},
		{
			`2`,
			2,
		},
		{
			`0.2`,
			0.2,
		},
		{
			`1.2`,
			1.2,
		},
		{
			`1e3`,
			1000.0,
		},
		{
			`0xF`,
			15,
		},
		{
			`"hello ninja warrior"`,
			"hello ninja warrior",
		},
		{
			`false`,
			false,
		},
		{
			`true`,
			true,
		},
		{
			`false`,
			false,
		},
		{
			`[0, 1, 2]`,
			[]interface{}{0, 1, 2},
		},
		{
			`[1, "hello", false, 2.3]`,
			[]interface{}{1, "hello", false, 2.3},
		},
		{
			`{"a":1,"b":2,"c":3}`,
			map[interface{}]interface{}{"a": 1, "b": 2, "c": 3},
		},
		{
			`[1 + 1, 1.0 <= 2.0 , !false]`,
			[]interface{}{2, true, true},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestLiteral[%d]", i), func(t *testing.T) {

			v := interpreter(t, tt.input)

			if v == nil {
				t.Fatalf("Interpreter return nil as result")
			}

			if _, ok := v.(*object.Error); ok {
				t.Fatalf("Interpreter return error. %s", v.Inspect())
			}

			if !testLiteralObject(t, v, tt.expected) {
				t.Fatalf("testLiteralObject got false, expected true.")
			}
		})
	}
}

func TestIfExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`if (true) { 1 } else { 0 }`,
			1,
		},
		{
			`if (false) { 1 } else { 0 }`,
			0,
		},
		{
			`if (false) { 1 } elseif (true) { 2 } else { 3 }`,
			2,
		},
		{
			`if (false) { 1 } elseif (false) { 2 } else { 3 }`,
			3,
		},
		{
			`if (false) { 1 }`,
			nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestIfExpression[%d]", i), func(t *testing.T) {

			v := interpreter(t, tt.input)

			if v == nil {
				t.Fatalf("Interpreter return nil as result")
			}

			if _, ok := v.(*object.Error); ok {
				t.Fatalf("Interpreter return error. %s", v.Inspect())
			}

			if !testLiteralObject(t, v, tt.expected) {
				t.Fatalf("testLiteralObject got false, expected true.")
			}
		})
	}
}

func TestIndexExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`[0, 2, 4][0]`,
			0,
		},
		{
			`[0, 2, 4][-1]`,
			nil,
		},
		{
			`[0, 2, 4][3]`,
			nil,
		},
		{
			`[0, 2, 4][2]`,
			4,
		},
		{
			`[0, 2, 4][1+1]`,
			4,
		},

		{
			`{"a":1,"b":2,"c":3}["a"]`,
			1,
		},
		{
			`{"a":1,"b":2,"c":3}["d"]`,
			nil,
		},

		{
			`"hello ninja"[0]`,
			"h",
		},
		{
			`"hello ninja"[-1]`,
			nil,
		},
		{
			`"hello ninja"[50]`,
			nil,
		},
		{
			`"hello ninja"[10]`,
			"a",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestIndexExpression[%d]", i), func(t *testing.T) {

			v := interpreter(t, tt.input)

			if v == nil {
				t.Fatalf("Interpreter return nil as result")
			}

			if _, ok := v.(*object.Error); ok {
				t.Fatalf("Interpreter return error. %s", v.Inspect())
			}

			if !testLiteralObject(t, v, tt.expected) {
				t.Fatalf("testLiteralObject got false, expected true.")
			}
		})
	}
}

func TestTernaryOperatorExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`true ? 1 : 0`,
			1,
		},
		{
			`false ? 1 : 0`,
			0,
		},
		{
			`1 ? 1 : 0`,
			1,
		},
		{
			`0 ? 1 : 0`,
			1,
		},
		{
			`!1 ? 1 : 0`,
			0,
		},
		{
			`2 == 2 ? 1 : 0`,
			1,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestTernaryOperatorExpression[%d]", i), func(t *testing.T) {

			v := interpreter(t, tt.input)

			if v == nil {
				t.Fatalf("Interpreter return nil as result")
			}

			if _, ok := v.(*object.Error); ok {
				t.Fatalf("Interpreter return error. %s", v.Inspect())
			}

			if !testLiteralObject(t, v, tt.expected) {
				t.Fatalf("testLiteralObject got false, expected true.")
			}
		})
	}
}

func TestElvisOperatorExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`true ?: false`,
			true,
		},
		{
			`"hello" ?: "world"`,
			"hello",
		},
		{
			`!true ?: "world"`,
			"world",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestElvisOperatorExpression[%d]", i), func(t *testing.T) {

			v := interpreter(t, tt.input)

			if v == nil {
				t.Fatalf("Interpreter return nil as result")
			}

			if _, ok := v.(*object.Error); ok {
				t.Fatalf("Interpreter return error. %s", v.Inspect())
			}

			if !testLiteralObject(t, v, tt.expected) {
				t.Fatalf("testLiteralObject got false, expected true.")
			}
		})
	}
}

func createParser(t *testing.T, input string) ast.Node {
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		t.Fatalf("Parsing program got some errors: %v", p.Errors()[0])
	}

	return program
}

func interpreter(t *testing.T, input string) object.Object {
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		t.Fatalf("Parsing program got some errors: %v", p.Errors()[0])
	}

	i := New(os.Stdout)
	return i.Interpreter(program)
}

func testLiteralObject(t *testing.T, result object.Object, expected interface{}) bool {
	switch expected := expected.(type) {
	case int:
		if result == nil {
			t.Errorf("Expected %s, got: nil", object.INTEGER_OBJ)
			return false
		}

		v, ok := result.(*object.Integer)
		if !ok {
			t.Errorf("Expected %s. Got: %s", object.INTEGER_OBJ, result.Type())
			return false
		}

		if v.Value != int64(expected) {
			t.Errorf("Expected %d. Got: %d", expected, v.Value)
			return false
		}

		return true
	case nil:
		if _, ok := result.(*object.Null); ok {
			return true
		}
		t.Errorf("Expected NULL. Got: nil")
		return false
	case float64:
		if result == nil {
			t.Errorf("Expected %s, got: nil", object.FLOAT_OBJ)
			return false
		}

		v, ok := result.(*object.Float)
		if !ok {
			t.Errorf("Expected %s. Got: %s", object.FLOAT_OBJ, result.Type())
			return false
		}

		if v.Value != float64(expected) {
			t.Errorf("Expected %.f. Got: %.f", expected, v.Value)
			return false
		}

		return true
	case string:
		if result == nil {
			t.Errorf("Expected %s, got: nil", object.STRING_OBJ)
			return false
		}

		v, ok := result.(*object.String)
		if !ok {
			t.Errorf("Expected %s. Got: %s", object.STRING_OBJ, result.Type())
			return false
		}

		if v.Value != expected {
			t.Errorf("Expected %s. Got: %s", expected, v.Value)
			return false
		}
		return true
	case bool:
		if result == nil {
			t.Errorf("Expected %s, got: nil", object.BOOLEAN_OBJ)
			return false
		}

		v, ok := result.(*object.Boolean)
		if !ok {
			t.Errorf("Expected %s. Got: %s", object.BOOLEAN_OBJ, result.Type())
			return false
		}

		if v.Value != expected {
			t.Errorf("Expected %v. Got: %v", expected, v.Value)
			return false
		}
		return true
	case []interface{}:
		if result == nil {
			t.Errorf("Expected %s, got: nil", object.ARRAY_OBJ)
			return false
		}

		v, ok := result.(*object.Array)
		if !ok {
			t.Errorf("Expected %s. Got: %s", object.ARRAY_OBJ, result.Type())
			return false
		}

		if len(v.Elements) != len(expected) {
			t.Errorf("Didn't get same number of element on array. Expected %d. got: %d", len(expected), len(v.Elements))
		}

		for i, vE := range v.Elements {
			if !testLiteralObject(t, vE, expected[i]) {
				t.Errorf("ArrayElement %d, didn't match expected %v. Got: %v", i, expected[i], vE)
				return false
			}
		}
		return true
	case map[interface{}]interface{}:
		if result == nil {
			t.Errorf("Expected %s, got: nil", object.HASH_OBJ)
			return false
		}

		// @todo need more checks
		_, ok := result.(*object.Hash)
		if !ok {
			t.Errorf("Expected %s. Got: %s", object.HASH_OBJ, result.Type())
			return false
		}

		return true
	}
	return false
}

// testBooleanObject helper for testing object.Object is equal expected.
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

// testIntegerObject helper for testing object.Object is equal expected.
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

// testStringObject helper for testing object.Object is equal expected.
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

// testFloatObject helper for testing object.Object is equal expected.
func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return false
	}

	max := math.Max(result.Value, expected)
	min := math.Min(result.Value, expected)

	if max-min >= object.EPSILON {
		t.Errorf("object has wrong value. got=%.30f, want=%.30f", result.Value, expected)
		return false
	}

	return true
}

// testNullObject helper for testing object.Object is equal expected NULL.
func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != object.NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

// testObjectLiteral helper for testing if object is equal expected interface{}
// we will decide which object test based on value passed in interface{}
func testObjectLiteral(
	t *testing.T,
	objectResult object.Object,
	expected interface{},
) bool {

	switch expected.(type) {
	case object.Hash:
		hash, ok := objectResult.(*object.Hash)
		if !ok {
			t.Errorf("type of exp expected to be object.Hash. Got: . got=%T", objectResult)
		}

		hashExpected, _ := expected.(object.Hash)

		if len(hashExpected.Pairs) != len(hash.Pairs) {
			t.Fatalf("object.Hash pairs elements expected %d. got=%d", len(hashExpected.Pairs), len(hash.Pairs))
		}

		for k, hashPair := range hash.Pairs {
			if !testObjectLiteral(t, hashPair.Key, hashExpected.Pairs[k].Key) {
				return false
			}

			if !testObjectLiteral(t, hashPair.Value, hashExpected.Pairs[k].Value) {
				return false
			}
		}
		return true

	case object.Array:

		arr, ok := objectResult.(*object.Array)
		if !ok {
			t.Fatalf("type of exp expected to be object.Array. Got: . got=%s", objectResult.Inspect())
		}

		arrExpected, _ := expected.(object.Array)

		if len(arrExpected.Elements) != len(arr.Elements) {
			t.Fatalf("object.Array elements expected %d. got=%d", len(arrExpected.Elements), len(arr.Elements))
		}

		for index, item := range arr.Elements {
			if !testObjectLiteral(t, item, arrExpected.Elements[index]) {
				return false
			}
		}
		return true
	case *object.Integer:
		expected := expected.(*object.Integer)
		return testIntegerObject(t, objectResult, expected.Value)
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
	case *object.String:
		expected := expected.(*object.String)
		return testStringObject(t, objectResult, expected.Value)
	case string:
		expected := expected.(string)
		return testStringObject(t, objectResult, expected)
	case nil:
		return testNullObject(t, objectResult)
	}

	t.Errorf("type of exp not handled. got=%T", expected)
	return false
}

// checkParserErrors check if there are parser errors
func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
