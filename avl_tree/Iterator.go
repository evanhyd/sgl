package avl_tree

// Iterator that iterate through the tree using in-order traversal.
type Iterator[T any] struct {
	stack []*node[T]
}

// Return the pointer to the value.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) Get() *T {
	return &i.stack[len(i.stack)-1].key
}

// Add all the left child rooted at root to the stack.
//
// time complexity: O(log(len))
//
// space complexity: O(log(len))
func (i *Iterator[T]) addLeftTree(root *node[T]) {
	for ; root != nil; root = root.left {
		i.stack = append(i.stack, root)
	}
}

// Advance the iterator.
//
// time complexity: amortized O(1)
//
// space complexity: amortized O(1)
func (i *Iterator[T]) Next() {
	last := i.stack[len(i.stack)-1]
	i.stack = i.stack[:len(i.stack)-1]
	i.addLeftTree(last.right)
}

// Return true if can advance.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[T]) HasNext() bool {
	return len(i.stack) > 0
}
