package insertionsort

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		j := i

		for j > 0 && array[j-1] > array[j] {
			array = swap(array, j, j-1)
			j--
		}
	}

	return array
}

func swap(array []int, i, j int) []int {
	array[i], array[j] = array[j], array[i]

	return array
}
