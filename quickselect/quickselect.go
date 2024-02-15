package quickselect

func Partition(array []int, left, right int) ([]int, int) {
	pivot := right
	right--

	for left < right {
		for array[left] < array[pivot] {
			left++
		}

		for array[right] > array[pivot] && right > 0 {
			right--
		}

		if left < right {
			array[left], array[right] = array[right], array[left]
		} else {
			array[left], array[pivot] = array[pivot], array[left]
		}
	}

	return array, left
}

func QuickSelect(arr []int, low, hi, k int) int {
	pivotedArray, pivotIdx := Partition(arr, low, hi)

	// if partition value is equal to the kth position,
	// return value at k.
	if pivotIdx == k {
		return pivotedArray[k]
	}

	// if partition value is less than kth position,
	// search right side of the array.
	if pivotIdx < k {
		return QuickSelect(pivotedArray, pivotIdx+1, hi, k)
	} else {
		// if partition value is greater than kth position,
		// search left side of the array.
		return QuickSelect(pivotedArray, low, pivotIdx-1, k)
	}
}
