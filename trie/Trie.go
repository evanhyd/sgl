package trie

import (
	"unicode/utf8"
)

type node[V any] struct {
	value    V
	end      bool
	children map[rune]*node[V]
}

// A trie that maps a string to a value, supports unicode.
type Trie[V any] struct {
	root node[V]
	len  int
}

func New[V any]() Trie[V] {
	return Trie[V]{root: node[V]{children: map[rune]*node[V]{}}}
}

// Return the number of element.
func (t *Trie[V]) Len() int {
	return t.len
}

// Insert a key value pair to the trie.
//
// If the key value pair entry already exists, it updates the value.
func (t *Trie[V]) Insert(key string, value V) {
	curr := &t.root

	for _, r := range key {
		child, exist := curr.children[r]
		if !exist {
			child = &node[V]{children: map[rune]*node[V]{}}
			curr.children[r] = child
		}
		curr = child
	}

	if !curr.end {
		curr.end = true
		t.len++
	}
	curr.value = value
}

// Return the value and the exist indicator.
//
// If the key exists, it returns (value, true).
//
// Otherwise, it returns (zero value, false).
func (t *Trie[V]) Get(key string) (value V, exist bool) {
	curr := &t.root
	for _, r := range key {
		curr, exist = curr.children[r]
		if !exist {
			return
		}
	}
	return curr.value, curr.end
}

// Remove key entry from the tree.
func (t *Trie[V]) Remove(key string) {
	curr := &t.root
	nodes := make([]*node[V], 0, utf8.RuneCountInString(key)+1)
	nodes = append(nodes, curr)
	for _, r := range key {
		var exist bool
		curr, exist = curr.children[r]
		if !exist {
			return
		}
		nodes = append(nodes, curr)
	}

	if !curr.end {
		return
	}
	curr.end = false
	t.len--

	i := len(nodes) - 1
	for _, r := range key {
		node := nodes[i]
		if len(node.children) == 0 && !node.end {
			i--
			delete(nodes[i].children, r)
		} else {
			break
		}
	}
}
