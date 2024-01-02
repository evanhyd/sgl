package avl_tree

import (
	"fmt"
	"math/bits"
)

type node[T any] struct {
	key    T
	left   *node[T]
	right  *node[T]
	height int
}

// Get the height of the left and right child nodes.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (n *node[T]) childHeights() (int, int) {
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
//
// time complexity: O(1)
//
// space complexity: O(1)
func (n *node[T]) updateHeight() {
	l, r := n.childHeights()
	n.height = max(l, r) + 1
}

// Return the node balance factor.
//
// < 0 if left heavy, > 0 if right heavy, 0 if balance
//
// time complexity: O(1)
//
// space complexity: O(1)
func (n *node[T]) balanceFactor() int {
	l, r := n.childHeights()
	return r - l
}

type AVLTree[T any] struct {
	root *node[T]
	len  int
	Cmp  func(T, T) int
}

// Return the number of element.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (a *AVLTree[T]) Len() int {
	return a.len
}

// Insert e to the tree.
//
// time complexity: O(log(len))
//
// space complexity: O(1)
func (a *AVLTree[T]) Insert(e T) {
	stack := make([]**node[T], 0, bits.Len(uint(a.len))+1)

	curr := &a.root
	for *curr != nil {
		stack = append(stack, curr)
		if cmp := a.Cmp(e, (*curr).key); cmp < 0 {
			curr = &(*curr).left
		} else if cmp > 0 {
			curr = &(*curr).right
		} else {
			return
		}
	}
	*curr = &node[T]{e, nil, nil, 0}
	a.len++

	for i := len(stack) - 1; i >= 0; i-- {
		a.balance(stack[i])
	}
}

// Balance the subtree rooted at p.
//
// time complexity: O(1)
//
// space complexity: O(1)
func (a *AVLTree[T]) balance(p **node[T]) {
	if factor := (*p).balanceFactor(); factor < -1 {
		if (*p).left.balanceFactor() < 0 {
			a.rightRotate(p)
		} else {
			a.leftRightRotate(p)
		}
	} else if factor > 1 {
		if (*p).right.balanceFactor() > 0 {
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
//
// time complexity: O(1)
//
// space complexity: O(1)

func (a *AVLTree[T]) leftRotate(p **node[T]) {
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

func (a *AVLTree[T]) rightRotate(p **node[T]) {
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

func (a *AVLTree[T]) leftRightRotate(p **node[T]) {
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

func (a *AVLTree[T]) rightLeftRotate(p **node[T]) {
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
//
// time complexity: O(1)
//
// space complexity: O(1)
func (a *AVLTree[T]) Begin() Iterator[T] {
	height := 0
	if a.root != nil {
		height = a.root.height + 1
	}
	iter := Iterator[T]{make([]*node[T], 0, height)}
	iter.addLeftTree(a.root)
	fmt.Println(iter.stack)
	return iter
}
