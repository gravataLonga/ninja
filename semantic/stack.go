package semantic

type Scope map[string]bool

type Stack []*Scope

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(i *Scope) {
	*s = append(*s, i)
}

func (s *Stack) Pop() (*Scope, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *Stack) Peek() (*Scope, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(*s) - 1
	scope := (*s)[index]
	return scope, true
}

func (s *Scope) Put(name string, ready bool) {
	(*s)[name] = ready
}

func (s *Stack) Get(index int) *Scope {
	return (*s)[index]
}

func (s *Stack) Size() int {
	return len(*s)
}
