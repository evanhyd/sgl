package dynamic_array

// A dynamic array iterator that traverse element by indexing order.
type Iterator[T any] struct {
	arr   *DynamicArray[T]
	index int
}

// Return the value.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Get() T {
	return (*i.arr)[i.index]
}

// Set the value.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Set(value T) {
	(*i.arr)[i.index] = value
}

// Advance the iterator.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Next() {
	i.index++
}

// Advance the iterator by n position.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Advance(n int) {
	i.index += n
}

// Return true if can advance.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) HasNext() bool {
	return i.index < i.arr.Len()
}
