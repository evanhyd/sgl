package dynamic_array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDynamicArray_PushBack(t *testing.T) {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)
	da.PushBack(99)

	expected := DynamicArray[int]{42, 12, 99}

	if !reflect.DeepEqual(da, expected) {
		t.Errorf("PushBack failed, got: %v, want: %v", da, expected)
	}
}

func TestDynamicArray_PopBack(t *testing.T) {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)
	da.PushBack(99)

	da.PopBack()

	expected := DynamicArray[int]{42, 12}

	if !reflect.DeepEqual(da, expected) {
		t.Errorf("PopBack failed, got: %v, want: %v", da, expected)
	}
}

func TestDynamicArray_Front(t *testing.T) {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)
	da.PushBack(99)

	front := *da.Front()

	if front != 42 {
		t.Errorf("Front() = %v, want: %v", front, 42)
	}
}

func TestDynamicArray_Back(t *testing.T) {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)
	da.PushBack(99)

	back := *da.Back()

	if back != 99 {
		t.Errorf("Back() = %v, want: %v", back, 99)
	}
}

func TestDynamicArray_Len(t *testing.T) {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)
	da.PushBack(99)

	length := da.Len()

	if length != 3 {
		t.Errorf("Len() = %v, want: %v", length, 3)
	}
}

func TestDynamicArray_Cap(t *testing.T) {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)
	da.PushBack(99)

	capacity := da.Cap()

	if capacity != 4 {
		t.Errorf("Cap() = %v, want: %v", capacity, 4)
	}
}

func TestIterator_Get(t *testing.T) {
	da := DynamicArray[int]{1, 2, 3}
	iterator := Iterator[int]{arr: &da, index: 1}

	result := iterator.Get()
	expected := &da[1]

	if result != expected {
		t.Errorf("Get() = %v, want %v", result, expected)
	}
}

func TestIterator_Next(t *testing.T) {
	da := DynamicArray[int]{1, 2, 3}
	iterator := Iterator[int]{arr: &da, index: 1}

	iterator.Next()

	result := iterator.index
	expected := 2

	if result != expected {
		t.Errorf("Next() resulted in index %v, want %v", result, expected)
	}
}

func TestIterator_HasNext(t *testing.T) {
	da := DynamicArray[int]{1, 2, 3}
	iterator := Iterator[int]{arr: &da, index: 1}

	result := iterator.HasNext()
	expected := true

	if result != expected {
		t.Errorf("HasNext() = %v, want %v", result, expected)
	}

	// Advance the iterator to the last element
	iterator.Next()
	iterator.Next()

	result = iterator.HasNext()
	expected = false

	if result != expected {
		t.Errorf("HasNext() = %v, want %v", result, expected)
	}
}

func TestDynamicArray_Begin(t *testing.T) {
	da := DynamicArray[int]{1, 2, 3}
	iterator := da.Begin()

	if iterator.index != 0 {
		t.Errorf("Begin() resulted in index %v, want %v", iterator.index, 0)
	}

	value := iterator.Get()
	expectedValue := &da[0]
	if value != expectedValue {
		t.Errorf("Get() = %v, want %v", value, expectedValue)
	}
}

func TestDynamicArray_End(t *testing.T) {
	da := DynamicArray[int]{1, 2, 3}
	iterator := da.End()

	// Verify that End() returns an iterator one pass the last element
	if iterator.index != len(da) {
		t.Errorf("TestDynamicArray_End: End() did not return an iterator one pass the last element")
	}
}

func ExampleDynamicArray_PushBack() {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)

	fmt.Println(da)
	// Output: [42 12]
}

func ExampleDynamicArray_PopBack() {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)

	da.PopBack()

	fmt.Println(da)
	// Output: [42]
}

func ExampleDynamicArray_Front() {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)
	da.PushBack(99)

	fmt.Println(*da.Front())
	// Output: 42
}

func ExampleDynamicArray_Back() {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)
	da.PushBack(99)

	fmt.Println(*da.Back())
	// Output: 99
}

func ExampleDynamicArray_Len() {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)
	da.PushBack(99)

	fmt.Println(da.Len())
	// Output: 3
}

func ExampleDynamicArray_Cap() {
	da := make(DynamicArray[int], 0)

	da.PushBack(42)
	da.PushBack(12)
	da.PushBack(99)

	fmt.Println(da.Cap())
	// Output: 4
}

func ExampleDynamicArray_Begin() {
	da := DynamicArray[int]{1, 2, 3}
	iterator := da.Begin()
	fmt.Print(*iterator.Get())
	// Output: 1
}

func ExampleDynamicArray_End() {
	da := DynamicArray[int]{1, 2, 3}
	iterator := da.End()
	fmt.Print(iterator.index)
	// Output: 3
}

func ExampleIterator() {
	da := DynamicArray[int]{1, 2, 3}
	for iter := da.Begin(); iter != da.End(); iter.Next() {
		fmt.Println(*iter.Get())
	}

	for iter := da.Begin(); iter.HasNext(); iter.Next() {
		fmt.Println(*iter.Get())
	}

	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
}
