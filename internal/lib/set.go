package lib

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	elements map[T]struct{}
}

func MakeSet[T comparable](ts ...T) Set[T] {
	s := Set[T]{elements: make(map[T]struct{})}
	for _, t := range ts {
		s.Insert(t)
	}
	return s
}

func (s *Set[T]) Elements(yield func(T) bool) {
	for val := range s.elements {
		if !yield(val) {
			return
		}
	}
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set[T]) Insert(t T) {
	s.elements[t] = struct{}{}
}

func (s *Set[T]) Delete(t T) {
	delete(s.elements, t)
}

func (s *Set[T]) Slice() []T {
	elements := make([]T, 0, len(s.elements))
	for k := range s.Elements {
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
	for val := range s2.Elements {
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

func (s *Set[T]) Pop() T {
	t := s.Slice()[0]
	s.Delete(t)
	return t
}

func (s Set[T]) String() string {
	strs := make([]string, s.Size())
	for i, val := range s.Slice() {
		strs[i] = fmt.Sprint(val)
	}
	return "{" + strings.Join(strs, ", ") + "}"
}
