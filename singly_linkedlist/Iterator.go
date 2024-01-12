package singly_linkedlist

// A singly linked list iterator that traverse elements by chaining order.
type Iterator[T any] struct {
	n **node[T]
}

// Return the value.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Get() T {
	return (*i.n).value
}

// Set the value.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Set(value T) {
	(*i.n).value = value
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
