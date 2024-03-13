package binary_heap

// A binary heap iterator that traverse from max to min elements.
type Iterator[T any] struct {
	heap  *BinaryHeap[T]
	queue BinaryHeap[int]
}

// Create a binary heap iterator.
func newIterator[T any, C func(T, T) int](heap *BinaryHeap[T]) Iterator[T] {
	iter := Iterator[T]{
		heap,
		BinaryHeap[int]{cmp: func(i, j int) int {
			return heap.cmp(heap.slice[i], heap.slice[j])
		}},
	}

	if heap.Len() > 0 {
		iter.queue.Push(0)
	}
	return iter
}

// Return the value.
func (i *Iterator[T]) Get() T {
	return i.heap.slice[i.queue.Top()]
}

// Advance the iterator.
func (i *Iterator[T]) Next() {
	top := i.queue.Top()
	i.queue.Pop()

	if l := i.heap.left(top); l < i.heap.Len() {
		i.queue.Push(l)
		if r := l + 1; r < i.heap.Len() {
			i.queue.Push(r)
		}
	}
}

// Return true if can advance.
func (i *Iterator[T]) HasNext() bool {
	return i.queue.Len() > 0
}
