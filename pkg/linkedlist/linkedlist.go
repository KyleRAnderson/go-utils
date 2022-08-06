/* Provides a basic linked list. */
package linkedlist

type node[I any] struct {
	Item I
	Next *node[I]
}

/* Linked list that only keeps track of the first item. */
type ForwardList[I any] struct {
	First *node[I]
}

func (fl *ForwardList[I]) Prepend(item I) {
	n := &node[I]{Item: item, Next: fl.First}
	fl.First = n
}

/* Returns and removes the first item in the list.
Panics if the list is empty. Use IsEmpty() to check this prior
to calling this. */
func (fl *ForwardList[I]) PopFirst() (item I) {
	item = fl.First.Item
	fl.First = fl.First.Next
	return
}

func (fl *ForwardList[I]) IsEmpty() bool {
	return fl.First == nil
}

/* Linked list that keeps track of the first and last item. */
type LinkedList[I any] struct {
	ForwardList[I]
	Last *node[I]
}

func (ll *LinkedList[I]) Prepend(item I) {
	if ll.IsEmpty() {
		defer func() {
			ll.Last = ll.First
		}()
	}
	ll.ForwardList.Prepend(item)
}

func (ll *LinkedList[I]) Append(item I) {
	if ll.IsEmpty() {
		defer func() {
			ll.First = ll.Last
		}()
	}
	n := &node[I]{Item: item}
	if ll.Last != nil {
		ll.Last.Next = n
	}
	ll.Last = n
}
