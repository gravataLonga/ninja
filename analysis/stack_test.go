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
