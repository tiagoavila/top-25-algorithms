package kadane

import (
	"testing"
)

func TestFindMaximumSumSubarray(t *testing.T) {
	testScenarios := []struct {
		array       []int
		expectedSum int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		// {[]int{-2, -3, 4, -1, -2, 1, 5, -3}, 7},
		// {[]int{1}, 1},
		// {[]int{5, 4, -1, 7, 8}, 23},
		// {[]int{-1}, -1},
		{[]int{-2, -1}, -1},
		{[]int{-1, -2}, -1},
	}

	for _, test := range testScenarios {
		resultSum := FindMaximumSumSubarray(test.array)
		if resultSum != test.expectedSum {
			t.Errorf("Expected sum %d, but got %d", test.expectedSum, resultSum)
		}
	}
}
