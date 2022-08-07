package interpreter

import (
	"fmt"
	"github.com/gravataLonga/ninja/object"
	"testing"
)

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
