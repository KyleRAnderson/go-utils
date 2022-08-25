/*
Implements a heap data structure using Go 1.18 generics.
*/

package heap

/*
Compares a and b, returning a negative value if a is smaller,
a positive on is a is larger, and 0 if they are equal.
*/
type Comparator[T any] func(a, b T) int

/*
Implements a Min heap, that is, a heap where the smallest item
is first, and all items each parent are valued as being larger than or equal to the parent.
*/
type Heap[T any] struct {
	comparator Comparator[T]
	data       []T
}

func New[T any](comp Comparator[T]) *Heap[T] {
	return &Heap[T]{comparator: comp}
}

func (h *Heap[T]) NodeAt(index uint) heapNode[T] {
	return heapNode[T]{h, index}
}

func (h *Heap[T]) Push(item T) {
	h.data = append(h.data, item)
	h.fixUp(h.NodeAt(h.Len() - 1))
}

/*
Retrieves the topmost item in the heap.
Panics if the heap is empty.
*/
func (h *Heap[T]) Peek() T {
	return h.data[0]
}

func (h *Heap[T]) Len() uint {
	return uint(len(h.data))
}

func (h *Heap[T]) Pop() (value T) {
	switch h.Len() {
	case 0:
		panic("heap.Pop on empty heap")
	case 1:
		value = h.data[0]
		h.data = nil
	default:
		value = h.data[0]
		h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
		h.data = h.data[:len(h.data)-1]
		h.fixDown(h.NodeAt(0))
	}
	return
}

/*
Fixes the heap going upward from `node`.
*/
func (h *Heap[T]) fixUp(node heapNode[T]) {
	for {
		parent := node.Parent()
		if parent == nil {
			break
		}
		if !(h.comparator(parent.Value(), node.Value()) <= 0) {
			swapNodes(parent, &node)
		} else {
			break
		}
	}
}

/*
Fixes the heap starting at data[0].
To change the starting point, simply slice the input.
*/
func (h *Heap[T]) fixDown(node heapNode[T]) {
	didSwap := true
	for didSwap {
		didSwap = false
		left, right := node.LeftChild(), node.RightChild()
		if left != nil {
			var toCompare *heapNode[T]
			if right != nil {
				if h.comparator(left.Value(), right.Value()) <= 0 {
					toCompare = left
				} else {
					toCompare = right
				}
			} else {
				toCompare = left
			}
			if !(h.comparator(node.Value(), toCompare.Value()) <= 0) {
				swapNodes(toCompare, &node)
				didSwap = true
			}
		}
	}
}

func leftChildIndex(index uint) uint {
	return 2*index + 1
}

func rightChildIndex(index uint) uint {
	return leftChildIndex(index) + 1
}

/*
Gets the parent index of the child at position `index`.
Assumes that such a parent exists, which isn't true if `index == 0`,
thus the behaviour of this function is undefined when `index == 0`.
*/
func parentIndex(index uint) uint {
	if index == 0 {
		panic("heap.parentIndex: a node at index 0 has no parent")
	}
	return (index - 1) / 2
}

type heapNode[T any] struct {
	heap  *Heap[T]
	index uint
}

func (hn heapNode[T]) Parent() *heapNode[T] {
	if hn.index == 0 {
		return nil
	}
	return &heapNode[T]{hn.heap, parentIndex(hn.index)}
}

/*
Gets the left child of the node, or returns nil if no such child exists.
*/
func (hn heapNode[T]) LeftChild() *heapNode[T] {
	leftIdx := leftChildIndex(uint(hn.index))
	if leftIdx < hn.heap.Len() {
		return &heapNode[T]{hn.heap, leftIdx}
	}
	return nil
}

func (hn heapNode[T]) RightChild() *heapNode[T] {
	rightIdx := rightChildIndex(uint(hn.index))
	if rightIdx < hn.heap.Len() {
		return &heapNode[T]{hn.heap, rightIdx}
	}
	return nil
}

func (hn heapNode[T]) Value() T {
	return hn.heap.data[int(hn.index)]
}

func swapNodes[T any](n1, n2 *heapNode[T]) {
	n1.heap.data[n1.index], n2.heap.data[n2.index] = n2.heap.data[n2.index], n1.heap.data[n1.index]
	n1.index, n2.index = n2.index, n1.index
}

/*
Converts the heap to a slice sorted in the order of the heap.
Postcondition: The heap will be empty.
*/
func (h *Heap[T]) ToSortedSlice() (sorted []T) {
	sorted = make([]T, h.Len())
	for i := range sorted {
		sorted[i] = h.Pop()
	}
	return
}
