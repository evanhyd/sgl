package avl_tree

import (
	"fmt"
	"math/rand"
	"testing"
)

func assertTree(t *testing.T, tree AVLTree[int, int]) bool {
	len := 0

	var height func(*node[int, int]) int
	height = func(n *node[int, int]) int {
		if n == nil {
			return 0
		}
		return max(height(n.left), height(n.right)) + 1
	}

	var validateNode func(*node[int, int]) bool
	validateNode = func(root *node[int, int]) bool {
		if root == nil {
			return true
		}
		len++

		balanceFactor := height(root.right) - height(root.left)
		if balanceFactor < -1 || balanceFactor > 1 {
			t.Fatalf("balance factor of node %d is %d", root.key, balanceFactor)
			return false
		}
		return validateNode(root.left) && validateNode(root.right)
	}

	result := validateNode(tree.root)
	if tree.Len() != len {
		t.Fatalf("tree node count = %d, wanted %d", tree.Len(), len)
	}
	return result
}

// Helper function to print the AVL tree horizontally.
func printAVLTree(tree AVLTree[int, int]) {
	ifs := func(condition bool, a, b string) string {
		if condition {
			return a
		}
		return b
	}

	var printTree func(*node[int, int], string)
	printTree = func(node *node[int, int], prefix string) {
		fmt.Printf("(%v, %v)\n", node.key, node.value)

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

func TestAVLTree_Insert_Ascend(t *testing.T) {
	const testSize = 1 << 10
	tree := New[int, int](func(a, b int) int { return a - b })
	keys := make([]int, testSize)
	for i := range keys {
		keys[i] = i
	}

	for i, e := range keys {
		tree.Insert(e, e)
		if expected := i + 1; tree.Len() != expected {
			t.Fatalf("Len() = %d, want %d", tree.Len(), expected)
		}
		assertTree(t, tree)
	}
}

func TestAVLTree_Insert_Descend(t *testing.T) {
	const testSize = 1 << 10
	tree := New[int, int](func(a, b int) int { return a - b })
	keys := make([]int, testSize)
	for i := range keys {
		keys[i] = len(keys) - i
	}

	for i, e := range keys {
		tree.Insert(e, e)
		if expected := i + 1; tree.Len() != expected {
			t.Fatalf("Len() = %d, want %d", tree.Len(), expected)
		}
		assertTree(t, tree)
	}
}

func TestAVLTree_Insert_Duplication(t *testing.T) {
	const testSize = 1 << 10
	tree := New[int, int](func(a, b int) int { return a - b })
	keys := make([]int, testSize)
	for i := range keys {
		keys[i] = i / 2
	}

	for i, e := range keys {
		tree.Insert(e, e)
		if expected := i/2 + 1; tree.Len() != expected {
			t.Fatalf("Len() = %d, want %d", tree.Len(), expected)
		}
		assertTree(t, tree)
	}
}

func TestAVLTree_Insert_Random(t *testing.T) {
	const testSize = 1 << 10
	tree := New[int, int](func(a, b int) int { return a - b })
	sTree := map[int]int{}
	keys := append(rand.Perm(testSize), rand.Perm(testSize)...)

	for i, e := range keys {
		tree.Insert(e, i)
		sTree[e] = i
		if expected := len(sTree); tree.Len() != expected {
			t.Fatalf("Len() = %d, want %d", tree.Len(), expected)
		}
		assertTree(t, tree)
	}
}

func TestAVLTree_Get(t *testing.T) {
	const testSize = 1 << 10
	tree := New[int, int](func(a, b int) int { return a - b })
	sTree := map[int]int{}
	keys := append(rand.Perm(testSize), rand.Perm(testSize)...)

	for i := 0; i < testSize; i++ {
		eValue, eBool := sTree[keys[i]]
		if aValue, aBool := tree.Get(keys[i]); eValue != aValue || eBool != aBool {
			t.Fatalf("Get(%v) = (%v, %v), want (%v, %v)", keys[i], aValue, aBool, eValue, eBool)
		}

		tree.Insert(keys[i], i)
		sTree[keys[i]] = i

		eValue, eBool = sTree[keys[i]]
		if aValue, aBool := tree.Get(keys[i]); eValue != aValue || eBool != aBool {
			t.Fatalf("Get(%v) = (%v, %v), want (%v, %v)", keys[i], aValue, aBool, eValue, eBool)
		}

		assertTree(t, tree)
	}
}

func TestAVLTree_Remove(t *testing.T) {
	const testSize = 1 << 10
	tree := New[int, int](func(a, b int) int { return a - b })
	sTree := map[int]int{}
	keys := rand.Perm(testSize)

	for i, key := range keys {
		tree.Insert(key, i)
		sTree[key] = i
		assertTree(t, tree)
	}

	//remove non existed keys
	fakeKeys := rand.Perm(testSize)
	for i := range fakeKeys {
		fakeKeys[i] += testSize
	}
	keys = append(keys, fakeKeys...)

	for i, key := range keys {
		tree.Remove(key)
		delete(sTree, key)

		if len(sTree) != tree.Len() {
			t.Fatalf("Len() = %v, want %v", tree.Len(), len(sTree))
		}

		_, expected := sTree[key]
		_, actual := tree.Get(key)
		if expected != actual {
			t.Fatalf("Contain(%v) = %v, want %v", keys[i], actual, expected)
		}
		assertTree(t, tree)
	}
}

func TestAVLTree_Min(t *testing.T) {
	const testSize = 1 << 10
	tree := New[int, int](func(a, b int) int { return a - b })
	keys := rand.Perm(testSize)
	for _, key := range keys {
		tree.Insert(key, key)
	}

	expected := 0
	if key, _ := tree.Min(); key != expected {
		t.Errorf("Min() = %d, want %d", key, expected)
	}
}

func TestAVLTree_Max(t *testing.T) {
	const testSize = 1 << 10
	tree := New[int, int](func(a, b int) int { return a - b })
	keys := rand.Perm(testSize)
	for _, key := range keys {
		tree.Insert(key, key)
	}

	expected := testSize - 1
	if key, _ := tree.Max(); key != expected {
		t.Errorf("Max() = %d, want %d", key, expected)
	}
}

func TestIterator(t *testing.T) {
	const testSize = 1 << 10
	tree := New[int, int](func(a, b int) int { return a - b })
	keys := make([]int, testSize)
	for i := range keys {
		keys[i] = i
	}
	for i, key := range keys {
		tree.Insert(key, i)
	}

	i := 0
	for iter := tree.Begin(); iter.HasNext(); iter.Next() {
		if key, value := iter.Get(); key != keys[i] || value != i {
			t.Errorf("Get() = (%d, %d), want (%d, %d)", key, value, keys[i], i)
		}
		i++
	}
}

func BenchmarkAVLTree_Insert_Small(b *testing.B) {
	// int64
	// BenchmarkAVLTree_Insert_Small-16    	 6201751	       199.5 ns/op	      48 B/op	       1 allocs/op
	// BenchmarkAVLTree_Insert_Small-16    	 6220729	       201.5 ns/op	      48 B/op	       1 allocs/op
	// BenchmarkAVLTree_Insert_Small-16    	 5196642	       209.7 ns/op	      48 B/op	       1 allocs/op

	tree := New[int64, int64](func(a, b int64) int { return int(a - b) })
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(int64(i), int64(i))
	}
}

func BenchmarkAVLTree_Insert_Big(b *testing.B) {
	// [20]int64
	// BenchmarkAVLTree_Insert_Big-16    	 2670862	       421.3 ns/op	     192 B/op	       1 allocs/op
	// BenchmarkAVLTree_Insert_Big-16    	 2667331	       413.9 ns/op	     192 B/op	       1 allocs/op
	// BenchmarkAVLTree_Insert_Big-16    	 2506341	       416.8 ns/op	     192 B/op	       1 allocs/op

	type Large [20]int64
	tree := New[Large, int64](func(a, b Large) int { return int(a[0] - b[0]) })
	keys := make([]Large, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = Large{int64(i)}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(keys[i], int64(i))
	}
}

func ExampleAVLTree_Insert() {
	tree := New[int, int](func(a, b int) int { return a - b })
	tree.Insert(10, 20)
	tree.Insert(5, 10)
	tree.Insert(15, 125)
	fmt.Println(tree.Len())
	// Output:
	// 3
}

func ExampleAVLTree_Get() {
	tree := New[int, int](func(a, b int) int { return a - b })
	tree.Insert(10, 20)
	tree.Insert(5, 10)
	tree.Insert(15, 125)

	fmt.Println(tree.Get(10))
	fmt.Println(tree.Get(199))
	// Output:
	// 20 true
	// 0 false
}

func ExampleAVLTree_Remove() {
	tree := New[int, int](func(a, b int) int { return a - b })
	tree.Insert(10, 20)
	tree.Insert(5, 10)
	tree.Insert(15, 125)

	fmt.Println(tree.Get(5))
	tree.Remove(5)
	fmt.Println(tree.Get(5))
	// Output:
	// 10 true
	// 0 false
}

func ExampleAVLTree_Min() {
	tree := New[int, int](func(a, b int) int { return a - b })
	tree.Insert(10, 20)
	tree.Insert(5, 10)
	tree.Insert(15, 125)

	fmt.Println(tree.Min())
	// Output:
	// 5 10
}

func ExampleAVLTree_Max() {
	tree := New[int, int](func(a, b int) int { return a - b })
	tree.Insert(10, 20)
	tree.Insert(5, 10)
	tree.Insert(15, 125)

	fmt.Println(tree.Max())
	// Output:
	// 15 125
}

func ExampleIterator() {
	tree := New[int, int](func(a, b int) int { return a - b })
	tree.Insert(10, 4)
	tree.Insert(-436, 8)
	tree.Insert(5, 3)
	tree.Insert(15, 12)
	tree.Insert(12, 54)
	tree.Insert(8, 123)
	tree.Insert(6, 83)

	for iter := tree.Begin(); iter.HasNext(); iter.Next() {
		fmt.Println(iter.Get())
	}
	// Output:
	// -436 8
	// 5 3
	// 6 83
	// 8 123
	// 10 4
	// 12 54
	// 15 12
}
