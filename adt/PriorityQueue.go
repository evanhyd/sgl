package adt

type PriorityQueue[T any] interface {
	Len() int
	Push(T)
	Pop()
	Top() T
}
