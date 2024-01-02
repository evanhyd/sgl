package main

import (
	"sgl/binomial_heap"
	"strings"
)

func main() {
	a := binomial_heap.BinomialHeap[string]{}
	a.Cmp = strings.Compare
}
