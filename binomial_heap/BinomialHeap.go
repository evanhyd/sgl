package binomial_heap

import (
	"math/bits"

	"github.com/evanhyd/sgl/adt"
)

type flagTree[T any] struct {
	key   T
	left  *flagTree[T]
	right *flagTree[T]
}

// A max binomial heap.
//
// It has higher constant for insert and pop, but supports merging in O(log(len)).
//
// interface: PriorityQueue
type BinomialHeap[T any] struct {
	trees []*flagTree[T]
	cmp   func(T, T) int
	len   int
}

var _ adt.PriorityQueue[int] = &BinomialHeap[int]{}

func New[T any](predicate func(T, T) int) BinomialHeap[T] {
	return BinomialHeap[T]{cmp: predicate}
}

// Return the number of element.
func (b *BinomialHeap[T]) Len() int {
	return b.len
}

// Push e to the heap.
func (b *BinomialHeap[T]) Push(e T) {
	b.len++
	b.reserve()
	b.mergeTree(&flagTree[T]{e, nil, nil}, 0)
}

// Remove the top element from the heap.
func (b *BinomialHeap[T]) Pop() {
	b.len--
	height := b.max()
	tree := b.trees[height].left
	b.trees[height] = nil

	//split the flag tree into smaller flag trees
	for height--; height >= 0; height-- {
		subTree := tree.right
		tree.right = nil
		b.mergeTree(tree, height)
		tree = subTree
	}
}

// Return the top of the heap.
func (b *BinomialHeap[T]) Top() T {
	return b.trees[b.max()].key
}

// Merge and destruct the heap into the current heap.
func (b *BinomialHeap[T]) Merge(heap BinomialHeap[T]) {
	b.len += heap.len
	b.reserve()
	for height, tree := range heap.trees {
		if tree != nil {
			b.mergeTree(tree, height)
		}
	}
}

// Find the top of the heap, return the index of the flag tree.
func (b *BinomialHeap[T]) max() int {
	m := -1
	for i, tree := range b.trees {
		if tree != nil && (m == -1 || b.cmp(b.trees[i].key, b.trees[m].key) > 0) {
			m = i
		}
	}
	return m
}

// Reserve enough memory for tree list based on len.
func (b *BinomialHeap[T]) reserve() {
	maxHeight := bits.Len(uint(b.len))
	if len(b.trees) < maxHeight {
		trees := make([]*flagTree[T], maxHeight)
		copy(trees, b.trees)
		b.trees = trees
	}
}

// Merge tree into the heap indexing by height
func (b *BinomialHeap[T]) mergeTree(tree *flagTree[T], height int) {
	for ; height < len(b.trees); height++ {
		t := b.trees[height]
		if t == nil {
			b.trees[height] = tree
			break
		}
		b.trees[height] = nil

		if b.cmp(t.key, tree.key) > 0 {
			tree, t = t, tree
		}
		tree.left, t.right = t, tree.left
	}
}
