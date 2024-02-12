package heapsort

// https://stackoverflow.com/questions/70025649/heapify-vs-heap-sort-vs-construct-heap

type Heap struct {
	heapArray   []int
	latestIndex int
}

func NewHeap(size int) *Heap {
	return &Heap{
		heapArray:   make([]int, size),
		latestIndex: -1,
	}
}

func HeapSort(inputArray []int) []int {
	heap := Heapify(inputArray)

	for i := len(inputArray) - 1; i >= 0; i-- {
		inputArray[i] = Pop(heap)
	}

	return inputArray
}

// In a binary heap stored as an array, the relationship between parent nodes and their children can be described using
// the indices within the array. Given a zero-based index array:

// The children of the node at index i are at indices 2*i + 1 (left child) and 2*i + 2 (right child).
// Conversely, the parent of the node at index i can be found at index (i - 1) / 2 (using integer division).
// To find the index of the last node that is a parent, consider that any parent node must have at least one child.
// The last parent node in the array, therefore, would be the parent of the very last node in the array.

// Given an array of length n, the last node in the array is at index n - 1.
// The parent of this node is at index (n - 1 - 1) / 2 = (n - 2) / 2.

// So, the index of the last node that is a parent is not exactly half of the array length but rather (n - 2) / 2,
// where n is the total number of nodes

func Heapify(inputArray []int) *Heap {
	outputArray := make([]int, len(inputArray))
	copy(outputArray, inputArray)

	for i := (len(outputArray) - 2) / 2; i >= 0; i-- {
		doHeapify(&outputArray, i)
	}

	heap := NewHeap(len(outputArray))
	heap.heapArray = outputArray
	heap.latestIndex = len(outputArray) - 1

	return heap
}

func Pop(heap *Heap) int {
	poppedValue := heap.heapArray[0]

	heap.heapArray[0] = heap.heapArray[heap.latestIndex]
	heap.heapArray[heap.latestIndex] = 0
	heap.latestIndex--

	currentItemIndex := 0
	for hasChildWithAGreaterValue(&heap.heapArray, currentItemIndex) == true {
		childIndexWithGreatestValue := getChildWithGreatestValue(&heap.heapArray, currentItemIndex, heap.latestIndex)
		heap.heapArray[currentItemIndex], heap.heapArray[childIndexWithGreatestValue] = heap.heapArray[childIndexWithGreatestValue], heap.heapArray[currentItemIndex]
		currentItemIndex = childIndexWithGreatestValue
	}

	return poppedValue
}

func doHeapify(inputArray *[]int, index int) {
	indexOfLargestValue := index
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2

	if leftChildIndex < len(*inputArray) && (*inputArray)[leftChildIndex] > (*inputArray)[indexOfLargestValue] {
		indexOfLargestValue = leftChildIndex
	}

	if rightChildIndex < len(*inputArray) && (*inputArray)[rightChildIndex] > (*inputArray)[indexOfLargestValue] {
		indexOfLargestValue = rightChildIndex
	}

	if indexOfLargestValue != index {
		(*inputArray)[index], (*inputArray)[indexOfLargestValue] = (*inputArray)[indexOfLargestValue], (*inputArray)[index]
		doHeapify(inputArray, indexOfLargestValue)
	}
}

func hasChildWithAGreaterValue(inputArray *[]int, index int) bool {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2

	return leftChildIndex < len(*inputArray) && (*inputArray)[leftChildIndex] > (*inputArray)[index] ||
		rightChildIndex < len(*inputArray) && (*inputArray)[rightChildIndex] > (*inputArray)[index]
}

func getChildWithGreatestValue(inputArray *[]int, index int, heapLatestIndex int) int {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2

	if rightChildIndex > heapLatestIndex {
		return leftChildIndex
	}

	if (*inputArray)[rightChildIndex] > (*inputArray)[leftChildIndex] {
		return rightChildIndex
	}

	return leftChildIndex
}
