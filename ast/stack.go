package ast

// Scope is a map structure that hold a identifier and if is ready
// or not.
type Scope map[string]bool

// Stack will keep track frames of a program, where a specific Scope
// are in analysis
type Stack []*Scope

// IsEmpty check if analysis is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push will push new Scope into analysis, will create a new frame.
func (s *Stack) Push(i *Scope) {
	*s = append(*s, i)
}

// Pop will pop last Scope pushed into Stack
func (s *Stack) Pop() (*Scope, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

// Peek only will peek last Scope pushed into Stack
func (s *Stack) Peek() (*Scope, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(*s) - 1
	scope := (*s)[index]
	return scope, true
}

// Get will get specified index of scope
func (s *Stack) Get(index int) *Scope {
	return (*s)[index]
}

// Size will get total size of analysis
func (s *Stack) Size() int {
	return len(*s)
}

// Put will put a new variable into Scope
func (s *Scope) Put(name string, ready bool) {
	(*s)[name] = ready
}
