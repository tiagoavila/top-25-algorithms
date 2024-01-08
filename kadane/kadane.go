package kadane

import "math"

func FindMaximumSumSubarray(nums []int) int {
	maxSoFar, maxEndingHere := math.MinInt, 0

	if len(nums) == 1 {
		return nums[0]
	}

	for i := 0; i < len(nums); i++ {
		maxEndingHere = maxEndingHere + nums[i]

		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}

		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}

	return maxSoFar
}
