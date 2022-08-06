package queue

/* An implementation of a queue using a slice as the underlying
storage mechanism. Mostly useful to cross-compare with the other queue implementations, since this one is simpler. */

type SliceQueue[I any] []I

func NewSliceQueue[I any](initialCap int) Queue[I] {
	sq := make(SliceQueue[I], 0, initialCap)
	return &sq
}

func (sq *SliceQueue[I]) Enqueue(item I) {
	*sq = append(*sq, item)
}

func (sq *SliceQueue[I]) Dequeue() (item I) {
	item = sq.Peek()
	*sq = (*sq)[1:]
	return
}

func (sq *SliceQueue[I]) Peek() I {
	return (*sq)[0]
}

func (sq SliceQueue[I]) IsEmpty() bool {
	return len(sq) <= 0
}
