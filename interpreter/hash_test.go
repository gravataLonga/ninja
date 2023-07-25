package interpreter

import (
	"fmt"
	"github.com/gravataLonga/ninja/object"
	"testing"
)

func TestHashLiterals(t *testing.T) {
	input := `var two = "two";
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 3,
		4: 4,
		true: 5,
		false: 6
	}`

	evaluated := interpreter(t, input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		object.TRUE.HashKey():                      5,
		object.FALSE.HashKey():                     6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestHashLiteralsAssing(t *testing.T) {
	input := `var a = {};
a["hello"] = "world";
a;
`

	evaluated := interpreter(t, input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]string{
		(&object.String{Value: "hello"}).HashKey(): "world",
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testStringObject(t, pair.Value, expectedValue)
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`var key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{true: 5}[true]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
		{
			`{1.0: 35}[1.0]`,
			35,
		},
		{
			`{1.000000: 35}[1.0]`,
			35,
		},
		{
			`{1.000001: 35}[1.0]`,
			nil,
		},
	}

	for o, tt := range tests {
		t.Run(fmt.Sprintf("TestHashIndexExpressions[%d]", o), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)
			integer, ok := tt.expected.(int)
			if ok {
				testIntegerObject(t, evaluated, int64(integer))
			} else {
				testNullObject(t, evaluated)
			}
		})

	}
}

func TestEvalHashExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"{};", object.Hash{}},
		{"!{}", false},
		{"!!{}", true},
		{"{1: 1}[1]", 1},
		{"{1: \"ola ola\"}[1]", "ola ola"},
		{"{1: 2.5}[1]", 2.5},
		{"{1: true}[1]", true},
		{"{1: 1 + 1}[1]", 2},
		{"{1: !true}[1]", false},
		{"{1 + 1: 4}[2]", 4},
		{"{} == {}", false},
		{"{} == {1: 0}", false},
		{"{} != {}", true},
		{"{} != {1: 2}", true},
		{"{} && {}", true},
		{"{} || {}", true},
		{"{} && false", false},
	}

	for o, tt := range tests {
		t.Run(fmt.Sprintf("TestEvalHashExpression[%d]", o), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)
			testObjectLiteral(t, evaluated, tt.expected)
		})

	}
}

func TestErrorHashHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"-{};",
			"unknown operator: -HASH - at [Line: 1, Offset: 1]",
		},
		{
			"{} + {}",
			"unknown operator: HASH + HASH + at [Line: 1, Offset: 4]",
		},
		{
			"{} - {}",
			"unknown operator: HASH - HASH - at [Line: 1, Offset: 4]",
		},
		{
			"{} > {}",
			"unknown operator: HASH > HASH > at [Line: 1, Offset: 4]",
		},
		{
			"{} < {}",
			"unknown operator: HASH < HASH < at [Line: 1, Offset: 4]",
		},
		{
			"{} <= {}",
			"unknown operator: HASH <= HASH <= at [Line: 1, Offset: 5]",
		},
		{
			"{} >= {}",
			"unknown operator: HASH >= HASH >= at [Line: 1, Offset: 5]",
		},
	}

	for o, tt := range tests {
		t.Run(fmt.Sprintf("TestErrorHashHandling[%d]", o), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
				return
			}

			if errObj.Message != tt.expectedMessage {
				t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
			}
		})

	}
}

func TestHashMethod(t *testing.T) {

	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{}.type()`,
			"HASH",
		},
		{
			`{}.keys()`,
			object.Array{Elements: []object.Object{}},
		},
		{
			`{}.has("a")`,
			false,
		},
		{
			`{"a": 1}.has("a")`,
			true,
		},
		{
			`{"a": 1}.has("b")`,
			false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestHashMethod_%d", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)

			testObjectLiteral(t, evaluated, tt.expected)
		})
	}
}

func TestHashHasKeyMethod(t *testing.T) {
	input := `{"a": 1, "b": true}.keys()`
	slicesContain := []string{"a", "b"}

	contain := func(slice []string, input string) bool {
		for _, v := range slice {
			if v == input {
				return true
			}
		}
		return false
	}

	evaluated := interpreter(t, input)

	arr, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("expect to be an array. Got: %T", evaluated)
	}

	stringPos1, ok := arr.Elements[0].(*object.String)
	if !ok {
		t.Fatalf("expect index 0 be a string. Got: %T", evaluated)
	}

	stringPos2, ok := arr.Elements[0].(*object.String)
	if !ok {
		t.Fatalf("expect index 0 be a string. Got: %T", evaluated)
	}

	if !contain(slicesContain, stringPos1.Value) {
		t.Fatalf("hash.keys() don't contain %v. Got: %v", slicesContain, stringPos1)
	}

	if !contain(slicesContain, stringPos2.Value) {
		t.Fatalf("hash.keys() don't contain %v. Got: %v", slicesContain, stringPos2)
	}
}

func TestHashHasValueMethod(t *testing.T) {
	input := `{"a": "a1", "b": "b2"}.values()`
	slicesContain := []string{"a1", "b2"}

	contain := func(slice []string, input string) bool {
		for _, v := range slice {
			if v == input {
				return true
			}
		}
		return false
	}

	evaluated := interpreter(t, input)

	arr, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("expect to be an array. Got: %T", evaluated)
	}

	stringPos1, ok := arr.Elements[0].(*object.String)
	if !ok {
		t.Fatalf("expect index 0 be a string. Got: %T", evaluated)
	}

	stringPos2, ok := arr.Elements[0].(*object.String)
	if !ok {
		t.Fatalf("expect index 0 be a string. Got: %T", evaluated)
	}

	if !contain(slicesContain, stringPos1.Value) {
		t.Fatalf("hash.keys() don't contain %v. Got: %v", slicesContain, stringPos1)
	}

	if !contain(slicesContain, stringPos2.Value) {
		t.Fatalf("hash.keys() don't contain %v. Got: %v", slicesContain, stringPos2)
	}
}

func TestHashHasMergeMethod(t *testing.T) {
	input := `{"a": 1, "b": 2}.merge({"c": 3})`
	contain := `{"a": 1, "b": 2, "c": 3}`

	evaluated := interpreter(t, input)

	expected := object.Hash{Pairs: make(map[object.HashKey]object.HashPair)}
	aKey := &object.String{Value: "a"}
	bKey := &object.String{Value: "b"}
	cKey := &object.String{Value: "c"}
	aValue := &object.Integer{Value: 1}
	bValue := &object.Integer{Value: 2}
	cValue := &object.Integer{Value: 3}

	expected.Pairs[aKey.HashKey()] = object.HashPair{Key: aKey, Value: aValue}
	expected.Pairs[bKey.HashKey()] = object.HashPair{Key: bKey, Value: bValue}
	expected.Pairs[cKey.HashKey()] = object.HashPair{Key: cKey, Value: cValue}

	if !testObjectLiteral(t, evaluated, expected) {
		t.Fatalf("Expected %s, got: %s", contain, evaluated.Inspect())
	}
}

func TestHashMethodWrongUsage(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{
			`{}.keys(1)`,
			"TypeError: hash.keys() takes exactly 0 argument (1 given)",
		},
		{
			`{}.has(1, 2)`,
			"TypeError: hash.has() takes exactly 1 argument (2 given)",
		},
		{
			`{}.has()`,
			"TypeError: hash.has() takes exactly 1 argument (0 given)",
		},
		{
			`{}.type(1)`,
			"TypeError: hash.type() takes exactly 0 argument (1 given)",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestHashMethodWrongUsage[%d]", i), func(t *testing.T) {
			evaluated := interpreter(t, tt.input)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			if errObj.Message != tt.expectedErrorMessage {
				t.Errorf("erro expected \"%s\". Got: %s", tt.expectedErrorMessage, errObj.Message)
			}
		})

	}
}
