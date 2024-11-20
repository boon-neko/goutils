package que

// Stack implementation
type Stack[T any] struct {
	data []T
	size int
}

// NewStack instantiates a new stack
func NewStack[T any](cap int) *Stack[T] { return &Stack[T]{data: make([]T, 0, cap), size: 0} }

// Push adds a new element at the end of the stack
func (s *Stack[T]) Push(d T) {
	s.data = append(s.data, d)
	s.size++
}

// Pop removes the last element from stack
func (s *Stack[T]) Pop() bool {
	if s.IsEmpty() {
		return false
	}
	s.size--
	s.data = s.data[:s.size]
	return true
}

// Top returns the last element of stack
func (s *Stack[T]) Top() T {
	return s.data[s.size-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}
