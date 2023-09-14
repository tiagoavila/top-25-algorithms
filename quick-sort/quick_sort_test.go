package quicksort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 5, 2, 1, 6, 3}, []int{0, 1, 2, 3, 6, 5}},
		{[]int{9, -3, 5, 2, 6, 8, -6, 1, 3}, []int{1, -3, -6, 2, 3, 8, 5, 9, 6}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := Partition(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 5, 2, 1, 6, 3}, []int{0, 1, 2, 3, 5, 6}},
		{[]int{9, -3, 5, 2, 6, 8, -6, 1, 3}, []int{-6, -3, 1, 2, 3, 5, 6, 8, 9}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := QuickSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
