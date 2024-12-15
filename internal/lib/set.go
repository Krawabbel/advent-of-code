package lib

type Set[T comparable] struct {
	elements map[T]struct{}
}

func (s *Set[T]) Insert(t T) {
	s.elements[t] = struct{}{}
}

func (s *Set[T]) Delete(t T) {
	delete(s.elements, t)
}

func (s *Set[T]) Slice() []T {
	elements := make([]T, 0, len(s.elements))
	for k := range s.elements {
		elements = append(elements, k)
	}
	return elements
}

func (s *Set[T]) Contains(t T) bool {
	_, found := s.elements[t]
	return found
}

func (s *Set[T]) Clone(t T) Set[T] {
	return Set[T]{elements: CloneMap(s.elements)}
}

func (s1 *Set[T]) Join(s2 Set[T]) {
	for val := range s2.elements {
		s1.Insert(val)
	}
}

func (s1 *Set[T]) Intersect(s2 Set[T]) {
	for _, val := range s1.Slice() {
		if !s2.Contains(val) {
			s1.Delete(val)
		}
	}
}

func (s *Set[T]) Modify(f func(t T) T) {
	for _, val := range s.Slice() {
		s.Delete(val)
		s.Insert(f(val))
	}
}
