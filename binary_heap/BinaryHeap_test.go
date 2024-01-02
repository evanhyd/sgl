package binary_heap

import (
	"fmt"
	"testing"
)

func checkHeapProperty[T any, C func(T, T) int](heap BinaryHeap[T], t *testing.T) {
	for i := 0; i < heap.Len(); i++ {
		left := heap.left(i)
		right := left + 1

		if left < heap.Len() && heap.Cmp(heap.slice[left], heap.slice[i]) > 0 {
			t.Errorf("Heap property violated at index %d and %d", i, left)
		}

		if right < heap.Len() && heap.Cmp(heap.slice[right], heap.slice[i]) > 0 {
			t.Errorf("Heap property violated at index %d and %d", i, right)
		}
	}
}

func TestHeapify(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	cmp := func(a, b int) int { return a - b }
	heap := Heapify(slice, cmp)

	checkHeapProperty(heap, t)
}

func TestBinaryHeap_Len(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	cmp := func(a, b int) int { return a - b }
	heap := Heapify(slice, cmp)

	// Check if Len returns the correct number of elements
	expectedLen := len(slice)
	actualLen := heap.Len()
	if actualLen != expectedLen {
		t.Errorf("Len() = %v, want %v", actualLen, expectedLen)
	}
}

func TestBinaryHeap_Cap(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	cmp := func(a, b int) int { return a - b }
	heap := Heapify(slice, cmp)

	// Check if Cap returns the correct capacity
	expectedCap := cap(slice)
	actualCap := heap.Cap()
	if actualCap != expectedCap {
		t.Errorf("Cap() = %v, want %v", actualCap, expectedCap)
	}
}

func TestBinaryHeap_Push(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	cmp := func(a, b int) int { return a - b }
	heap := Heapify(slice, cmp)

	// Push a new element to the heap
	newElement := 7
	heap.Push(newElement)

	// Check if the top element is correct
	expectedTop := 9
	actualTop := heap.Top()
	if actualTop != expectedTop {
		t.Errorf("Top() = %v, want %v", actualTop, expectedTop)
	}

	checkHeapProperty(heap, t)
}

func TestBinaryHeap_Pop(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	cmp := func(a, b int) int { return a - b }
	heap := Heapify(slice, cmp)

	// Pop the top element from the heap
	heap.Pop()

	// Check if the top element is correct
	expectedTop := 6
	actualTop := heap.Top()
	if actualTop != expectedTop {
		t.Errorf("Top() = %v, want %v", actualTop, expectedTop)
	}

	checkHeapProperty(heap, t)
}

func TestBinaryHeap_Top(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	cmp := func(a, b int) int { return a - b }
	heap := Heapify(slice, cmp)

	// Check if Top returns the correct top element
	expectedTop := 9
	actualTop := heap.Top()
	if actualTop != expectedTop {
		t.Errorf("Top() = %v, want %v", actualTop, expectedTop)
	}

	checkHeapProperty(heap, t)
}

func TestBinaryHeap_Begin(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	cmp := func(a, b int) int { return a - b }
	heap := Heapify(slice, cmp)

	// Check if Begin() points to the top
	expectedTop := 9
	iter := heap.Begin()
	if iter.Get() != expectedTop {
		t.Errorf("Top() = %v, want %v", iter.Get(), expectedTop)
	}

	checkHeapProperty(heap, t)
}

func TestIterator_Top(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 99, 2, 6, 5, 3, 5}
	cmp := func(a, b int) int { return a - b }
	heap := Heapify(slice, cmp)

	// Check if Begin() points to the top
	expectedTop := 99
	iter := heap.Begin()
	if iter.Get() != expectedTop {
		t.Errorf("Get() = %v, want %v", iter.Get(), expectedTop)
	}

	checkHeapProperty(heap, t)
}

func TestIterator_Next(t *testing.T) {
	slice := []int{3, 1, 4, 88, 5, 99, 2, 6, 5, 3, 5}
	cmp := func(a, b int) int { return a - b }
	heap := Heapify(slice, cmp)

	expectedTop := 99
	iter := heap.Begin()
	if iter.Get() != expectedTop {
		t.Errorf("Get() = %v, want  %v", iter.Get(), expectedTop)
	}

	iter.Next()
	expectedTop = 88
	if iter.Get() != expectedTop {
		t.Errorf("Get() = %v, want  %v", iter.Get(), expectedTop)
	}

	checkHeapProperty(heap, t)
}

func TestIterator_HasNext(t *testing.T) {
	slice := []int{1, 3, 2}
	cmp := func(a, b int) int { return a - b }
	heap := Heapify(slice, cmp)

	iter := heap.Begin()
	if !iter.HasNext() {
		t.Errorf("HasNext() = false, want true")
	}
	iter.Next()

	if !iter.HasNext() {
		t.Errorf("HasNext() = false, want true")
	}
	iter.Next()

	if !iter.HasNext() {
		t.Errorf("HasNext() = false, want true")
	}
	iter.Next()

	if iter.HasNext() {
		t.Errorf("HasNext() = true, want false")
	}

	checkHeapProperty(heap, t)
}

func ExampleMake() {
	// Create a max heap from a slice
	slice := []int{4, 2, 7, 1, 9, 5}
	maxHeap := Heapify(slice, func(i, j int) int { return i - j })
	fmt.Println(maxHeap.slice)
	// Output: [9 4 7 1 2 5]
}

func ExampleBinaryHeap_Len() {
	// Get the length of the heap
	slice := []int{4, 2, 7, 1, 9, 5}
	maxHeap := Heapify(slice, func(i, j int) int { return i - j })
	length := maxHeap.Len()
	fmt.Println(length)
	// Output: 6
}

func ExampleBinaryHeap_Cap() {
	// Get the capacity of the heap
	slice := []int{4, 2, 7, 1, 9, 5}
	maxHeap := Heapify(slice, func(i, j int) int { return i - j })
	capacity := maxHeap.Cap()
	fmt.Println(capacity)
	// Output: 6
}

func ExampleBinaryHeap_Push() {
	// Push an element into the heap
	slice := []int{4, 2, 7, 1, 9, 5}
	maxHeap := Heapify(slice, func(i, j int) int { return i - j })
	maxHeap.Push(8)
	fmt.Println(maxHeap.slice)
	// Output: [9 4 8 1 2 5 7]
}

func ExampleBinaryHeap_Pop() {
	// Pop the top element from the heap
	slice := []int{4, 2, 7, 1, 9, 5}
	maxHeap := Heapify(slice, func(i, j int) int { return i - j })
	maxHeap.Pop()
	fmt.Println(maxHeap.slice)
	// Output: [7 4 5 1 2]
}

func ExampleBinaryHeap_Top() {
	// Get the top element of the heap
	slice := []int{4, 2, 7, 1, 9, 5}
	maxHeap := Heapify(slice, func(i, j int) int { return i - j })
	topElement := maxHeap.Top()
	fmt.Println(topElement)
	// Output: 9
}

func ExampleBinaryHeap_Begin() {
	// Get the top element of the heap
	slice := []int{4, 2, 7, 1, 9, 5}
	maxHeap := Heapify(slice, func(i, j int) int { return i - j })
	topElement := maxHeap.Begin()
	fmt.Println(topElement.Get())
	// Output: 9
}

func ExampleIterator() {
	// Iterate through the heap
	slice := []int{4, 2, 7, 1, 9, 5}
	maxHeap := Heapify(slice, func(i, j int) int { return i - j })
	for iter := maxHeap.Begin(); iter.HasNext(); iter.Next() {
		fmt.Println(iter.Get())
	}
	// Output:
	// 9
	// 7
	// 5
	// 4
	// 2
	// 1
}
