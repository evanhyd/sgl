package avl_tree

import (
	"fmt"
	"math/rand"
	"testing"
)

// Helper function to validateTree AVL tree properties.
func validateTree(t *testing.T, tree AVLTree[int]) bool {
	var height func(*node[int]) int
	height = func(n *node[int]) int {
		if n == nil {
			return 0
		}
		return max(height(n.left), height(n.right)) + 1
	}

	var validateNode func(*testing.T, *node[int]) bool
	validateNode = func(t *testing.T, root *node[int]) bool {
		if root == nil {
			return true
		}

		balanceFactor := height(root.right) - height(root.left)
		if balanceFactor < -1 || balanceFactor > 1 {
			t.Errorf("Balance factor of node %d is %d", root.key, balanceFactor)
			return false
		}

		return validateNode(t, root.left) && validateNode(t, root.right)
	}

	return validateNode(t, tree.root)
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
		if !validateTree(t, tree) {
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
		if !validateTree(t, tree) {
			printAVLTree(tree)
			t.Error("tree is not balanced")
		}
	}

	// Test case 3: Inserting duplicate elements
	tree = AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	elements = []int{3, 1, 5, 2, 4, 3, 1, 5, 2, 4}
	for _, e := range elements {
		tree.Insert(e)
		if !validateTree(t, tree) {
			t.Error("Tree is not balanced")
		}
	}
	if tree.Len() != 5 {
		t.Errorf("Len() = %d, want %d", tree.Len(), 5)
	}

	// Test case 4: Inserting random elements
	tree = AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	elements = rand.Perm(100)
	for _, e := range elements {
		tree.Insert(e)
		if !validateTree(t, tree) {
			t.Error("tree is not balanced")
		}
	}

	elements = rand.Perm(1 << 20)
	for _, e := range elements {
		tree.Insert(e)
	}
	if !validateTree(t, tree) {
		t.Error("tree is not balanced")
	}
}

func TestAVLTree_Contain(t *testing.T) {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	for i := 0; i < 100; i += 2 {
		tree.Insert(i)
	}
	for i := 0; i < 100; i++ {
		shoulContain := (i%2 == 0)
		if contain := tree.Contain(i); contain != shoulContain {
			t.Errorf("Contain(%v) = %v, want %v", i, contain, shoulContain)
		}
	}
}

func TestAVLTree_Remove(t *testing.T) {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}

	elements := rand.Perm(300)
	// elements := []int{7, 6, 3, 9, 0, 1, 4, 8, 2, 5}
	for _, e := range elements {
		tree.Insert(e)
	}

	for _, e := range elements {
		if contain := tree.Contain(e); !contain {
			t.Errorf("Contain(%v) = %v, want %v", e, contain, true)
			fmt.Println(elements)
			return
		}

		tree.Remove(e)

		if contain := tree.Contain(e); contain {
			t.Errorf("Contain(%v) = %v, want %v", e, contain, false)
			fmt.Println(elements)
			return
		}

		if !validateTree(t, tree) {
			t.Error("Tree is not balanced")
			printAVLTree(tree)
			fmt.Println(elements)
			return
		}
	}
}

func TestAVLTree_Min(t *testing.T) {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	for i := 0; i < 100; i++ {
		tree.Insert(i)
	}
	if m := tree.Min(); m != 0 {
		t.Errorf("Min() = %d, want %d", m, 0)
	}
}

func TestAVLTree_Max(t *testing.T) {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	for i := 0; i < 100; i++ {
		tree.Insert(i)
	}
	if m := tree.Max(); m != 99 {
		t.Errorf("Max() = %d, want %d", m, 99)
	}
}

func TestIterator(t *testing.T) {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	elements := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33}
	for _, e := range elements {
		tree.Insert(e)
	}

	i := 0
	for iter := tree.Begin(); iter.HasNext(); iter.Next() {
		if iter.Get() != i {
			t.Errorf("Get() = %d, want %d", iter.Get(), i)
		}
		i++
	}
}

func BenchmarkAVLTree_Insert_Small(b *testing.B) {
	// int64
	// iterative slice: BenchmarkAVLTree_Insert_Small-16    	 3736772	       337.5 ns/op	     202 B/op	       2 allocs/op
	// iterative array: BenchmarkAVLTree_Insert_Small-16    	 5551730	       236.3 ns/op	      32 B/op	       1 allocs/op
	// recursion:       BenchmarkAVLTree_Insert_Small-16    	 5480019	       230.7 ns/op	      32 B/op	       1 allocs/op
	//
	// Allocating slice brings a huge performance penalty.
	// Recursion is consistently faster than the iterative array version.

	tree := AVLTree[int64]{Cmp: func(a, b int64) int { return int(a - b) }}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(int64(i))
	}
}

func BenchmarkAVLTree_Insert_Big(b *testing.B) {
	// [20]int64
	// iterative slice: BenchmarkAVLTree_Insert_Big-16    	 2328404	       492.9 ns/op	     358 B/op	       2 allocs/op
	// iterative array: BenchmarkAVLTree_Insert_Big-16    	 2671484	       430.6 ns/op	     192 B/op	       1 allocs/op
	// recursion:       BenchmarkAVLTree_Insert_Big-16    	 2213972	       543.8 ns/op	     192 B/op	       1 allocs/op

	type Large [20]int64
	tree := AVLTree[Large]{Cmp: func(a, b Large) int { return int(a[0] - b[0]) }}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(Large{int64(i)})
	}
}

func ExampleAVLTree_Insert() {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	fmt.Println(tree.Len())
	// Output:
	// 3
}

func ExampleAVLTree_Contain() {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	fmt.Println(tree.Contain(5))
	fmt.Println(tree.Contain(6))
	// Output:
	// true
	// false
}

func ExampleAVLTree_Remove() {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	fmt.Println(tree.Contain(5))
	tree.Remove(5)
	fmt.Println(tree.Contain(5))
	// Output:
	// true
	// false
}

func ExampleAVLTree_Min() {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	fmt.Println(tree.Min())
	// Output:
	// 5
}

func ExampleAVLTree_Max() {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	fmt.Println(tree.Max())
	// Output:
	// 15
}

func ExampleIterator() {
	tree := AVLTree[int]{Cmp: func(a, b int) int { return a - b }}
	tree.Insert(10)
	tree.Insert(-436)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(12)
	tree.Insert(8)
	tree.Insert(6)

	for iter := tree.Begin(); iter.HasNext(); iter.Next() {
		fmt.Println(iter.Get())
	}
	// Output:
	// -436
	// 5
	// 6
	// 8
	// 10
	// 12
	// 15
}
