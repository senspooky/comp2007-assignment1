package queue

import "errors"

var (
	ErrIsEmpty = errors.New("queue is empty")
	ErrIsFull  = errors.New("queue is full")
)
