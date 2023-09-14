package quicksort

func Partition(array []int) []int {
	pivot, left, right := len(array)-1, 0, len(array)-2

	for left < right {
		for array[left] < array[pivot] {
			left++
		}

		for array[right] > array[pivot] {
			right--
		}

		if left < right {
			array[left], array[right] = array[right], array[left]
		} else {
			array[left], array[pivot] = array[pivot], array[left]
		}
	}

	return array
}

func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	if len(array) == 2 {
		if array[0] > array[1] {
			array[0], array[1] = array[1], array[0]
		}

		return array
	}

	pivot, left, right := len(array)-1, 0, len(array)-2

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

	leftSort := QuickSort(array[:left])
	rightSort := QuickSort(array[left+1:])

	return append(append(leftSort, array[left]), rightSort...)
}
