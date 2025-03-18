package promqlex

type Stack[T any] struct {
	elements []T
}

// Push adds an element to the stack
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element of the stack
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zeroVal T
		return zeroVal, false
	}
	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value, true
}

func (s *Stack[T]) MustPop() T {
	value, ok := s.Pop()
	if !ok {
		panic("stack is empty")
	}
	return value
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zeroVal T
		return zeroVal, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *Stack[T]) MustPeek() T {
	value, ok := s.Peek()
	if !ok {
		panic("stack is empty")
	}
	return value
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
