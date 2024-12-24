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

func (s *Set[T]) Insert(ts ...T) {
	for _, t := range ts {
		s.elements[t] = struct{}{}
	}
}

func (s *Set[T]) Delete(ts ...T) {
	for _, t := range ts {
		delete(s.elements, t)
	}
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

func (s *Set[T]) Clone() Set[T] {
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

func (s1 *Set[T]) Diff(s2 Set[T]) {
	for _, val := range s1.Slice() {
		if s2.Contains(val) {
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
	for val := range s.Elements {
		s.Delete(val)
		return val
	}
	panic("attempt to pop empty set")
}

func (s *Set[T]) Peek() T {
	for val := range s.Elements {
		return val
	}
	panic("attempt to peek empty set")
}

func (s Set[T]) String() string {
	strs := make([]string, s.Size())
	for i, val := range s.Slice() {
		strs[i] = fmt.Sprint(val)
	}
	return "{" + strings.Join(strs, ",") + "}"
}

func SetJoin[T comparable](s1, s2 Set[T]) Set[T] {
	s := MakeSet[T]()
	for v := range s1.Elements {
		s.Insert(v)
	}
	for v := range s2.Elements {
		s.Insert(v)
	}
	return s
}

func SetIntersect[T comparable](s1, s2 Set[T]) Set[T] {
	s := MakeSet[T]()
	for v := range s1.Elements {
		if s2.Contains(v) {
			s.Insert(v)
		}
	}
	return s
}

func SetDiff[T comparable](s1, s2 Set[T]) Set[T] {
	s := MakeSet[T]()
	for v := range s1.Elements {
		if !s2.Contains(v) {
			s.Insert(v)
		}
	}
	return s
}

func SetInsert[T comparable](s Set[T], t T) Set[T] {
	set := s.Clone()
	set.Insert(t)
	return set
}
