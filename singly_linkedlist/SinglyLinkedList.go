package singlylinkedlist

type node[T any] struct {
	value T
	next  *node[T]
}

// A singly linked list that supports traversing forward only.
type SinglyLinkedList[T any] struct {
	head *node[T]
	len  int
}

// Return the number of element.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (d *SinglyLinkedList[T]) Len() int {
	return d.len
}

// Prepend e to the list.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (s *SinglyLinkedList[T]) PushFront(e T) {
	s.head = &node[T]{e, s.head}
	s.len++
}

// Remove the first element.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (s *SinglyLinkedList[T]) PopFront() {
	s.head = s.head.next
	s.len--
}

// Return an iterator points to the first element.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (s *SinglyLinkedList[T]) Begin() Iterator[T] {
	return Iterator[T]{&s.head}
}

// Insert e before i.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (s *SinglyLinkedList[T]) Insert(i Iterator[T], e T) {
	*i.n = &node[T]{e, *i.n}
	s.len++
}

// Remove the element at i.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (s *SinglyLinkedList[T]) Remove(i Iterator[T]) {
	*i.n = (*i.n).next
	s.len--
}
