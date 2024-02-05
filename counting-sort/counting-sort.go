package countingsort

// https://www.geeksforgeeks.org/counting-sort/
func CountSort(inputArray []int) []int {
	// Step1: Find out the maximum element from the given array.
	max := 0
	for _, value := range inputArray {
		if value > max {
			max = value
		}
	}

	//Step 2: Initialize a countArray[] of length max+1 with all elements as 0.
	// This array will be used for storing the occurrences of the elements of the input array.
	countArray := make([]int, max+1)

	// Step 3: In the countArray[], store the count of each unique element of the input array at their respective indices.
	// For Example: The count of element 2 in the input array is 2. So, store 2 at index 2 in the countArray[].
	// Similarly, the count of element 5 in the input array is 1, hence store 1 at index 5 in the countArray[].
	for _, value := range inputArray {
		countArray[value]++
	}

	// Step 4: Store the cumulative sum or prefix sum of the elements of the countArray[] by doing
	// countArray[i] = countArray[i – 1] + countArray[i].
	// This will help in placing the elements of the input array at the correct index in the output array.
	for i := 1; i < len(countArray); i++ {
		countArray[i] += countArray[i-1]
	}

	// Step 5: Iterate from end of the input array and because traversing input array from end preserves the order
	// of equal elements, which eventually makes this sorting algorithm stable.
	outputArray := make([]int, len(inputArray))
	for i := len(inputArray) - 1; i >= 0; i-- {
		// Update outputArray[ countArray[ inputArray[i] ] – 1] = inputArray[i].
		value := inputArray[i]
		count := countArray[value] //Get the count of the value
		outputArray[count-1] = value

		// Also, update countArray[ inputArray[i] ] = countArray[ inputArray[i] ]– -.
		countArray[value]--
	}

	return outputArray
}
