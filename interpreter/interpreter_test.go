package interpreter

import (
	"fmt"
	"github.com/gravataLonga/ninja/lexer"
	"github.com/gravataLonga/ninja/object"
	"github.com/gravataLonga/ninja/parser"
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

func TestPrefixOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`-1`,
			-1,
		},
		{
			`-1.0`,
			-1.0,
		},
		{
			`++1`,
			2,
		},
		{
			`--1`,
			0,
		},
		{
			`++1.0`,
			2.0,
		},
		{
			`--1.0`,
			0.0,
		},
		{
			`!true`,
			false,
		},
		{
			`!false`,
			true,
		},
		{
			`!0`,
			false,
		},
		{
			`!1`,
			false,
		},
		{
			`!!0`,
			true,
		},
		{
			`!!1`,
			true,
		},
		{
			`!0.0`,
			false,
		},
		{
			`!1.0`,
			false,
		},
		{
			`!!0.0`,
			true,
		},
		{
			`!!1.0`,
			true,
		},
		{
			`![]`,
			false,
		},
		{
			`![1, 2]`,
			false,
		},
		{
			`!{}`,
			false,
		},
		{
			`!{"a":1}`,
			false,
		},
		{
			`!![]`,
			true,
		},
		{
			`!![1, 2]`,
			true,
		},
		{
			`!!{}`,
			true,
		},
		{
			`!!{"a":1}`,
			true,
		},
		{
			`!"hello ninja warrior"`,
			false,
		},
		{
			`!!"hello ninja warrior"`,
			true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestPrefixOperator[%d]", i), func(t *testing.T) {

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

func TestPostfixOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`1++`,
			2,
		},
		{
			`0++`,
			1,
		},
		{
			`0.0++`,
			1.0,
		},
		{
			`1.0++`,
			2.0,
		},

		{
			`1--`,
			0,
		},
		{
			`0--`,
			-1,
		},
		{
			`0.0--`,
			-1.0,
		},
		{
			`1.0--`,
			0.0,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestPostfixOperator[%d]", i), func(t *testing.T) {

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

func TestInfixOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`1 + 1`,
			2,
		},
		{
			`1.0 + 1.0`,
			2.0,
		},
		{
			`1.0 + 1`,
			2.0,
		},
		{
			`1 + 1.0`,
			2.0,
		},
		{
			`1 - 1`,
			0,
		},
		{
			`1.0 - 1.0`,
			0.0,
		},
		{
			`1 - 1.0`,
			0.0,
		},
		{
			`1.0 - 1`,
			0.0,
		},
		{
			`2 * 2`,
			4,
		},
		{
			`2.0 * 2.0`,
			4.0,
		},
		{
			`2.0 * 2`,
			4.0,
		},
		{
			`2 * 2.0`,
			4.0,
		},
		{
			`2 / 2`,
			1,
		},
		{
			`2.0 / 2.0`,
			1.0,
		},
		{
			`2.0 / 2`,
			1.0,
		},
		{
			`2 / 2.0`,
			1.0,
		},
		{
			`4 % 2`,
			0,
		},
		{
			`4.0 % 2.0`,
			0.0,
		},
		{
			`4 % 2.0`,
			0.0,
		},
		{
			`4.0 % 2`,
			0.0,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestInfixOperator[%d]", i), func(t *testing.T) {

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
	case float64:
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