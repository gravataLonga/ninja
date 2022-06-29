package evaluator

import (
	"ninja/lexer"
	"ninja/object"
	"ninja/parser"
	"strings"
	"testing"
)

// testBooleanObject helper for testing if object.Object is equal expected.
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

// testIntegerObject helper for testing if object.Object is equal expected.
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

// testStringObject helper for testing if object.Object is equal expected.
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

// testFloatObject helper for testing if object.Object is equal expected.
func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value-expected >= object.EPSILON {
		t.Errorf("object has wrong value. got=%.30f, want=%.30f", result.Value, expected)
		return false
	}

	return true
}

// testNullObject helper for testing if object.Object is equal expected NULL.
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

		for _, hashPair := range hash.Pairs {
			if !testObjectLiteral(t, hashPair.Value, expected) {
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

// testEval execute input code and check if there are parser error
// and return result object.Object
func testEval(input string, t *testing.T) object.Object {
	l := lexer.New(strings.NewReader(input))
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	checkParserErrors(t, p)

	return Eval(program, env)
}
