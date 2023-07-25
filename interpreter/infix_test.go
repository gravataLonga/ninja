package interpreter

import (
	"fmt"
	"github.com/gravataLonga/ninja/object"
	"testing"
)

func TestInfixMathOperator(t *testing.T) {
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
			`"hello" + " ninja"`,
			"hello ninja",
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
			1.0,
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
		{
			`10 ** 0`,
			1,
		},
		{
			`10.0 ** 0.0`,
			1.0,
		},
		{
			`10.0 ** 0`,
			1.0,
		},
		{
			`10 ** 0.0`,
			1.0,
		},
		{
			`1 | 0`,
			1,
		},
		{
			`1 & 0`,
			0,
		},
		{
			`1 ^ 0`,
			1,
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

func TestInfixLogicOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`1 == 1`,
			true,
		},
		{
			`0 == 1`,
			false,
		},
		{
			`1 == 0`,
			false,
		},
		{
			`1 == 0`,
			false,
		},
		{
			`1.0 == 1.0`,
			true,
		},
		{
			`1.0 == 0.0`,
			false,
		},
		{
			`0.0 == 1.0`,
			false,
		},
		{
			`1 == 1.0`,
			true,
		},
		{
			`1.0 == 1`,
			true,
		},
		{
			`1.0 == 1`,
			true,
		},
		{
			`true == true`,
			true,
		},
		{
			`true == false`,
			false,
		},
		{
			`true == "ninja"`,
			false,
		},
		{
			`true == 1`,
			false,
		},
		{
			`true == 1.0`,
			false,
		},
		{
			`"ninja" == true`,
			false,
		},
		{
			`1 == true`,
			false,
		},
		{
			`1.0 == true`,
			false,
		},
		{
			`"ninja" == "ninja"`,
			true,
		},
		{
			`"ninja" == "hello"`,
			false,
		},
		{
			`"ninja" == 1`,
			false,
		},
		{
			`1 == "ninja"`,
			false,
		},
		{
			`1.0 == "ninja"`,
			false,
		},
		{
			`"ninja" == 1.0`,
			false,
		},
		{
			`[] == []`,
			false,
		},
		{
			`{} == {}`,
			false,
		},
		{
			`1 != 1`,
			false,
		},
		{
			`0 != 1`,
			true,
		},
		{
			`1.0 != 1.0`,
			false,
		},
		{
			`1 != 1.0`,
			false,
		},
		{
			`1.0 != 1`,
			false,
		},
		{
			`"ninja" != "wow"`,
			true,
		},
		{
			`"ninja" != "ninja"`,
			false,
		},
		{
			`"ninja" != "ninja"`,
			false,
		},
		{
			`"ninja" != 1`,
			true,
		},
		{
			`1 != "ninja"`,
			true,
		},
		{
			`1.0 != "ninja"`,
			true,
		},
		{
			`"ninja" != 1.0`,
			true,
		},
		{
			`"ninja" != []`,
			true,
		},
		{
			`[] != "ninja"`,
			true,
		},
		{
			`"ninja" != {}`,
			true,
		},
		{
			`{} != "ninja"`,
			true,
		},
		{
			`1 != []`,
			true,
		},
		{
			`[] != 1`,
			true,
		},
		{
			`1 != {}`,
			true,
		},
		{
			`{} != 1`,
			true,
		},
		{
			`1.0 != []`,
			true,
		},
		{
			`[] != 1.0`,
			true,
		},
		{
			`1.0 != {}`,
			true,
		},
		{
			`{} != 1.0`,
			true,
		},
		{
			`true != true`,
			false,
		},
		{
			`false != true`,
			true,
		},
		{
			`true != 1`,
			true,
		},
		{
			`true != 1.0`,
			true,
		},
		{
			`true != []`,
			true,
		},
		{
			`true != {}`,
			true,
		},

		{
			`1 != true`,
			true,
		},
		{
			`1.0 != true`,
			true,
		},
		{
			`[] != true`,
			true,
		},
		{
			`{} != true`,
			true,
		},

		{
			`1 < 1`,
			false,
		},
		{
			`0 < 1`,
			true,
		},
		{
			`1.0 < 1.0`,
			false,
		},
		{
			`0.0 < 1.0`,
			true,
		},
		{
			`1 < 1.0`,
			false,
		},
		{
			`0 < 1.0`,
			true,
		},
		{
			`1.0 < 1`,
			false,
		},
		{
			`1.0 < 0`,
			false,
		},

		{
			`1 > 1`,
			false,
		},
		{
			`0 > 1`,
			false,
		},
		{
			`1.0 > 1.0`,
			false,
		},
		{
			`0.0 > 1.0`,
			false,
		},
		{
			`1 > 1.0`,
			false,
		},
		{
			`0 > 1.0`,
			false,
		},
		{
			`1.0 > 1`,
			false,
		},
		{
			`1.0 > 0`,
			true,
		},

		{
			`1 <= 1`,
			true,
		},
		{
			`0 <= 1`,
			true,
		},
		{
			`1.0 <= 1.0`,
			true,
		},
		{
			`0.0 <= 1.0`,
			true,
		},
		{
			`1 <= 1.0`,
			true,
		},
		{
			`0 <= 1.0`,
			true,
		},
		{
			`1.0 <= 1`,
			true,
		},
		{
			`1.0 <= 0`,
			false,
		},

		{
			`1 >= 1`,
			true,
		},
		{
			`0 >= 1`,
			false,
		},
		{
			`1.0 >= 1.0`,
			true,
		},
		{
			`0.0 >= 1.0`,
			false,
		},
		{
			`1 >= 1.0`,
			true,
		},
		{
			`0 >= 1.0`,
			false,
		},
		{
			`1.0 >= 1`,
			true,
		},
		{
			`1.0 >= 0`,
			true,
		},

		{
			"1 && 1",
			true,
		},
		{
			"1.0 && 1",
			true,
		},
		{
			"1 && 1.0",
			true,
		},
		{
			"[] && 1.0",
			true,
		},
		{
			"1 && []",
			true,
		},
		{
			"{} && 1.0",
			true,
		},
		{
			"1 && {}",
			true,
		},
		{
			`1 && true`,
			true,
		},
		{
			`1 && false`,
			false,
		},
		{
			`false && 1`,
			false,
		},
		{
			`1.0 && true`,
			true,
		},
		{
			`1.0 && false`,
			false,
		},
		{
			`false && 1.0`,
			false,
		},
		{
			`[] && true`,
			true,
		},
		{
			`[] && false`,
			false,
		},
		{
			`false && []`,
			false,
		},
		{
			`{} && true`,
			true,
		},
		{
			`{} && false`,
			false,
		},
		{
			`false && {}`,
			false,
		},
		{
			`false && false`,
			false,
		},

		{
			"1 || 1",
			true,
		},
		{
			"1.0 || 1",
			true,
		},
		{
			"1 || 1.0",
			true,
		},
		{
			"[] || 1.0",
			true,
		},
		{
			"1 || []",
			true,
		},
		{
			"{} || 1.0",
			true,
		},
		{
			"1 || {}",
			true,
		},
		{
			`1 || true`,
			true,
		},
		{
			`1 || false`,
			true,
		},
		{
			`false || 1`,
			true,
		},
		{
			`1.0 || true`,
			true,
		},
		{
			`1.0 || false`,
			true,
		},
		{
			`false || 1.0`,
			true,
		},
		{
			`[] || true`,
			true,
		},
		{
			`[] || false`,
			true,
		},
		{
			`false || []`,
			true,
		},
		{
			`{} || true`,
			true,
		},
		{
			`{} || false`,
			true,
		},
		{
			`false || {}`,
			true,
		},
		{
			`false || false`,
			false,
		},

		{
			`1 << 1`,
			2,
		},
		{
			`2 >> 1`,
			1,
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
