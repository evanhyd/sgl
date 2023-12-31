package binary_heap

// A max binary heap.
type BinaryHeap[T any, C func(T, T) int] struct {
	slice []T
	cmp   C
}

// Heapify the slice using cmp as predicate.
//
// time complexity: O(n)
//
// space complexity: O(1)
func Make[T any, C func(T, T) int](slice []T, cmp C) BinaryHeap[T, C] {
	heap := BinaryHeap[T, C]{slice, cmp}
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
func (d *BinaryHeap[T, C]) Len() int {
	return len(d.slice)
}

// Return the capacity.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *BinaryHeap[T, C]) Cap() int {
	return cap(d.slice)
}

// Add e to the heap.
//
// time complexity: amortized O(log(len))
//
// space complexity: amortized O(1)
func (b *BinaryHeap[T, C]) Push(e T) {
	b.slice = append(b.slice, e)
	b.fixUp(len(b.slice) - 1)
}

// Remove the top element from the heap.
//
// time complexity: O(log(len))
//
// space complexity: O(1)
func (b *BinaryHeap[T, C]) Pop() {
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
func (b *BinaryHeap[T, C]) Top() T {
	return b.slice[0]
}

// Return an iterator points to the top.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *BinaryHeap[T, C]) Begin() Iterator[T, C] {
	return make[T, C](d)
}

// Fix the heap property start from child upward.
//
// time complexity: O(log(len))
//
// space complexity: O(1)
func (b *BinaryHeap[T, C]) fixUp(child int) {
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
//
// time complexity: O(log(len))
//
// space complexity: O(1)
func (b *BinaryHeap[T, C]) fixDown(root int) {
	for i, j := root, len(b.slice)/2; i < j; {
		child := b.left(i)
		if r := child + 1; r < len(b.slice) && b.cmp(b.slice[r], b.slice[child]) > 0 {
			child = r
		}

		if b.cmp(b.slice[child], b.slice[i]) > 0 {
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
func (b *BinaryHeap[T, C]) parent(i int) int {
	return (i - 1) / 2
}

// Return the left child of i.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (b *BinaryHeap[T, C]) left(i int) int {
	return i*2 + 1
}
