package interpreter

import (
	"fmt"
	"github.com/gravataLonga/ninja/object"
	"testing"
)

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
		nullable bool
	}{
		{`len("")`, 0, false},
		{`len("four")`, 4, false},
		{`len("hello world")`, 11, false},
		{`len(1)`, "TypeError: len() expected argument to be `ARRAY,STRING` got `INTEGER`", false},
		{`len("one", "two")`, "TypeError: len() takes exactly 1 argument (2 given)", false},
		{`len([1, 2, 3])`, 3, false},
		{`len([])`, 0, false},
		{`puts("hello", "world!")`, nil, true},
		{`first([1, 2, 3])`, 1, false},
		{`first([])`, nil, false},
		{`first(1)`, "TypeError: first() expected argument #1 to be `ARRAY` got `INTEGER`", false},
		// builtin function last must be immutable
		{`var a = [[0, 1]];var b = first(a);b[0] = b[0] + 1; a[0][0];`, 0, false},
		{`last([1, 2, 3])`, 3, false},
		// builtin function last must be immutable
		{`var a = [[0, 1]];var b = last(a);b[0] = b[0] + 1; a[0][0];`, 0, false},
		{`last([])`, nil, false},
		{`last(1)`, "TypeError: last() expected argument #1 to be `ARRAY` got `INTEGER`", false},
		{`rest([1, 2, 3])`, []int{2, 3}, false},
		// builtin function last must be immutable
		{`var a = [[0, 1], [0, 1]];var b = rest(a);b[0] = b[0] + 1; a[1][0];`, 0, false},
		{`rest([])`, nil, false},
		{`push([], 1)`, []int{1}, false},
		{`push(1, 1)`, "TypeError: push() expected argument #1 to be `ARRAY` got `INTEGER`", false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestBuiltinFunctions[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)

			switch expected := tt.expected.(type) {
			case int:
				testIntegerObject(t, evaluated, int64(expected))
			case nil:
				if tt.nullable {
					if evaluated != nil {
						t.Fatalf("Test must return nil. Got: %T", evaluated)
					}
				} else {
					testNullObject(t, evaluated)
				}

			case string:
				errObj, ok := evaluated.(*object.Error)
				if !ok {
					t.Fatalf("object is not Error. got=%T (%+v)",
						evaluated, evaluated)
				}
				if errObj.Message != expected {
					t.Errorf("wrong error message. expected=%q, got=%q",
						expected, errObj.Message)
				}
			case []int:
				array, ok := evaluated.(*object.Array)
				if !ok {
					t.Fatalf("obj not Array. got=%T (%+v)", evaluated, evaluated)
				}

				if len(array.Elements) != len(expected) {
					t.Fatalf("wrong num of elements. want=%d, got=%d",
						len(expected), len(array.Elements))
				}

				for i, expectedElem := range expected {
					testIntegerObject(t, array.Elements[i], int64(expectedElem))
				}
			}
		})
	}
}

func TestBuiltinTime(t *testing.T) {
	evaluated := interpreter(t, `time()`)

	if _, ok := evaluated.(*object.Integer); !ok {
		t.Fatalf("builtin time() expected got integer. Got: %T", evaluated)
	}
}

func TestBuiltinRand(t *testing.T) {
	evaluated := interpreter(t, `rand()`)

	if _, ok := evaluated.(*object.Float); !ok {
		t.Fatalf("builtin rand() expected got float. Got: %T", evaluated)
	}
}

func TestArgs(t *testing.T) {
	object.Arguments = []string{"test", "hello"}
	input := `args();`

	evaluated := interpreter(t, input)

	arr, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("TestArgs expected array got %s", evaluated.Type())
	}

	if len(arr.Elements) != 2 {
		t.Errorf("Elements expect to be 2 length. Got: %d", len(arr.Elements))
	}

	arg1, _ := arr.Elements[0].(*object.String)
	arg2, _ := arr.Elements[1].(*object.String)

	if arg1.Value != "test" || arg2.Value != "hello" {
		t.Errorf("state of env isn't equal")
	}
}

func TestPlugin(t *testing.T) {
	input := `plugin("../fixtures/hello")`

	evaluated := interpreter(t, input)

	_, ok := evaluated.(*object.Plugin)
	if !ok {
		t.Fatalf("TestPlugin expected plugin got %s (%s)", evaluated.Type(), evaluated.Inspect())
	}

	_, ok = evaluated.(object.CallableMethod)
	if !ok {
		t.Fatalf("TestPlugin expected implement CallableMethod")
	}
}

func TestPluginCallSymbols(t *testing.T) {
	input := `var h = plugin("../fixtures/hello"); h.hello();`

	evaluated := interpreter(t, input)

	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("TestPlugin expected plugin got %s (%s)", evaluated.Type(), evaluated.Inspect())
	}

	if str.Value != "Hello World!" {
		t.Fatalf("TestPluginCallSymbols expected \"Hello World!\" got %s", str.Value)
	}
}
