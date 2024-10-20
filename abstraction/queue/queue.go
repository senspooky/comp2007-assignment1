// package
package queue

import "fmt"

func New[T any](cap int) Queue[T] {
	return &queue[T]{
		front:    nil,
		rear:     nil,
		capacity: cap,
		size:     0,
	}
}

// Interface
type Queue[T any] interface {
	Enqueue(T) error
	Dequeue() (T, error)
	Front() (T, error)
	Rear() (T, error)
	Full() bool
	Empty() bool
}

type queue[T any] struct {
	// Information hiding
	size, capacity int
	front, rear    *element[T]
}

type element[T any] struct {
	next, prev *element[T]
	value      T
}

func (q *queue[T]) Enqueue(v T) error {
	if q.Full() {
		return fmt.Errorf("unable to enqueue element: %w", ErrIsFull)
	}
	new := element[T]{
		value: v,
		next:  nil,
		prev:  q.rear,
	}
	if new.prev != nil {
		new.prev.next = &new
	} else {
		q.front = &new
	}
	q.rear = &new
	return nil
}

func (q *queue[T]) Dequeue() (T, error) {
	var v T
	if q.Empty() {
		return v, fmt.Errorf("unable to dequeue element: %w", ErrIsEmpty)
	}
	dequeuedElement := q.front
	q.front = dequeuedElement.next
	if q.front != nil {
		q.front.prev = nil
	}

	return v, nil
}

func (q queue[T]) Front() (T, error) {
	var v T
	if q.front == nil {
		return v, fmt.Errorf("unable to retrieve front element: %w", ErrIsEmpty)
	}
	return v, nil
}

func (q queue[T]) Rear() (T, error) {
	var v T
	if q.rear == nil {
		return v, fmt.Errorf("unable to retrieve rear element: %w", ErrIsEmpty)
	}
	return v, nil
}

func (q queue[T]) Empty() bool {
	return q.rear == nil
}

func (q queue[T]) Full() bool {
	return q.capacity == q.size
}
