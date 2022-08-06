package queue

import "errors"

var errQueueEmpty = errors.New(`queue is empty`)

func ErrQueueEmpty() error { return errQueueEmpty }

type Queue[I any] interface {
	IsEmpty() bool
	Enqueue(I)
	/* Removes and returns the item at the front of the queue.
	panics with ErrQueueEmpty() if the queue is empty. */
	Dequeue() I
	/* Same as Dequeue() but does not remove the item. */
	Peek() I
}
