package binary_search

import (
	"testing"
)

func TestBinarySearchFindsNumberInArray(t *testing.T) {
	nums := []int{2, 3, 5, 7, 9}
	target := 7
	result := BinarySearch(nums, target)
	if result != 3 {
		t.Errorf("binarySearch(%v, %d) = %d; expected 3", nums, target, result)
	}
}

func TestBinarySearchReturnsMinusOneWhenNumberNotInArray(t *testing.T) {
	nums := []int{2, 3, 5, 7, 9}
	target := 8
	result := BinarySearch(nums, target)
	if result != -1 {
		t.Errorf("binarySearch(%v, %d) = %d; expected -1", nums, target, result)
	}
}
