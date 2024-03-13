package avl_tree

type node[K any, V any] struct {
	left   *node[K, V]
	right  *node[K, V]
	key    K
	value  V
	height int
}

// Get the height of the left and right child nodes.
func (n *node[K, V]) childHeights() (int, int) {
	l, r := -1, -1
	if n.left != nil {
		l = n.left.height
	}
	if n.right != nil {
		r = n.right.height
	}
	return l, r
}

// Update the node height based on child nodes' height.
func (n *node[K, V]) updateHeight() {
	l, r := n.childHeights()
	n.height = max(l, r) + 1
}

// Return the node balance factor.
//
// < 0 if left heavy, > 0 if right heavy, 0 if balance
func (n *node[K, V]) balanceFactor() int {
	l, r := n.childHeights()
	return r - l
}

type AVLTree[K any, V any] struct {
	root *node[K, V]
	len  int
	cmp  func(K, K) int
}

func New[K any, V any](predicate func(K, K) int) AVLTree[K, V] {
	return AVLTree[K, V]{cmp: predicate}
}

// Return the number of element.
func (a *AVLTree[K, V]) Len() int {
	return a.len
}

// Insert a key value pair to the tree.
//
// If the key value pair entry already exists, it updates the value.
func (a *AVLTree[K, V]) Insert(key K, value V) {
	stack := [32]**node[K, V]{}
	pos := -1

	curr := &a.root
	for *curr != nil {
		pos++
		stack[pos] = curr

		if cmp := a.cmp(key, (*curr).key); cmp < 0 {
			curr = &(*curr).left
		} else if cmp > 0 {
			curr = &(*curr).right
		} else {
			break
		}
	}

	if *curr == nil {
		*curr = &node[K, V]{left: nil, right: nil, key: key, value: value, height: 0}
		a.len++
		for ; pos >= 0; pos-- {
			a.balance(stack[pos])
		}
	} else {
		(*curr).value = value
	}
}

// Return the value and the exist indicator.
//
// If the key exists, it returns (value, true).
//
// Otherwise, it returns (zero value, false).
func (a *AVLTree[K, V]) Get(key K) (V, bool) {
	for curr := a.root; curr != nil; {
		if cmp := a.cmp(key, (*curr).key); cmp < 0 {
			curr = curr.left
		} else if cmp > 0 {
			curr = curr.right
		} else {
			return curr.value, true
		}
	}

	var zero V
	return zero, false
}

// Remove key entry from the tree.
func (a *AVLTree[K, V]) Remove(key K) {
	stack := [32]**node[K, V]{}
	pos := -1

	//find the node to remove
	curr := &a.root
	for *curr != nil {
		pos++
		stack[pos] = curr

		if cmp := a.cmp(key, (*curr).key); cmp < 0 {
			curr = &(*curr).left
		} else if cmp > 0 {
			curr = &(*curr).right
		} else {
			break
		}
	}

	//key does not exist
	if *curr == nil {
		return
	}

	//find the replacement
	if (*curr).right == nil {
		*curr = (*curr).left
	} else {
		replace := &(*curr).right
		for ; (*replace).left != nil; replace = &(*replace).left {
			pos++
			stack[pos] = replace
		}
		(*curr).key = (*replace).key
		(*curr).value = (*replace).value
		*replace = (*replace).right
	}

	//delete a leave node
	if *stack[pos] == nil {
		pos--
	}
	a.len--

	for ; pos >= 0; pos-- {
		a.balance(stack[pos])
	}
}

// Return the min key value pair.
func (a *AVLTree[K, V]) Min() (K, V) {
	curr := a.root
	for curr.left != nil {
		curr = curr.left
	}
	return curr.key, curr.value
}

// Return the max key value pair.
func (a *AVLTree[K, V]) Max() (K, V) {
	curr := a.root
	for curr.right != nil {
		curr = curr.right
	}
	return curr.key, curr.value
}

// Balance the subtree rooted at p.
func (a *AVLTree[K, V]) balance(p **node[K, V]) {
	if factor := (*p).balanceFactor(); factor < -1 {
		if (*p).left.balanceFactor() <= 0 {
			a.rightRotate(p)
		} else {
			a.leftRightRotate(p)
		}
	} else if factor > 1 {
		if (*p).right.balanceFactor() >= 0 {
			a.leftRotate(p)
		} else {
			a.rightLeftRotate(p)
		}
	} else {
		(*p).updateHeight()
	}
}

// These functions rotate the tree based on different cases.
//
// It only involves moving pointers around.

func (a *AVLTree[K, V]) leftRotate(p **node[K, V]) {
	//   p
	//   x
	//  / \
	// 0   y
	//    / \
	//   1   z
	//      / \
	//     2   3

	x := *p
	y := x.right
	x.right, y.left = y.left, x
	*p = y
	x.updateHeight()
	y.updateHeight()
}

func (a *AVLTree[K, V]) rightRotate(p **node[K, V]) {
	//       p
	//       x
	//      / \
	//     y   3
	//    / \
	//   z   2
	//  / \
	// 0   1

	x := *p
	y := x.left
	x.left, y.right = y.right, x
	*p = y
	x.updateHeight()
	y.updateHeight()
}

func (a *AVLTree[K, V]) leftRightRotate(p **node[K, V]) {
	//      p
	//      x
	//     / \
	//    y   3
	//   / \
	//  0   z
	//     / \
	//    1   2

	x := *p
	y := x.left
	z := y.right

	y.right, x.left = z.left, z.right
	z.left, z.right = y, x
	*p = z
	x.updateHeight()
	y.updateHeight()
	z.updateHeight()
}

func (a *AVLTree[K, V]) rightLeftRotate(p **node[K, V]) {
	//    p
	//    x
	//   / \
	//  0   y
	//     / \
	//    z   3
	//   / \
	//  1  2

	x := *p
	y := x.right
	z := y.left

	y.left, x.right = z.right, z.left
	z.left, z.right = x, y
	*p = z
	x.updateHeight()
	y.updateHeight()
	z.updateHeight()
}

// Return an iterator points to the first element.
func (a *AVLTree[K, V]) Begin() Iterator[K, V] {
	height := 0
	if a.root != nil {
		height = a.root.height + 1
	}
	iter := Iterator[K, V]{make([]*node[K, V], 0, height)}
	iter.addLeftTree(a.root)
	return iter
}
