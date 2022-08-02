package queue

import "gitlab.com/kyle_anderson/go-utils/pkg/linkedlist"

/* Linked list implementation of a queue. */

func NewLinkedListQueue[I any]() Queue[I] {
	return &linkedListQueue[I]{}
}

type linkedListQueue[I any] struct {
	linkedlist.LinkedList[I]
}

func (llq *linkedListQueue[I]) Enqueue(item I) {
	/* The end of the linked list is the end of the queue. */
	llq.LinkedList.Append(item)
}

func (llq *linkedListQueue[I]) Dequeue() I {
	/* The start of the linked list is the start of the queue. */
	return llq.LinkedList.PopFirst()
}

func (llq *linkedListQueue[I]) Peek() I {
	return llq.LinkedList.First.Item
}
