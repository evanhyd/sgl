package dynamic_array

// A dynamic array iterator that traverse element by indexing order.
type Iterator[T any] struct {
	arr   *DynamicArray[T]
	index int
}

// Return the pointer to the node value.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Get() *T {
	return &(*i.arr)[i.index]
}

// Advance the iterator.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Next() {
	i.index++
}

// Return true if can advance.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) HasNext() bool {
	return i.index < i.arr.Len()
}
