package dynamic_array

// A resizable array.
//
// Equivalent to a slice.
type DynamicArray[T any] []T

func New[T any]() DynamicArray[T] {
	return DynamicArray[T]{}
}

// Return the number of element.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *DynamicArray[T]) Len() int {
	return len(*d)
}

// Return the capacity.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *DynamicArray[T]) Cap() int {
	return cap(*d)
}

// Return a pointer to the first element.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *DynamicArray[T]) Front() *T {
	return &(*d)[0]
}

// Return a pointer to the last element.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *DynamicArray[T]) Back() *T {
	return &(*d)[len(*d)-1]
}

// Append e to the array.
//
// time complexity: amortized O(1)
//
// space complexity: amortized O(1)
func (d *DynamicArray[T]) PushBack(e T) {
	*d = append(*d, e)
}

// Remove the last element.
//
// It also zero it for the GC to clean up.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *DynamicArray[T]) PopBack() {
	last := len(*d) - 1
	var zero T
	(*d)[last] = zero
	*d = (*d)[:last]
}

// Return an iterator points to the first element.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *DynamicArray[T]) Begin() Iterator[T] {
	return Iterator[T]{d, 0}
}

// Return an iterator one pass the last element.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *DynamicArray[T]) End() Iterator[T] {
	return Iterator[T]{d, len(*d)}
}
