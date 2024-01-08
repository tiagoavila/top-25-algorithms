package kadane

import "math"

func FindMaximumSumSubarray(arr []int) int {
	maxSoFar, maxEndingHere := math.MinInt, 0

	if len(arr) == 1 {
		return arr[0]
	}

	for i := 0; i < len(arr); i++ {
		maxEndingHere = maxEndingHere + arr[i]

		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}

		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}

	return maxSoFar
}
