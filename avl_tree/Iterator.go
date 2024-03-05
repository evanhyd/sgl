package avl_tree

// Iterator that iterate through the tree using in-order traversal.
type Iterator[K any, V any] struct {
	stack []*node[K, V]
}

// Return the key value pair.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[K, V]) Get() (K, V) {
	node := i.stack[len(i.stack)-1]
	return node.key, node.value
}

// Add all the left child rooted at root to the stack.
//
// time complexity: O(log(len))
//
// space complexity: O(log(len))
func (i *Iterator[K, V]) addLeftTree(root *node[K, V]) {
	for ; root != nil; root = root.left {
		i.stack = append(i.stack, root)
	}
}

// Advance the iterator.
//
// time complexity: amortized O(1)
//
// space complexity: amortized O(1)
func (i *Iterator[K, V]) Next() {
	len := len(i.stack) - 1
	node := i.stack[len]
	i.stack = i.stack[:len]
	i.addLeftTree(node.right)
}

// Return true if the iterator is not the end.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (i *Iterator[K, V]) HasNext() bool {
	return len(i.stack) > 0
}
