package singlylinkedlist

import (
	"fmt"
	"testing"
)

func TestSinglyLinkedList_Len(t *testing.T) {
	list := SinglyLinkedList[int]{}
	if len := list.Len(); len != 0 {
		t.Errorf("Len() = %d, want %d", len, 0)
	}

	list.PushFront(1)
	if len := list.Len(); len != 1 {
		t.Errorf("Len() = %d, want %d", len, 1)
	}

	list.PushFront(2)
	if len := list.Len(); len != 2 {
		t.Errorf("Len() = %d, want %d", len, 2)
	}
}

func TestSinglyLinkedList_PushFront(t *testing.T) {
	list := SinglyLinkedList[int]{}
	list.PushFront(1)
	if len := list.Len(); len != 1 {
		t.Errorf("Len() = %d, want %d", len, 1)
	}

	list.PushFront(2)
	if len := list.Len(); len != 2 {
		t.Errorf("Len() = %d, want %d", len, 2)
	}
}

func TestSinglyLinkedList_PopFront(t *testing.T) {
	list := SinglyLinkedList[int]{}
	list.PushFront(1)
	list.PushFront(2)

	list.PopFront()
	if len := list.Len(); len != 1 {
		t.Errorf("Len() = %d, want %d", len, 1)
	}

	list.PopFront()
	if len := list.Len(); len != 0 {
		t.Errorf("Len() = %d, want %d", len, 0)
	}
}

func TestSinglyLinkedList_Begin(t *testing.T) {
	list := SinglyLinkedList[int]{}
	list.PushFront(1)
	list.PushFront(2)

	iterator := list.Begin()
	if value := iterator.Get(); value != 2 {
		t.Errorf("Begin().Get() = %d, want %d", value, 2)
	}
}

func TestSinglyLinkedList_Insert(t *testing.T) {
	list := SinglyLinkedList[int]{}

	// Test case 1: Insert into an empty list
	list.Insert(list.Begin(), 1)
	if len := list.Len(); len != 1 {
		t.Errorf("SinglyLinkedList.Len() = %d, want 1", len)
	}

	// Test case 2: Insert at the beginning
	list.Insert(list.Begin(), 0)
	if len := list.Len(); len != 2 {
		t.Errorf("SinglyLinkedList.Len() = %d, want 2", len)
	}

	if iter := list.Begin(); iter.Get() != 0 {
		t.Errorf("Iterator.Get() = %d, want 0", iter.Get())
	}

	// Test case 3: Insert at the end
	end := list.Begin()
	end.Next()
	end.Next()
	list.Insert(end, 2)
	if len := list.Len(); len != 3 {
		t.Errorf("SinglyLinkedList.Len() = %d, want 3", len)
	}
	if end.Get() != 2 {
		t.Errorf("Iterator.Get() = %d, want 2", end.Get())
	}

	// Test case 4: Insert in the middle
	iter := list.Begin()
	iter.Next()
	list.Insert(iter, 1)
	if len := list.Len(); len != 4 {
		t.Errorf("SinglyLinkedList.Len() = %d, want 4", len)
	}

	expected := []int{0, 1, 1, 2}
	iter = list.Begin()
	for _, val := range expected {
		if iter.Get() != val {
			t.Errorf("Iterator.Get() = %d, want %d", iter.Get(), val)
		}
		iter.Next()
	}
}

func TestSinglyLinkedList_Remove(t *testing.T) {

	// Test case 1: Remove from the beginning
	list := SinglyLinkedList[int]{}
	list.PushFront(2)
	list.PushFront(1)
	list.PushFront(0)
	list.Remove(list.Begin())
	if len := list.Len(); len != 2 {
		t.Errorf("SinglyLinkedList.Len() = %d, want 2", len)
	}

	iter := list.Begin()
	if iter.Get() != 1 {
		t.Errorf("Iterator.Get() = %d, want 1", iter.Get())
	}
	iter.Next()
	if iter.Get() != 2 {
		t.Errorf("Iterator.Get() = %d, want 2", iter.Get())
	}

	// Test case 2: Remove from the middle
	list = SinglyLinkedList[int]{}
	list.PushFront(2)
	list.PushFront(1)
	list.PushFront(0)
	iter = list.Begin()
	iter.Next()
	list.Remove(iter)
	if len := list.Len(); len != 2 {
		t.Errorf("SinglyLinkedList.Len() = %d, want 2", len)
	}

	iter = list.Begin()
	if iter.Get() != 0 {
		t.Errorf("Iterator.Get() = %d, want 0", iter.Get())
	}
	iter.Next()
	if iter.Get() != 2 {
		t.Errorf("Iterator.Get() = %d, want 2", iter.Get())
	}

	// Test case 3: Remove from the end
	list = SinglyLinkedList[int]{}
	list.PushFront(2)
	list.PushFront(1)
	list.PushFront(0)
	iter = list.Begin()
	iter.Next()
	iter.Next()
	list.Remove(iter)
	if len := list.Len(); len != 2 {
		t.Errorf("SinglyLinkedList.Len() = %d, want 2", len)
	}

	iter = list.Begin()
	if iter.Get() != 0 {
		t.Errorf("Iterator.Get() = %d, want 0", iter.Get())
	}
	iter.Next()
	if iter.Get() != 1 {
		t.Errorf("Iterator.Get() = %d, want 1", iter.Get())
	}

	// Test case 4: Remove from a list with one element
	list = SinglyLinkedList[int]{}
	list.PushFront(123)
	list.Remove(list.Begin())
	if len := list.Len(); len != 0 {
		t.Errorf("SinglyLinkedList.Len() = %d, want 0", len)
	}
}

func TestIterator_HasNext(t *testing.T) {
	list := SinglyLinkedList[int]{}
	iterator := list.Begin()
	if hasNext := iterator.HasNext(); hasNext {
		t.Errorf("HasNext() = %t, want %t", hasNext, false)
	}

	list.PushFront(1)
	iterator = list.Begin()
	if hasNext := iterator.HasNext(); !hasNext {
		t.Errorf("HasNext() = %t, want %t", hasNext, true)
	}
}

func ExampleSinglyLinkedList_PushFront() {
	list := SinglyLinkedList[int]{}
	list.PushFront(42)
	fmt.Println(list.Len())
	// Output: 1
}

func ExampleSinglyLinkedList_PopFront() {
	list := SinglyLinkedList[int]{}
	list.PushFront(42)
	list.PopFront()
	fmt.Println(list.Len())
	// Output: 0
}

func ExampleSinglyLinkedList_Insert() {
	list := SinglyLinkedList[int]{}
	iterator := list.Begin()
	list.Insert(iterator, 10)
	fmt.Println(list.Len())
	// Output: 1
}

func ExampleSinglyLinkedList_Remove() {
	list := SinglyLinkedList[int]{}
	iterator := list.Begin()
	list.Insert(iterator, 10)
	fmt.Println(list.Len())
	list.Remove(iterator)
	fmt.Println(list.Len())
	// Output:
	// 1
	// 0
}

func ExampleIterator_Get() {
	list := SinglyLinkedList[int]{}
	list.PushFront(5)
	list.PushFront(3)
	iter := list.Begin()
	fmt.Println(iter.Get())
	// Output:
	// 3
}

func ExampleIterator_Set() {
	list := SinglyLinkedList[int]{}
	list.PushFront(5)
	list.PushFront(3)
	iter := list.Begin()
	iter.Set(10)
	fmt.Println(iter.Get())
	// Output:
	// 10
}

func ExampleIterator_Next() {
	list := SinglyLinkedList[int]{}
	list.PushFront(5)
	list.PushFront(3)
	iterator := list.Begin()
	iterator.Next()
	fmt.Println(iterator.HasNext())
	// Output: true
}

func ExampleIterator_HasNext() {
	list := SinglyLinkedList[int]{}
	list.PushFront(5)
	list.PushFront(3)
	iterator := list.Begin()
	fmt.Println(iterator.HasNext())
	iterator.Next()
	fmt.Println(iterator.HasNext())
	iterator.Next()
	fmt.Println(iterator.HasNext())
	// Output:
	// true
	// true
	// false
}
