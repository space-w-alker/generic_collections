package set

type Set[T comparable] map[T]bool

func (s *Set[T]) Add(element T) {
	_, ok := (*s)[element]
	if ok {
		return
	}
	(*s)[element] = true
}

func (s *Set[T]) Remove(element T) {
	delete((*s), element)
}

func (s *Set[T]) IsEmpty() (ok bool) {
	return len(*s) == 0
}

func (s *Set[T]) Size() int {
	return len(*s)
}

func (s *Set[T]) Contains(element T) (ok bool) {
	_, ok = (*s)[element]
	return
}

func (s *Set[T]) Clear(element T) {
	for k := range *s {
		delete(*s, k)
	}
}

func (s *Set[T]) Union(other *Set[T]) {
	for v := range *other {
		s.Add(v)
	}
}

func (s *Set[T]) Intersect(other *Set[T]) {
	for v := range *s {
		if !other.Contains(v) {
			s.Remove(v)
		}
	}
}

func (s *Set[T]) Difference(other *Set[T]) {
	for v := range *s {
		if other.Contains(v) {
			s.Remove(v)
		}
	}
}
