package mergesort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{38, 27, 43, 3, 9, 82, 10}, []int{3, 9, 10, 27, 38, 43, 82}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := MergeSort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
