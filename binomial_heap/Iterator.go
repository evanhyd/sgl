package binomial_heap

import "github.com/evanhyd/sgl/binary_heap"

// A binomial heap iterator that traverse from max to min elements.
type Iterator[T any] struct {
	heap  *BinomialHeap[T]
	queue binary_heap.BinaryHeap[*flagTree[T]]
}

// Create a binomial heap iterator.
func newIterator[T any, C func(T, T) int](heap *BinomialHeap[T]) Iterator[T] {
	iter := Iterator[T]{
		heap,
		binary_heap.BinaryHeap[*flagTree[T]]{Cmp: func(l, r *flagTree[T]) int {
			return heap.Cmp(l.key, r.key)
		}},
	}

	for _, tree := range heap.trees {
		iter.queue.Push(tree)
	}
	return iter
}

// Return the value.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Get() T {
	return i.queue.Top().key
}

// Advance the iterator.
//
// time complexity: O(log(len))
//
// space complexity: O(log(len))
func (i *Iterator[T]) Next() {
	// top := i.queue.Top()
	// i.queue.Pop()
}

// Return true if can advance.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) HasNext() bool {
	return i.queue.Len() > 0
}
