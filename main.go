package main

import (
	"strings"

	"github.com/evanhyd/sgl/binomial_heap"
)

func main() {
	a := binomial_heap.BinomialHeap[string]{}
	a.Cmp = strings.Compare
}
