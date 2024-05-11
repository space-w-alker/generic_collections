package stack

type Stack[T any] []T

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Push(elem T) {
	*s = append(*s, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	index := len(*s) - 1
	elem := (*s)[index]
	*s = (*s)[:index]
	return elem, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	return (*s)[len(*s)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
