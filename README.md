[![GoDoc](https://godoc.org/github.com/evanhyd/sgl?status.svg)](https://godoc.org/github.com/evanhyd/sgl)
[![Go Report Card](https://goreportcard.com/badge/github.com/evanhyd/sgl)](https://goreportcard.com/report/github.com/evanhyd/sgl)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/evanhyd/sgl/blob/master/LICENSE)

# Standard Generic Library (SGL)
A generic data structure library in go.

# Feature Data Structures

- [AVL Tree](#avl_tree)
- [Priority Queue](#priority_queue)
  - [Binary Heap](#binary_heap)
  - [Binomial Heap](#binomial_heap)
- [Dynamic Array](#dynamic_array)
- [Singly LinkedList](#singly_linkedlist)
- [Trie](#trie)

# Priority_Queue
An abstract data type that defines operations to access the highest priority element.
```go
type PriorityQueue[T any] interface {
	Len() int
	Push(T)
	Pop()
	Top() T
}
```

# AVL_Tree  
![image](https://i.imgur.com/IfNd3vg.png)  
A self-balance binary tree in linked list representation.  
Supports fast insert, search, and delete key-value pairs.  
  
Insert: Θ(log n)  
Get:  Θ(log n)  
Delete:  Θ(log n)  
## [Benchmark](https://github.com/evanhyd/sgl/blob/main/avl_tree/AVLTree_test.go)  
	// BenchmarkAVLTree_Insert_Small-16    6201751         199.5 ns/op        48 B/op        1 allocs/op
	// BenchmarkAVLTree_Insert_Big-16      2670862         421.3 ns/op       192 B/op        1 allocs/op
 

# Binary_Heap  
![image](https://i.imgur.com/lUg8KI3.png)  
A complete binary tree in array representation.   
Support fast push, get + delete min/max elements.  

Push: Θ(log n)  
Top: Θ(1)  
Pop: Θ(log n)  

## [Benchmark](https://github.com/evanhyd/sgl/blob/main/binary_heap/BinaryHeap_test.go)    
    // BenchmarkBinaryHeap_Push_Small-16    51245478	       22.89 ns/op	      45 B/op	       0 allocs/op
    // BenchmarkBinaryHeap_Push_Big-16    	 7348918	       148.1 ns/op	     218 B/op	       1 allocs/op
    // BenchmarkBinaryHeap_Pop_Small-16    	 5932274	       210.9 ns/op	       0 B/op	       0 allocs/op
    // BenchmarkBinaryHeap_Pop_Big-16      	 2242784	       605.9 ns/op	       0 B/op	       0 allocs/op
    

# Binomial_Heap  
![image](https://i.imgur.com/rzox0pW.png)  
A heap in linked list representation.  
Support fast push, get + delete min/max elements, and merge.  

Push: Θ(log n)  
Top: Θ(log n)  
Pop: Θ(log n)  
Merge: Θ(log n)  

## [Benchmark](https://github.com/evanhyd/sgl/blob/main/binomial_heap/BinomialHeap_test.go)    
    // BenchmarkBinomialHeap_Push_Small-16      20313676          54.18 ns/op       24 B/op        1 allocs/op
    // BenchmarkBinomialHeap_Push_Big-16         6652496          155.1 ns/op      192 B/op        1 allocs/op
    // BenchmarkBinomialHeap_Pop_Small-16        8291036          154.5 ns/op        0 B/op        0 allocs/op
    // BenchmarkBinomialHeap_Pop_Big-16          6747440          187.2 ns/op        0 B/op        0 allocs/op

# Dynamic_Array  
![image](https://i.imgur.com/Ig9i7uV.png)  
A classic dynamic array.  
Support fast pushback, and popback in good amortized time.  

PushBack: Θ(1) average  
PopBack: Θ(1)  

## [Benchmark](https://github.com/evanhyd/sgl/blob/main/dynamic_array/DynamicArray_test.go)    
The dynamic array internally uses GO's built-in slice implementation.  

# Singly_LinkedList  
A linked list with each node tracks its child nodes.  
Support fast push front, and pop front to access the first element.  

PushBack: Θ(1)    
PopBack: Θ(1)  
Insert: Θ(1) require iterator  
Delete: Θ(1) require iterator  

## [Benchmark](https://github.com/evanhyd/sgl/blob/main/singly_linkedlist/SinglyLinkedList_test.go)    
- Missing  

# Trie  
An optimized prefix tree in linked list representation.  
Support fast insert, get, delete string-value pairs.  

Insert: Θ(|string|)  
Get: Θ(|string|)  
Remove: Θ(|string|)  

## [Benchmark](https://github.com/evanhyd/sgl/blob/main/trie/Trie_test.go)    
    // BenchmarkTrie_Insert_Small-16    	 5659839	       216.3 ns/op	      97 B/op	       2 allocs/op  
    // BenchmarkTrie_Insert_Big-16      	 3737061	       320.9 ns/op	     257 B/op	       2 allocs/op  
    // BenchmarkTrie_Get_Small-16        	15389232	       89.86 ns/op	       0 B/op	       0 allocs/op  
    // BenchmarkTrie_Get_Big-16          	14818125	       102.1 ns/op	       0 B/op	       0 allocs/op  
    // BenchmarkTrie_Remove_Small-16    	 6894930	       180.0 ns/op	      63 B/op	       1 allocs/op  
    // BenchmarkTrie_Remove_Big-16    	 6774646	       190.3 ns/op	      63 B/op	       1 allocs/op  
