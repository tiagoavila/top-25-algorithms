package heapsort

import (
	"reflect"
	"testing"
)

func TestHeapifyAnArrayWorks(t *testing.T) {
	testHeap := []int{100, 88, 87, 86, 50, 25, 16, 15, 12, 8, 3, 2}
	arrayToHeapify := []int{87, 3, 50, 86, 2, 88, 100, 12, 25, 15, 8, 16}

	heap := Heapify(arrayToHeapify) // Assuming Heapify is a method of MaxHeap

	for i, expected := range testHeap {
		poppedValue := Pop(heap) // Assuming Pop returns (int, error)

		if poppedValue != expected {
			t.Errorf("Expected %d, got %d at index %d", expected, poppedValue, i)
		}
	}
}

func TestHeapSort(t *testing.T) {
	// Given an unsorted array
	arrayToHeapify := []int{87, 3, 50, 86, 2, 88, 100, 12, 25, 15, 8, 16}

	// When HeapSort is called
	HeapSort(arrayToHeapify)

	// Then the array should be sorted in ascending order
	expected := []int{2, 3, 8, 12, 15, 16, 25, 50, 86, 87, 88, 100}

	if !reflect.DeepEqual(arrayToHeapify, expected) {
		t.Errorf("HeapSort failed: expected %v, got %v", expected, arrayToHeapify)
	}
}
