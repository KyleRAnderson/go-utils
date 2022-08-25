package heap

import (
	"math"
	"math/rand"
	"testing"

	"fmt"
)

/*
Tests the heap implementation against the standard library one.
*/
func TestHeapPresetCases(t *testing.T) {
	t.Run(`single push and pop`, func(t *testing.T) {
		runCase := func(value int) {
			t.Run(fmt.Sprint(value), func(t *testing.T) {
				h := New(func(a, b int) int { return a - b })
				h.Push(value)
				if received := h.Pop(); received != value {
					t.Errorf(`expected: %v, received: %v`, value, received)
				}
			})
		}
		for _, value := range []int{0, 1, 5} {
			runCase(value)
		}

		r := rand.New(rand.NewSource(3))
		for i := 0; i < 50; i++ {
			runCase(r.Int())
		}
	})

	t.Run(`pop on empty`, func(t *testing.T) {
		/* pop on empty should panic */
		defer func() {
			if err := recover(); err == nil {
				t.Error(`expected to panic but did not`)
			}
		}()
		New[int](nil).Pop()
	})

	t.Run(`ToSortedSlice`, func(t *testing.T) {
		caseNo := 0
		runCase := func(input []int) {
			caseNo++
			t.Run(fmt.Sprint(caseNo), func(t *testing.T) {
				h := New(func(a, b int) int { return a - b })
				expectedCounts := make(map[int]uint, len(input))
				for _, value := range input {
					h.Push(value)
					if _, ok := expectedCounts[value]; !ok {
						expectedCounts[value] = 0
					}
					expectedCounts[value]++
				}
				received := h.ToSortedSlice()
				t.Logf(`received: %v`, received)
				if len(input) != len(received) {
					t.Errorf(`expected len: %d, received len: %d`, len(input), len(received))
				} else {
					previous := math.MinInt
					for i, current := range received {
						if !(previous <= current) {
							t.Errorf(`bad ordering at index %d, previous: %v, current: %v`, i, previous, current)
						} else if count, ok := expectedCounts[current]; !ok {
							t.Errorf(`received unexpected value %v at index %d`, current, i)
						} else if count <= 0 {
							t.Errorf(`received duplicate occurrence of %v at index %d`, current, i)
						} else {
							expectedCounts[current]--
						}
						previous = current
					}
				}
				if l := h.Len(); l != 0 {
					t.Error("expected heap to be empty but reports len of ", l)
				}
			})
		}
		for _, input := range [][]int{
			{},
			{-1, 0, 1},
			{1, 2, 3},
			{5, 1034, 1_000, 0x1afd9},
			{7, 4, 9, 5, 4, 6, 5, 7},
		} {
			runCase(input)
		}
		r := rand.New(rand.NewSource(5))
		for _, size := range []uint{8, 10, 50, 1_015} {
			testCase := make([]int, size)
			for i := range testCase {
				testCase[i] = r.Int()
			}
			runCase(testCase)
		}
	})
}
