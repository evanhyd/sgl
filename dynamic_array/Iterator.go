package dynamic_array

// A dynamic array iterator that traverse element by indexing order.
type Iterator[T any] struct {
	arr   *DynamicArray[T]
	index int
}

// Return the value.
func (i *Iterator[T]) Get() T {
	return (*i.arr)[i.index]
}

// Set the value.
func (i *Iterator[T]) Set(value T) {
	(*i.arr)[i.index] = value
}

// Advance the iterator.
func (i *Iterator[T]) Next() {
	i.index++
}

// Advance the iterator by n position.
func (i *Iterator[T]) Advance(n int) {
	i.index += n
}

// Return true if can advance.
func (i *Iterator[T]) HasNext() bool {
	return i.index < i.arr.Len()
}
