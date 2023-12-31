package binary_heap

// A binary heap iterator that traverse from max to min elements.
type Iterator[T any, C func(T, T) int] struct {
	heap  *BinaryHeap[T, C]
	queue BinaryHeap[int, func(int, int) int]
}

// Create a binary heap iterator.
func make[T any, C func(T, T) int](heap *BinaryHeap[T, C]) Iterator[T, C] {
	iter := Iterator[T, C]{
		heap,
		BinaryHeap[int, func(int, int) int]{nil, func(i, j int) int {
			return heap.cmp(heap.slice[i], heap.slice[j])
		}},
	}

	if heap.Len() > 0 {
		iter.queue.slice = []int{0}
	}
	return iter
}

// Return the value.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T, C]) Get() T {
	return i.heap.slice[i.queue.Top()]
}

// Advance the iterator.
//
// time complexity: O(log(len))
//
// space complexity: O(1)
//
// To traverse through the entire heap.
//
// Let n be the number of times we call Next().
// At n-th times, there are n elements in the priority queue.
// It takes at most log(n) <= log(len) to query the next top element.
//
// log(1) + log(2) + ... + log(n-1) + log(n) <= n log(n)
//
// time complexity: O(n log(n))
//
// space complexity: O(n)
func (i *Iterator[T, C]) Next() {
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
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T, C]) HasNext() bool {
	return i.queue.Len() > 0
}
