package singlylinkedlist

// A singly linked list iterator that traverse elements by chaining order.
type Iterator[T any] struct {
	n **node[T]
}

// Return the pointer to the node value.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Get() *T {
	return &(*i.n).value
}

// Advance the iterator.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Next() {
	i.n = &(*i.n).next
}

// Return true if can advance.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) HasNext() bool {
	return *i.n != nil
}
