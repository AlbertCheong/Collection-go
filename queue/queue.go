package queue

import (
	"container/list"
	"fmt"
)

type Queue[T any] struct {
	elements *list.List
}

func NewQueue[T any](vl ...T) *Queue[T] {
	q := &Queue[T]{elements: list.New()}
	for _, v := range vl {
		q.elements.PushBack(v)
	}
	return q
}

func (q *Queue[T]) Len() int {
	return q.elements.Len()
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) Zero() T {
	var zeroValue T
	return zeroValue
}

func (q *Queue[T]) Enqueue(v T) {
	q.elements.PushBack(v)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.Empty() {
		return q.Zero(), fmt.Errorf("unable to dequeue: Queue is empty")
	}
	front := q.elements.Front()
	q.elements.Remove(front)
	return front.Value.(T), nil
}

func (q *Queue[T]) Front() (T, error) {
	if q.Empty() {
		return q.Zero(), fmt.Errorf("unable to retrieve front value: Queue is empty")
	}
	return q.elements.Front().Value.(T), nil
}
