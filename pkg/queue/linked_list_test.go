package queue

import (
	"math/rand"
	"testing"
)

func TestComparisonToSliceQueue(t *testing.T) {
	/* The idea of this test is to compare the more complicated linked list queue implementation
	with a simple but inefficient queue implementation that uses slices. */
	t.Run(`sanity check`, func(t *testing.T) {

		type operation[I any] func(q Queue[I]) I

		var sq Queue[int] = NewSliceQueue[int](0)
		var llq Queue[int] = NewLinkedListQueue[int]()

		for i, op := range []operation[int]{
			/* Return values for enqueues will always match, we don't check anything there. */
			func(q Queue[int]) int { q.Enqueue(1); return 0 },
			func(q Queue[int]) int { q.Enqueue(2); return 0 },
			func(q Queue[int]) int { q.Enqueue(3); return 0 },
			func(q Queue[int]) int { q.Enqueue(4); return 0 },
			func(q Queue[int]) int { return q.Dequeue() },
			func(q Queue[int]) int { return q.Dequeue() },
			func(q Queue[int]) int { return q.Dequeue() },
			func(q Queue[int]) int { return q.Dequeue() },
		} {
			if sqRes, llqRes := op(sq), op(llq); sqRes != llqRes {
				t.Errorf(`mismatch on operation results at index %d, sq: %v, llq: %v`, i, sqRes, llqRes)
			}
			if sqEmpty, llqEmpty := sq.IsEmpty(), llq.IsEmpty(); sqEmpty != llqEmpty {
				t.Errorf(`mismatch on IsEmpty, llq.IsEmpty(): %v, sq.IsEmpty(): %v`, llqEmpty, sqEmpty)
			}
			if t.Failed() {
				t.FailNow()
			}
		}
	})

	t.Run(`generated cases`, func(t *testing.T) {
		r := rand.New(rand.NewSource(3))
		var source [0x10_000]int
		for i := range source {
			source[i] = i
		}
		r.Shuffle(len(source), func(i, j int) { source[i], source[j] = source[j], source[i] })

		var sq Queue[int] = NewSliceQueue[int](0)
		var llq Queue[int] = NewLinkedListQueue[int]()

		checkDequeuePanic := func(q Queue[int]) (success bool) {
			defer func() {
				err := recover()
				success = err != nil
			}()
			q.Dequeue()
			return
		}

		for i := 0; i < len(source) && !t.Failed(); {
			t.Log(`index `, i, ` value `, source[i])
			switch r.Intn(2) { /* Choose an operation randomly. */
			case 0:
				t.Log(`enqueue `, source[i])
				sq.Enqueue(source[i])
				llq.Enqueue(source[i])
				i++
			case 1:
				t.Log(`dequeue`)
				if sq.IsEmpty() {
					if !checkDequeuePanic(sq) {
						t.Error(`sq did not panic when empty and dequeued`)
					}
					if !checkDequeuePanic(llq) {
						t.Error(`llq did not panic when empty and dequeued`)
					}
				} else {
					sqVal := sq.Dequeue()
					llqVal := llq.Dequeue()
					if llqVal != sqVal {
						t.Errorf(`dequeue mismatch, sqVal: %v, llqVal: %v`, sqVal, llqVal)
					}
				}
			}
			if sqEmpty, llqEmpty := sq.IsEmpty(), llq.IsEmpty(); sqEmpty != llqEmpty {
				t.Errorf(`IsEmpty() mismatch, sq: %v, llq: %v`, sqEmpty, llqEmpty)
			}
		}
	})
}
