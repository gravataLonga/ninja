package semantic

import (
	"fmt"
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		st       Stack
		expected bool
	}{
		{
			Stack{},
			true,
		},
		{
			Stack{&Scope{"hello": false}},
			false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestStack_IsEmpty[%d]", i), func(t *testing.T) {
			if tt.st.IsEmpty() != tt.expected {
				t.Errorf("Stack expected to be empty %v. Got %v", tt.st.IsEmpty(), tt.expected)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {

	stack := &Stack{}
	scope := Scope{"hello": false}

	stack.Push(&scope)
	popScope, ok := stack.Pop()

	if !ok {
		t.Fatalf("Unable to pop elements from Stack")
	}

	if popScope != &scope {
		t.Errorf("Popped scope isn't equal of scope created!")
	}

	if !stack.IsEmpty() {
		t.Errorf("Stack isn't empty")
	}
}

func TestStack_Peek(t *testing.T) {

	stack := &Stack{}
	scope := Scope{"hello": false}

	stack.Push(&scope)
	peekScope, ok := stack.Peek()

	if !ok {
		t.Fatalf("Unable to peek scope")
	}

	if stack.IsEmpty() {
		t.Errorf("Stack can't be empty, we are only peek")
	}

	if peekScope != &scope {
		t.Errorf("Peeked scope isn't equal")
	}

	*peekScope = Scope{"world": true}

	pooppedScope, _ := stack.Pop()

	if v := (*pooppedScope)["world"]; !v {
		t.Errorf("Popped value wasn't change")
	}
}

func TestScope_Put(t *testing.T) {
	scope := Scope{}
	scope.Put("test", true)

	v := scope["test"]

	if !v {
		t.Fatalf("scope test isn't true")
	}
}
