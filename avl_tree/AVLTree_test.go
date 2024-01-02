package avl_tree

import (
	"fmt"
	"math/rand"
	"testing"
)

// Helper function to calculate the height of a node.
func height(n *node[int]) int {
	if n == nil {
		return 0
	}
	return max(height(n.left), height(n.right)) + 1
}

// Helper function to validate AVL tree properties.
func validateAVLTree(t *testing.T, root *node[int]) bool {
	if root == nil {
		return true
	}

	balanceFactor := height(root.right) - height(root.left)
	if balanceFactor < -1 || balanceFactor > 1 {
		t.Errorf("Balance factor of node %d is %d", root.key, balanceFactor)
		return false
	}

	return validateAVLTree(t, root.left) && validateAVLTree(t, root.right)
}

// Helper function to print the AVL tree horizontally.
func printAVLTree(tree AVLTree[int]) {
	ifs := func(condition bool, a, b string) string {
		if condition {
			return a
		}
		return b
	}

	var printTree func(*node[int], string)
	printTree = func(node *node[int], prefix string) {
		fmt.Println(node.key)

		if node.right != nil {
			fmt.Print(prefix, ifs(node.left != nil, "├─", "└─"))
			printTree(node.right, prefix+ifs(node.left != nil, "| ", "  "))
		}

		if node.left != nil {
			fmt.Print(prefix, "└─")
			printTree(node.left, prefix+"  ")
		}
	}

	printTree(tree.root, "")
}

func TestAVLTree_Insert(t *testing.T) {
	// Test case 1: Inserting elements in ascending order
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	elements := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33}
	for i, e := range elements {
		tree.Insert(e)
		if tree.Len() != i+1 {
			t.Errorf("Len() = %d, want %d", tree.Len(), i+1)
		}
		if !validateAVLTree(t, tree.root) {
			printAVLTree(tree)
			t.Error("tree is not balanced")
		}
	}

	// Test case 2: Inserting elements in descending order
	tree = AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	elements = []int{33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i, e := range elements {
		tree.Insert(e)
		if tree.Len() != i+1 {
			t.Errorf("Len() = %d, want %d", tree.Len(), i+1)
		}
		if !validateAVLTree(t, tree.root) {
			printAVLTree(tree)
			t.Error("tree is not balanced")
		}
	}

	// Test case 3: Inserting random elements
	tree = AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	elements = []int{}
	for i := 0; i < (1 << 13); i++ {
		elements = append(elements, rand.Int())
	}
	for _, e := range elements {
		tree.Insert(e)
		if !validateAVLTree(t, tree.root) {
			printAVLTree(tree)
			fmt.Println(elements)
			t.Error("tree is not balanced")
		}
	}

	// Test case 4: Inserting duplicate elements
	tree = AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	elements = []int{3, 1, 5, 2, 4, 3, 1, 5, 2, 4}
	for _, e := range elements {
		tree.Insert(e)
		if !validateAVLTree(t, tree.root) {
			t.Error("Tree is not balanced")
		}
	}
	if tree.Len() != 5 {
		t.Errorf("Len() = %d, want %d", tree.Len(), 5)
	}
}

func TestAVLTreeIterator(t *testing.T) {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	elements := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33}
	for _, e := range elements {
		tree.Insert(e)
	}

	i := 0
	for iter := tree.Begin(); iter.HasNext(); iter.Next() {
		if *iter.Get() != i {
			t.Errorf("Get() = %d, want %d", *iter.Get(), i)
		}
		i++
	}
}
