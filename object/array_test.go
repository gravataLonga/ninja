package object

import "testing"

func TestArray_Inspect(t *testing.T) {
	arr := &Array{Elements: []Object{&Integer{Value: 1}, &String{Value: "Hello"}, TRUE, NULL, &Float{Value: 1}}}

	if arr.Inspect() != "[1, Hello, true, null, 1.000000]" {
		t.Fatalf("Array.Inspect() expected [1, \"Hello\"]. Got: %s", arr.Inspect())
	}
}

func TestArray_Clone(t *testing.T) {
	arr := &Array{Elements: []Object{&Integer{Value: 1}, &String{Value: "Hello"}, TRUE, NULL, &Float{Value: 1}}}
	arrClone := arr.Clone()

	arr1, ok := arrClone.(*Array)
	if !ok {
		t.Fatalf("Array.Clone() didn't produce an array")
	}

	if arr == arr1 {
		t.Fatalf("Array.Clone() didn't clone array")
	}
}

func TestArray_CloneDeep(t *testing.T) {
	// array[array[1]]
	arr := &Array{Elements: []Object{&Array{Elements: []Object{&Integer{Value: 1}}}}}
	arrClone := arr.Clone()

	arr1, ok := arrClone.(*Array)
	if !ok {
		t.Fatalf("Array.Clone() didn't produce an array")
	}

	arrDeep, _ := arr.Elements[0].(*Array)
	// this will work but, if we change directly value on Elements[0].Right = 2 isn't working
	arrDeep.Elements[0] = &Integer{Value: 2}

	if arr.Inspect() == arr1.Inspect() {
		t.Fatalf("Array.Clone() didn't clone array deep values, %s == %s", arr.Inspect(), arr1.Inspect())
	}
}

func BenchmarkArrayInspect(b *testing.B) {
	arr := &Array{Elements: []Object{&Integer{Value: 1}, &String{Value: "Hello"}, TRUE, NULL, &Float{Value: 1}}}
	for n := 0; n < b.N; n++ {
		arr.Inspect()
	}
}
