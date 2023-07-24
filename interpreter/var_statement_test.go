package interpreter

import (
	"fmt"
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
			i := New(os.Stdout, object.NewEnvironment())

			v := i.Interpreter(nodes)

			if len(tt.expectedEnvironment) > 0 {
				for k, v := range tt.expectedEnvironment {
					vEnv, ok := i.env.Get(k)
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
