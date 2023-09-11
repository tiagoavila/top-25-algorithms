package mergesort

func MergeSort(array []int) []int {
	// Stopping condition for the recursion
	if len(array) <= 1 {
		return array
	}

	// Split the input array into two halves and recursively call MergeSort on them
	middle := len(array) / 2
	left := MergeSort(array[:middle])
	right := MergeSort(array[middle:])

	// Merge the sorted halves and return the result
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	leftIndex := 0
	rightIndex := 0
	result := make([]int, 0, len(left)+len(right))

	// Compare elements from both halves and merge them into the result slice
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	// Add remaining elements from left and right if any
	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result
}
