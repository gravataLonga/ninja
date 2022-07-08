package object

import (
	"fmt"
	"testing"
)

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

func TestHahsableString(t *testing.T) {
	s := &String{Value: "Hello World"}
	i := &Integer{Value: 1}
	f := &Float{Value: 1}

	tests := []struct {
		oLeft   Object
		oRight  Object
		isEqual bool
	}{
		{
			&String{Value: "Hello World"},
			&String{Value: "Hello World"},
			false,
		},
		{
			s,
			s,
			true,
		},
		{
			&Integer{Value: 1},
			&Integer{Value: 1},
			false,
		},
		{
			i,
			i,
			true,
		},
		{
			&Float{Value: 1},
			&Float{Value: 1},
			false,
		},
		{
			f,
			f,
			true,
		},
		{
			&Boolean{Value: true},
			&Boolean{Value: true},
			false,
		},
		{
			TRUE,
			TRUE,
			true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestHahsableString[%d]", i), func(t *testing.T) {
			hashLeft, _ := tt.oLeft.(Hashable)
			hashRight, _ := tt.oRight.(Hashable)
			equal := hashLeft == hashRight

			if equal != tt.isEqual {
				t.Fatalf("Object %s and %s are %v. Expected %v", tt.oLeft.Inspect(), tt.oRight.Inspect(), equal, tt.isEqual)
			}

		})
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
		{
			nil,
			false,
		},
	}

	for _, tt := range tests {

		if tt.expected != IsTruthy(tt.obj) {
			t.Errorf("%s expected to be %t. Got: %t", tt.obj.Inspect(), tt.expected, IsTruthy(tt.obj))
		}
	}
}
