package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}

	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestIsTruthy(t *testing.T) {
	tests := []struct {
		obj      Object
		expected bool
	}{
		{
			TRUE,
			true,
		},
		{
			NULL,
			false,
		},
		{
			&Null{},
			false,
		},
		{
			&Array{},
			true,
		},
		{
			&Hash{},
			true,
		},
		{
			&String{},
			true,
		},
		{
			&Integer{},
			true,
		},
		{
			&Integer{Value: 0},
			true,
		},
		{
			&Float{},
			true,
		},
		{
			TRUE,
			true,
		},
		{
			FALSE,
			false,
		},
		{
			&Boolean{Value: true},
			true,
		},
		{
			&Boolean{Value: false},
			false,
		},
	}

	for _, tt := range tests {

		if tt.expected != IsTruthy(tt.obj) {
			t.Errorf("%s expected to be %t. Got: %t", tt.obj.Inspect(), tt.expected, IsTruthy(tt.obj))
		}
	}
}
