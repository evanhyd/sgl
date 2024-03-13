package binary_heap

import "github.com/evanhyd/sgl/adt"

// A max binary heap.
//
// interface: PriorityQueue
type BinaryHeap[T any] struct {
	slice []T
	cmp   func(T, T) int
}

var _ adt.PriorityQueue[int] = &BinaryHeap[int]{}

func New[T any](predicate func(T, T) int) BinaryHeap[T] {
	return BinaryHeap[T]{cmp: predicate}
}

// Heapify the slice using cmp as predicate.
func Heapify[T any](slice []T, cmp func(T, T) int) BinaryHeap[T] {
	heap := BinaryHeap[T]{slice, cmp}
	for i := len(heap.slice)/2 - 1; i >= 0; i-- {
		heap.fixDown(i)
	}
	return heap
}

// Return the number of element.
func (d *BinaryHeap[T]) Len() int {
	return len(d.slice)
}

// Return the capacity.
func (d *BinaryHeap[T]) Cap() int {
	return cap(d.slice)
}

// Add e to the heap.
func (b *BinaryHeap[T]) Push(e T) {
	b.slice = append(b.slice, e)
	b.fixUp(len(b.slice) - 1)
}

// Remove the top element from the heap.
func (b *BinaryHeap[T]) Pop() {
	last := len(b.slice) - 1
	b.slice[0] = b.slice[last]
	var zero T
	b.slice[last] = zero
	b.slice = b.slice[:last]
	b.fixDown(0)
}

// Return the top of the heap.
func (b *BinaryHeap[T]) Top() T {
	return b.slice[0]
}

// Return an iterator points to the top.
func (d *BinaryHeap[T]) Begin() Iterator[T] {
	return newIterator(d)
}

// Fix the heap property start from child upward.
func (b *BinaryHeap[T]) fixUp(child int) {
	for i := child; i > 0; {
		p := b.parent(i)
		if b.cmp(b.slice[i], b.slice[p]) > 0 {
			b.slice[i], b.slice[p] = b.slice[p], b.slice[i]
			i = p
		} else {
			break
		}
	}
}

// Fix the heap property rooted at root downward.
func (b *BinaryHeap[T]) fixDown(root int) {
	for end := len(b.slice) / 2; root < end; {
		child := b.left(root)

		if r := child + 1; r < len(b.slice) {
			if b.cmp(b.slice[r], b.slice[child]) > 0 {
				child = r
			}
		}

		if b.cmp(b.slice[child], b.slice[root]) > 0 {
			b.slice[root], b.slice[child] = b.slice[child], b.slice[root]
			root = child
		} else {
			break
		}
	}
}

// Return the parent of i.
func (b *BinaryHeap[T]) parent(i int) int {
	return (i - 1) / 2
}

// Return the left child of i.
func (b *BinaryHeap[T]) left(i int) int {
	return i*2 + 1
}
