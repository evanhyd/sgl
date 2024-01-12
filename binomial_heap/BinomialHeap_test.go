package binomial_heap

import (
	"fmt"
	"testing"
)

func TestBinomialHeap_Len(t *testing.T) {
	heap := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	// Test Len on an empty heap
	if len := heap.Len(); len != 0 {
		t.Errorf("Len() = %d, want 0", len)
	}

	// Test Len after pushing elements
	heap.Push(5)
	heap.Push(3)
	heap.Push(7)
	if len := heap.Len(); len != 3 {
		t.Errorf("Len() = %d, want 3", len)
	}
}

func TestBinomialHeap_Push(t *testing.T) {
	heap := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	// Test Push and Top
	heap.Push(5)
	if top := heap.Top(); top != 5 {
		t.Errorf("Top() = %d, want 5", top)
	}

	// Test Push on an empty heap
	heap = BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	heap.Push(10)
	if top := heap.Top(); top != 10 {
		t.Errorf("Top() = %d, want 10", top)
	}

	// Test Push with multiple elements
	heap.Push(8)
	heap.Push(3)
	if top := heap.Top(); top != 10 {
		t.Errorf("Top() = %d, want 10", top)
	}

	// Test Push with elements in descending order
	descHeap := BinomialHeap[int]{Cmp: func(a, b int) int { return b - a }}
	descHeap.Push(20)
	descHeap.Push(15)
	descHeap.Push(10)
	if top := descHeap.Top(); top != 10 {
		t.Errorf("Top() = %d, want 10", top)
	}

	// Test Push with repeated elements
	repeatHeap := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	repeatHeap.Push(5)
	repeatHeap.Push(5)
	repeatHeap.Push(5)
	if top := repeatHeap.Top(); top != 5 {
		t.Errorf("Top() = %d, want 5", top)
	}

	// Test Push with lots of elements
	bigHeap := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	for i := 0; i < 100; i++ {
		bigHeap.Push(i)
	}

	for i := 0; i < 100; i++ {
		actualTop := bigHeap.Top()
		wantTop := 99 - i
		if actualTop != wantTop {
			t.Errorf("Top() = %d, want %d", actualTop, wantTop)
		}
		bigHeap.Pop()
	}
}

func TestBinomialHeap_Pop(t *testing.T) {
	heap := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}

	// Test Pop after pushing elements
	heap.Push(5)
	heap.Push(3)
	heap.Push(7)
	heap.Pop()
	if len := heap.Len(); len != 2 {
		t.Errorf("Pop(), Len() = %d, want 2", len)
	}
	if top := heap.Top(); top != 5 {
		t.Errorf("Pop(), Top() = %d, want 5", top)
	}

	// Test Pop until empty
	heap.Pop()
	heap.Pop()
	if len := heap.Len(); len != 0 {
		t.Errorf("Multiple Pop(), Len() = %d, want 0", len)
	}
}

func TestBinomialHeap_Top(t *testing.T) {
	heap := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}

	// Test Top after pushing elements
	heap.Push(5)
	heap.Push(3)
	heap.Push(7)
	if top := heap.Top(); top != 7 {
		t.Errorf("Top() = %d, want 7", top)
	}

	// Test Top after Pop
	heap.Pop()
	if top := heap.Top(); top != 5 {
		t.Errorf("Top() = %d, want 5", top)
	}
}

func TestBinomialHeap_Merge(t *testing.T) {
	// Test Merge with two non-empty heaps
	heap1 := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	heap1.Push(5)
	heap1.Push(3)

	heap2 := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	heap2.Push(10)
	heap2.Push(8)

	heap1.Merge(heap2)

	if len := heap1.Len(); len != 4 {
		t.Errorf("Merge(), Len() = %d, want 4", len)
	}
	if top := heap1.Top(); top != 10 {
		t.Errorf("Merge(), Top() = %d, want 10", top)
	}

	// Test Merge with one empty heap
	heap3 := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	heap3.Push(2)
	heap3.Push(4)

	heap4 := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	heap4.Merge(heap3)

	if len := heap4.Len(); len != 2 {
		t.Errorf("Merge(), Len() = %d, want 2", len)
	}
	if top := heap4.Top(); top != 4 {
		t.Errorf("Merge(), Top() = %d, want 4", top)
	}

	// Test merge two big heaps
	bigHeap1 := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	bigHeap2 := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	for i := -100; i < 0; i++ {
		bigHeap1.Push(i)
	}
	for i := 0; i <= 100; i++ {
		bigHeap2.Push(i)
	}

	expectedLen := bigHeap1.Len() + bigHeap2.Len()
	bigHeap1.Merge(bigHeap2)
	if bigHeap1.Len() != expectedLen {
		t.Errorf("Len() = %d, want %d", bigHeap1.Len(), expectedLen)
	}

	for i := 100; i >= -100; i-- {
		actualTop := bigHeap1.Top()
		if actualTop != i {
			t.Errorf("Top() = %d, want %d", actualTop, i)
		}
		bigHeap1.Pop()
	}
}

func BenchmarkBinomialHeap_Push_Small(b *testing.B) {
	// int64
	// BenchmarkBinomialHeap_Push_Small-16    	19541204	        55.06 ns/op	      24 B/op	       1 allocs/op
	heap := BinomialHeap[int64]{Cmp: func(l, r int64) int { return int(l - r) }}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		heap.Push(int64(i % 1024))
	}
}

func BenchmarkBinomialHeap_Push_Big(b *testing.B) {
	// [20]int64
	// BenchmarkBinomialHeap_Push_Big-16    	 7704684	       141.6 ns/op	     176 B/op	       1 allocs/op

	type Large [20]int64
	heap := BinomialHeap[Large]{Cmp: func(l, r Large) int { return int(l[0] - r[0]) }}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		heap.Push(Large{int64(i % 1024)})
	}
}

func BenchmarkBinomialHeap_Pop_Small(b *testing.B) {
	// int64
	// BenchmarkBinomialHeap_Pop_Small-16    	 8985741	       251.9 ns/op	       0 B/op	       0 allocs/op
	heap := BinomialHeap[int64]{Cmp: func(l, r int64) int { return int(l - r) }}
	for i := 0; i < b.N; i++ {
		heap.Push(int64(i % 1024))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		heap.Pop()
	}
}

func BenchmarkBinomialHeap_Pop_Big(b *testing.B) {
	// [20]int64
	// BenchmarkBinomialHeap_Pop_Big-16    	 3046507	       540.5 ns/op	       0 B/op	       0 allocs/op
	type Large [20]int64
	heap := BinomialHeap[Large]{Cmp: func(l, r Large) int { return int(l[0] - r[0]) }}
	for i := 0; i < b.N; i++ {
		heap.Push(Large{int64(i % 1024)})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		heap.Pop()
	}
}

func ExampleBinomialHeap_Len() {
	heap := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	fmt.Println(heap.Len())
	heap.Push(123)
	fmt.Println(heap.Len())

	// Output:
	// 0
	// 1
}

func ExampleBinomialHeap_Push() {
	heap := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	heap.Push(5)
	heap.Push(3)
	heap.Push(7)
	fmt.Println(heap.Len())
	heap.Push(10)
	fmt.Println(heap.Len())
	fmt.Println(heap.Top())
	// Output:
	// 3
	// 4
	// 10
}

func ExampleBinomialHeap_Pop() {
	heap := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	heap.Push(5)
	heap.Push(3)
	heap.Push(7)
	fmt.Println(heap.Top())
	heap.Pop()
	fmt.Println(heap.Top())
	// Output:
	// 7
	// 5
}

func ExampleBinomialHeap_Top() {
	heap := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	heap.Push(5)
	heap.Push(3)
	heap.Push(7)
	fmt.Println(heap.Top())
	// Output: 7
}

func ExampleBinomialHeap_Merge() {
	// Merge two binomial heaps and print the top element after merging
	heap1 := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	heap1.Push(5)
	heap1.Push(3)

	heap2 := BinomialHeap[int]{Cmp: func(a, b int) int { return a - b }}
	heap2.Push(10)
	heap2.Push(8)

	heap1.Merge(heap2)
	fmt.Println(heap1.Len())
	fmt.Println(heap1.Top())

	// Output:
	// 4
	// 10
}
