package interpreter_test

import (
	"fmt"
	"github.com/gravataLonga/ninja/interpreter"
	"github.com/gravataLonga/ninja/object"
	"os"
	"testing"
)

func TestVarStmt(t *testing.T) {
	tests := []struct {
		input               string
		expectedEnvironment map[string]interface{}
		expected            interface{}
		expectedNil         bool
	}{
		{
			`var a = "ninja";`,
			map[string]interface{}{"a": "ninja"},
			nil,
			true,
		},
		{
			`var a = "ninja"; var b = 2022;`,
			map[string]interface{}{"a": "ninja", "b": 2022},
			nil,
			true,
		},
		{
			`var a = "ninja"; a;`,
			map[string]interface{}{"a": "ninja"},
			"ninja",
			false,
		},
		// {
		//	`var a = "ninja"; b;`,
		//	map[string]interface{}{"a": "ninja"},
		//	nil,
		//	false,
		//},
		{
			`var a = "ninja"; a = "hello";`,
			map[string]interface{}{"a": "hello"},
			nil,
			true,
		},
		{
			`var a = "ninja"; a = "hello"; a;`,
			map[string]interface{}{"a": "hello"},
			"hello",
			false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestVarStmt[%d]", i), func(t *testing.T) {

			nodes := createParser(t, tt.input)
			env := object.NewEnvironment()
			i := interpreter.New(os.Stdout, env)

			v := i.Interpreter(nodes)

			if len(tt.expectedEnvironment) > 0 {
				for k, v := range tt.expectedEnvironment {
					vEnv, ok := env.Get(k)
					if !ok {
						t.Errorf("not found env on interpreter")
						continue
					}

					if !testLiteralObject(t, vEnv, v) {
						t.Errorf("value isn't equal on enviroment at interpreter")
					}
				}
			}

			if v == nil && tt.expected != nil {
				t.Fatalf("Interpreter return nil as result")
			}

			if _, ok := v.(*object.Error); ok {
				t.Fatalf("Interpreter return error. %s", v.Inspect())
			}

			if !tt.expectedNil && !testLiteralObject(t, v, tt.expected) {
				t.Fatalf("testLiteralObject got false, expected true.")
			}
		})
	}
}

func TestAssignDotPropertyOnHash(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`var h = {}; h.name = "ninja"; h.name;`,
			"ninja",
		},
		{
			`var h = {"name": "old"}; h.name = "new"; h.name;`,
			"new",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestAssignDotPropertyOnHash[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)

			testStringObject(t, evaluated, tt.expected)
		})
	}
}

func TestAssignDotPropertyOnNonHash(t *testing.T) {
	// `h.name() = 1` is covered by TestIllegalAssignmentsErrors (parser/assign_test.go)
	// instead: with `.` binding tighter than a call, its LHS is a *ast.CallExpression
	// (illegal assignment target altogether), never reaching the *ast.Dot write path.
	input := `var a = [1]; a.x = 1;`

	evaluated := evalProgram(t, input)

	if _, ok := evaluated.(*object.Error); !ok {
		t.Fatalf("expected *object.Error. got=%T (%+v)", evaluated, evaluated)
	}
}

func TestIndexAssignOnNonCollection(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{`var x = 5; x[0] = 1;`, `cannot assign property "x" on INTEGER`},
		{`var x = "hello"; x[0] = "y";`, `cannot assign property "x" on STRING`},
		{`var x = true; x[0] = 1;`, `cannot assign property "x" on BOOLEAN`},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestIndexAssignOnNonCollection[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
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

func TestAssignErrorPropagation(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{`var x = 1; x = "a" + 1;`, "unknown operator: STRING + INTEGER at [Line: 1, Offset: 20]"},
		{`var x = 1; x = 1 > "a";`, "unknown operator: INTEGER > STRING at [Line: 1, Offset: 18]"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestAssignErrorPropagation[%d]", i), func(t *testing.T) {
			evaluated := evalProgram(t, tt.input)
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
