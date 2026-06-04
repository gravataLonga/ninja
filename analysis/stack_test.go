package analysis

import "testing"

func Test(t *testing.T) {
	s := NewStack()
	scope := NewScope()

	if s.Size() != 0 {
		t.Errorf("Size of an empty analysis should be 0")
	}

	s.Push(scope)

	if s.Size() != 1 {
		t.Errorf("Size shouldn't be equal to 1")
	}

	if s.Peek() != scope {
		t.Errorf("Peek analysis should scope.")
	}

	if s.Pop() != scope {
		t.Errorf("Pop item must be same scope.")
	}

	if s.Size() != 0 {
		t.Errorf("Size  should be 0.")
	}

	s.Push(scope)
	s.Push(scope)

	if s.Size() != 2 {
		t.Errorf("Size should be 2")
	}

	if s.Peek() != scope {
		t.Errorf("Peek analysis sould be same scope.")
	}
}

func TestAt(t *testing.T) {
	s := NewStack()
	bottom := NewScope()
	bottom.Put("bottom", true)
	mid := NewScope()
	mid.Put("mid", true)
	top := NewScope()
	top.Put("top", true)

	s.Push(bottom)
	s.Push(mid)
	s.Push(top)

	if s.At(0) != bottom {
		t.Errorf("At(0) should return bottom scope")
	}
	if s.At(1) != mid {
		t.Errorf("At(1) should return mid scope")
	}
	if s.At(2) != top {
		t.Errorf("At(2) should return top scope (same as Peek)")
	}
	if s.At(2) != s.Peek() {
		t.Errorf("At(Size()-1) should equal Peek()")
	}
	if s.At(-1) != nil {
		t.Errorf("At(-1) should return nil")
	}
	if s.At(3) != nil {
		t.Errorf("At(out-of-bounds) should return nil")
	}
}
