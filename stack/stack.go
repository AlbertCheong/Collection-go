package stack

import (
	"container/list"
	"fmt"
)

type Stack[T any] struct {
	elements *list.List
}

func NewStack[T any](vl ...T) *Stack[T] {
	s := &Stack[T]{elements: list.New()}
	for _, v := range vl {
		s.elements.PushBack(v)
	}
	return s
}

func (s *Stack[T]) Len() int {
	return s.elements.Len()
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Zero() T {
	var zeroValue T
	return zeroValue
}

func (s *Stack[T]) Push(v T) {
	s.elements.PushBack(v)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		return s.Zero(), fmt.Errorf("unable to pop: Stack is empty")
	}
	v := s.elements.Back()
	s.elements.Remove(v)
	return v.Value.(T), nil
}

func (s *Stack[T]) Top() (T, error) {
	if s.Empty() {
		return s.Zero(), fmt.Errorf("unable to retrieve top value: Stack is empty")
	}
	return s.elements.Back().Value.(T), nil
}
