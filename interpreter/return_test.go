package interpreter

import (
	"fmt"
	"github.com/gravataLonga/ninja/object"
	"testing"
)

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"return 20;", 20},
		{"return;", object.NULL},
		{"return 20.50;", 20.50},
		{"return 30; 9;", 30},
		{"return 8 * 5; 9;", 40},
		{"9; return 12 * 5; 9;", 60},
		{"if (10 > 1) { return 10; }", 10},
		{"if (10 > 1) { return; }", object.NULL},
		/*
						// @todo we really care return without in function scope?
			{
			`
			if (10 > 1) {
				if (10 > 1) {
					return 90;
				}

				return 1;
			}
			`,
			90,
			},*/
		{
			`
		var f = function(x) {
		  return x;
		  return x + 10;
		};
		f(10);`,
			10,
		},
		{
			`
		var f = function() {
		  return;
		};
		f();`,
			object.NULL,
		},
		{
			`
		var f = function(x) {
		   var result = x + 10;
		   return result;
		   return 10;
		};
		f(10);`,
			20,
		},
		{
			`
		function f(x) {
		  return x;
		  x + 10;
		};
		f(10);`,
			10,
		},
		{
			`
		function f(x) {
		   var result = x + 10;
		   return result;
		   return 10;
		};
		f(10);`,
			20,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestReturnStatements[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)

			integer, ok := tt.expected.(int)
			float, okFloat := tt.expected.(float64)

			if ok {
				testIntegerObject(t, evaluated, int64(integer))
			} else if okFloat {
				testFloatObject(t, evaluated, float64(float))
			}
		})
	}
}
