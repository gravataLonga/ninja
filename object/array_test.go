package object

import "testing"

func TestArray_Inspect(t *testing.T) {
	arr := &Array{Elements: []Object{&Integer{Value: 1}, &String{Value: "Hello"}}}

	if arr.Inspect() == "[1, \"Hello\"]" {
		t.Fatalf("Array.Inspect() expected [1, \"Hello\"]. Got: %s", arr.Inspect())
	}
}

func BenchmarkArrayInspect(b *testing.B) {
	arr := &Array{Elements: []Object{&Integer{Value: 1}, &String{Value: "Hello"}}}
	for n := 0; n < b.N; n++ {
		arr.Inspect()
	}
}
