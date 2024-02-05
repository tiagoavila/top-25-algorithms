package countingsort

import (
	"reflect"
	"testing"
)

// TestCountSort tests the CountSort function to ensure it correctly sorts input arrays.
func TestCountSort(t *testing.T) {
	// Define test cases with input arrays and expected sorted arrays
	testCases := []struct {
		input    []int // Input array to be sorted
		expected []int // Expected sorted array
	}{
		{input: []int{4, 2, 2, 8, 3, 3, 1}, expected: []int{1, 2, 2, 3, 3, 4, 8}},
		{input: []int{1, 4, 1, 2, 7, 5, 2}, expected: []int{1, 1, 2, 2, 4, 5, 7}},
		{input: []int{10, 7, 12, 4, 9, 13}, expected: []int{4, 7, 9, 10, 12, 13}},
		{input: []int{}, expected: []int{}},   // Test with an empty array
		{input: []int{5}, expected: []int{5}}, // Test with a single-element array
		// Adding the new test case from the image
		{input: []int{2, 5, 3, 0, 2, 3, 0, 3}, expected: []int{0, 0, 2, 2, 3, 3, 3, 5}},
	}

	for _, tc := range testCases {
		// Run CountSort on the input array
		output := CountSort(tc.input)

		// Check if the output matches the expected sorted array
		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("CountSort(%v) = %v, expected %v", tc.input, output, tc.expected)
		}
	}
}
