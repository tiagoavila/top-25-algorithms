package selectionsort

func SelectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		indexOfLowest := i

		for j := i + 1; j < len(array); j++ {
			if array[j] < array[indexOfLowest] {
				indexOfLowest = j
			}
		}

		if indexOfLowest != i {
			array[i], array[indexOfLowest] = array[indexOfLowest], array[i]
		}
	}

	return array
}
