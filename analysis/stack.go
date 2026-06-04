package analysis

type Stack struct {
	head *leaf
	size int
}

type leaf struct {
	value *Scope
	prev  *leaf
}

// NewStack Create a new stack, with size 0 and head nil
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// Size Return the number of items
func (s *Stack) Size() int {
	return s.size
}

// Peek View most recent push item
func (s *Stack) Peek() *Scope {
	if s.size == 0 {
		return nil
	}
	return s.head.value
}

// Pop remove head element from stack
func (s *Stack) Pop() *Scope {
	if s.size == 0 {
		return nil
	}

	n := s.head
	s.head = n.prev
	s.size--
	return n.value
}

// At returns the scope at position index (0 = bottom/oldest, Size()-1 = top/newest)
func (s *Stack) At(index int) *Scope {
	if index < 0 || index >= s.size {
		return nil
	}
	n := s.head
	for i := 0; i < s.size-1-index; i++ {
		n = n.prev
	}
	return n.value
}

// Push a value onto the head of the analysis
func (s *Stack) Push(value *Scope) {
	n := &leaf{value, s.head}
	s.head = n
	s.size++
}
