package trie

import (
	"fmt"
	"slices"
	"strconv"
	"testing"
)

func makeStringSequence(n int) []string {
	seq := make([]string, n)
	for i := 0; i < n; i++ {
		seq[i] = strconv.FormatInt(int64(i), 10)
	}
	return seq
}

func TestNewTrie(t *testing.T) {
	trie := New[int]()
	if actual := trie.Len(); actual != 0 {
		t.Errorf("Len() = %v, want 0", actual)
	}
}

func TestTrie_Insert(t *testing.T) {
	trie := New[int]()

	if actual := trie.Len(); actual != 0 {
		t.Errorf("Len() = %v, want 0", actual)
	}

	//random string and prefix test
	words := []string{"a", "aaa", "b", "bbb", "aa", "bb", "abcd", "efgh", "aabb", "", " ", "   "}
	for i, word := range words {
		trie.Insert(word, i)
		if actual := trie.Len(); actual != i+1 {
			t.Errorf("Len() = %v, want 1", actual)
		}
	}

	//duplication test
	expected := trie.Len()
	words = []string{"efgh", "aaa", "   ", "abcd", "aabb", "", " ", "bbb", "a", "b"}
	for i, word := range words {
		trie.Insert(word, i)
		if actual := trie.Len(); actual != expected {
			t.Errorf("Len() = %v, want %v", actual, expected)
		}
	}

	//random test
	words = makeStringSequence(10000)
	trie = New[int]()
	table := map[string]int{}
	for i, word := range words {
		trie.Insert(word, i)
		table[word] = i
		actual := trie.Len()
		expected := len(table)
		if actual != expected {
			t.Errorf("Len() = %v, want %v", actual, expected)
		}
	}

	//random test
	words = makeStringSequence(10000)
	slices.Reverse(words)
	trie = New[int]()
	table = map[string]int{}
	for i, word := range words {
		trie.Insert(word, i)
		table[word] = i
		actual := trie.Len()
		expected := len(table)
		if actual != expected {
			t.Errorf("Len() = %v, want %v", actual, expected)
		}
	}
}

func TestTrie_Get(t *testing.T) {
	trie := New[int]()

	//random string and prefix test
	words := []string{"a", "bbb", "b", "aaa", "abcd", "efgh", "aabb", "", "   "}
	for i, word := range words {
		trie.Insert(word, i)
	}

	for i, word := range words {
		v, exist := trie.Get(word)
		if !exist || v != i {
			t.Errorf("Get(%v) = (%v, %v), want (%v, %v)", word, v, exist, i, true)
		}
	}

	//update test
	words = []string{"efgh", "aaa", "   ", "abcd", "aabb", "", "bbb", "a", "b"}
	for i, word := range words {
		trie.Insert(word, i)
	}
	for i, word := range words {
		v, exist := trie.Get(word)
		if !exist || v != i {
			t.Errorf("Get(%v) = (%v, %v), want (%v, %v)", word, v, exist, i, true)
		}
	}

	//empty test
	words = []string{"aa", "bb", "aaaa", "bbbb", "  ", "    ", "aaak", "abcdd", "a abb", "9", " ?", "bbbd"}
	for _, word := range words {
		v, exist := trie.Get(word)
		if exist || v != 0 {
			t.Errorf("Get(%v) = (%v, %v), want (%v, %v)", word, v, exist, false, 0)
		}
	}

	//random test
	words = makeStringSequence(10000)
	trie = New[int]()
	table := map[string]int{}
	for i, word := range words {
		trie.Insert(word, i)
		table[word] = i
		a1, a2 := trie.Get(word)
		e1, e2 := table[word]
		if a1 != e1 || a2 != e2 {
			t.Errorf("Get(%v) = (%v, %v), want (%v, %v)", word, a1, a2, e1, e2)
		}
	}

	//reverse random test
	words = makeStringSequence(10000)
	slices.Reverse(words)
	trie = New[int]()
	table = map[string]int{}
	for i, word := range words {
		trie.Insert(word, i)
		table[word] = i
		a1, a2 := trie.Get(word)
		e1, e2 := table[word]
		if a1 != e1 || a2 != e2 {
			t.Errorf("Get(%v) = (%v, %v), want (%v, %v)", word, a1, a2, e1, e2)
		}
	}
}

func TestTrie_Remove(t *testing.T) {
	trie := New[int]()

	//random string and prefix test
	words := []string{"a", "bbb", "b", "aaa", "abcd", "efgh", "aabb", "", "   "}
	for i, word := range words {
		trie.Insert(word, i)
	}

	//remove non existed words
	expected := trie.Len()
	nonExisted := []string{"aa", "bb", "  ", "abc", "efghi", " aabb"}
	for _, word := range nonExisted {
		trie.Remove(word)
		if actual := trie.Len(); actual != expected {
			t.Errorf("Len() = %v, want %v", actual, expected)
		}
		_, exist := trie.Get(word)
		if exist {
			t.Errorf("Get(%v) = (_, %v), want (_, %v)", word, true, false)
		}
	}

	//remove actual words
	words = []string{"aabb", "a", "aaa", "bbb", "b", "abcd", "efgh"}
	expected = trie.Len()
	for _, word := range words {
		trie.Remove(word)
		expected--
		if actual := trie.Len(); actual != expected {
			t.Errorf("Len() = %v, want %v", actual, expected)
		}
		_, exist := trie.Get(word)
		if exist {
			t.Errorf("Get(%v) = (_, %v), want (_, %v)", word, true, false)
		}
	}

	//random test
	words = makeStringSequence(10000)
	trie = New[int]()
	table := map[string]int{}
	for i, word := range words {
		trie.Insert(word, i)
		table[word] = i
	}

	for _, word := range words {
		trie.Remove(word)
		delete(table, word)
		_, a2 := trie.Get(word)
		_, e2 := table[word]
		if a2 != e2 {
			t.Errorf("Get(%v) = (_, %v), want (_, %v)", word, a2, e2)
		}
	}

	//random test
	words = makeStringSequence(10000)
	slices.Reverse(words)
	trie = New[int]()
	table = map[string]int{}
	for i, word := range words {
		trie.Insert(word, i)
		table[word] = i
	}

	for _, word := range words {
		trie.Remove(word)
		delete(table, word)
		_, a2 := trie.Get(word)
		_, e2 := table[word]
		if a2 != e2 {
			t.Errorf("Get(%v) = (_, %v), want (_, %v)", word, a2, e2)
		}
	}
}

// BenchmarkTrie_Insert_Small-16    	 5659839	       216.3 ns/op	      97 B/op	       2 allocs/op
// BenchmarkTrie_Insert_Small-16    	 5504846	       225.0 ns/op	      97 B/op	       2 allocs/op
// BenchmarkTrie_Insert_Small-16    	 5610189	       220.6 ns/op	      97 B/op	       2 allocs/op
func BenchmarkTrie_Insert_Small(b *testing.B) {
	trie := New[int32]()
	keys := makeStringSequence(b.N)

	b.ResetTimer()
	for i, key := range keys {
		trie.Insert(key, int32(i))
	}
}

// BenchmarkTrie_Insert_Big-16    	 3737061	       320.9 ns/op	     257 B/op	       2 allocs/op
// BenchmarkTrie_Insert_Big-16    	 3795207	       319.8 ns/op	     257 B/op	       2 allocs/op
// BenchmarkTrie_Insert_Big-16    	 3676011	       331.0 ns/op	     257 B/op	       2 allocs/op
func BenchmarkTrie_Insert_Big(b *testing.B) {
	type Large [20]int64
	trie := New[Large]()
	keys := makeStringSequence(b.N)
	values := make([]Large, b.N)
	for i := 0; i < b.N; i++ {
		values[i] = Large{int64(i)}
	}

	b.ResetTimer()
	for i, key := range keys {
		trie.Insert(key, values[i])
	}
}

// BenchmarkTrie_Get_Small-16    	15389232	        89.86 ns/op	       0 B/op	       0 allocs/op
// BenchmarkTrie_Get_Small-16    	12966272	        89.96 ns/op	       0 B/op	       0 allocs/op
// BenchmarkTrie_Get_Small-16    	13567185	        89.33 ns/op	       0 B/op	       0 allocs/op
func BenchmarkTrie_Get_Small(b *testing.B) {
	trie := New[int32]()
	keys := makeStringSequence(b.N)
	for i, key := range keys {
		trie.Insert(key, int32(i))
	}

	b.ResetTimer()
	for _, key := range keys {
		_, _ = trie.Get(key)
	}
}

// BenchmarkTrie_Get_Big-16    	14818125	       102.1 ns/op	       0 B/op	       0 allocs/op
// BenchmarkTrie_Get_Big-16    	12904086	        94.06 ns/op	       0 B/op	       0 allocs/op
// BenchmarkTrie_Get_Big-16    	14814741	        96.28 ns/op	       0 B/op	       0 allocs/op
func BenchmarkTrie_Get_Big(b *testing.B) {
	type Large [20]int64
	trie := New[Large]()
	keys := makeStringSequence(b.N)
	values := make([]Large, b.N)
	for i := 0; i < b.N; i++ {
		values[i] = Large{int64(i)}
	}
	for i, key := range keys {
		trie.Insert(key, values[i])
	}

	b.ResetTimer()
	for _, key := range keys {
		_, _ = trie.Get(key)
	}
}

// BenchmarkTrie_Remove_Small-16    	 6894930	       180.0 ns/op	      63 B/op	       1 allocs/op
// BenchmarkTrie_Remove_Small-16    	 6486952	       178.5 ns/op	      63 B/op	       1 allocs/op
// BenchmarkTrie_Remove_Small-16    	 6252040	       179.2 ns/op	      63 B/op	       1 allocs/op
func BenchmarkTrie_Remove_Small(b *testing.B) {
	trie := New[int32]()
	keys := makeStringSequence(b.N)
	for i, key := range keys {
		trie.Insert(key, int32(i))
	}

	b.ResetTimer()
	for _, key := range keys {
		trie.Remove(key)
	}
}

// BenchmarkTrie_Remove_Big-16    	 6774646	       190.3 ns/op	      63 B/op	       1 allocs/op
// BenchmarkTrie_Remove_Big-16    	 6775078	       192.3 ns/op	      63 B/op	       1 allocs/op
// BenchmarkTrie_Remove_Big-16    	 6451324	       188.8 ns/op	      63 B/op	       1 allocs/op
func BenchmarkTrie_Remove_Big(b *testing.B) {
	type Large [20]int64
	trie := New[Large]()
	keys := makeStringSequence(b.N)
	values := make([]Large, b.N)
	for i := 0; i < b.N; i++ {
		values[i] = Large{int64(i)}
	}
	for i, key := range keys {
		trie.Insert(key, values[i])
	}

	b.ResetTimer()
	for _, key := range keys {
		trie.Remove(key)
	}
}

func ExampleTrie_Insert() {
	trie := New[int]()
	trie.Insert("hello", 1)
	trie.Insert("world", 2)
	trie.Insert("你好", 3)
	trie.Insert("世界", 4)
	fmt.Println(trie.Len())
	// Output:
	// 4
}

func ExampleTrie_Get() {
	trie := New[int]()
	trie.Insert("hello, 世界", 123)

	fmt.Println(trie.Get("hello, world"))
	fmt.Println(trie.Get("hello, 世界"))
	// Output:
	// 0 false
	// 123 true
}

func ExampleTrie_Remove() {
	trie := New[int]()
	trie.Insert("hello, 世界", 123)

	value, exist := trie.Get("hello, 世界")
	fmt.Println(value, exist)

	trie.Remove("hello, 世界")
	_, exist = trie.Get("hello, 世界")
	fmt.Println(exist)

	// Output:
	// 123 true
	// false
}
