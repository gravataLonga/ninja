package interpreter_test

import (
	"fmt"
	"testing"

	"github.com/gravataLonga/ninja/object"
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
			`5 ** 2 ** 2`,
			625,
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

			v := evalProgram(t, tt.input)

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
		{
			`2 ** -1`,
			0.5,
		},
		{
			`2 ** -2`,
			0.25,
		},
		{
			`3 ** -1`,
			1.0 / 3.0,
		},
		{
			`var called = false; var se = function() { called = true; return true; }; false && se(); called;`,
			false,
		},
		{
			`var called = false; var se = function() { called = true; return false; }; true || se(); called;`,
			false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestInfixOperator[%d]", i), func(t *testing.T) {

			v := evalProgram(t, tt.input)

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

func TestInfixOperatorErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{`1 % 0`, "integer division by zero"},
		{`-5 % 0`, "integer division by zero"},
		{`0 % 0`, "integer division by zero"},
		{`1 < "a"`, "unknown operator: INTEGER < STRING at [Line: 1, Offset: 3]"},
		{`1 > "a"`, "unknown operator: INTEGER > STRING at [Line: 1, Offset: 3]"},
		{`1 - "a"`, "unknown operator: INTEGER - STRING at [Line: 1, Offset: 3]"},
		{`1 * "a"`, "unknown operator: INTEGER * STRING at [Line: 1, Offset: 3]"},
		{`"a" << 1`, "TypeError: <<() expected argument #1 to be `INTEGER` got `STRING` at [Line: 1, Offset: 6]"}, // @todo uniform all the errors message
		{`1 >> "b"`, "TypeError: <<() expected argument #2 to be `INTEGER` got `STRING` at [Line: 1, Offset: 4]"}, // @todo uniform all the errors message
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestInfixOperatorErrors[%d]", i), func(t *testing.T) {
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
