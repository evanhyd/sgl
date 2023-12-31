package binary_heap

// A max binary heap.
type BinaryHeap[T any] struct {
	slice []T
	Cmp   func(T, T) int
}

// Heapify the slice using cmp as predicate.
//
// time complexity: O(n)
//
// space complexity: O(1)
func Heapify[T any](slice []T, cmp func(T, T) int) BinaryHeap[T] {
	heap := BinaryHeap[T]{slice, cmp}
	for i := len(heap.slice)/2 - 1; i >= 0; i-- {
		heap.fixDown(i)
	}
	return heap
}

// Return the number of element.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *BinaryHeap[T]) Len() int {
	return len(d.slice)
}

// Return the capacity.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *BinaryHeap[T]) Cap() int {
	return cap(d.slice)
}

// Add e to the heap.
//
// time complexity: amortized O(log(len))
//
// space complexity: amortized O(1)
func (b *BinaryHeap[T]) Push(e T) {
	b.slice = append(b.slice, e)
	b.fixUp(len(b.slice) - 1)
}

// Remove the top element from the heap.
//
// time complexity: O(log(len))
//
// space complexity: O(1)
func (b *BinaryHeap[T]) Pop() {
	last := len(b.slice) - 1
	b.slice[0] = b.slice[last]
	var zero T
	b.slice[last] = zero
	b.slice = b.slice[:last]
	b.fixDown(0)
}

// Return the top of the heap.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (b *BinaryHeap[T]) Top() T {
	return b.slice[0]
}

// Return an iterator points to the top.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *BinaryHeap[T]) Begin() Iterator[T] {
	return make[T](d)
}

// Fix the heap property start from child upward.
//
// time complexity: O(log(len))
//
// space complexity: O(1)
func (b *BinaryHeap[T]) fixUp(child int) {
	for i := child; i > 0; {
		p := b.parent(i)
		if b.Cmp(b.slice[i], b.slice[p]) > 0 {
			b.slice[i], b.slice[p] = b.slice[p], b.slice[i]
			i = p
		} else {
			break
		}
	}
}

// Fix the heap property rooted at root downward.
//
// time complexity: O(log(len))
//
// space complexity: O(1)
func (b *BinaryHeap[T]) fixDown(root int) {
	for i, j := root, len(b.slice)/2; i < j; {
		child := b.left(i)
		if r := child + 1; r < len(b.slice) && b.Cmp(b.slice[r], b.slice[child]) > 0 {
			child = r
		}

		if b.Cmp(b.slice[child], b.slice[i]) > 0 {
			b.slice[i], b.slice[child] = b.slice[child], b.slice[i]
			i = child
		} else {
			break
		}
	}
}

// Return the parent of i.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (b *BinaryHeap[T]) parent(i int) int {
	return (i - 1) / 2
}

// Return the left child of i.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (b *BinaryHeap[T]) left(i int) int {
	return i*2 + 1
}
