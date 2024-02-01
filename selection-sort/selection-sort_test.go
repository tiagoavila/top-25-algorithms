package selectionsort

import (
	"reflect"
	"testing"
)

func TestISelectionSort(t *testing.T) {
	input := []int{4, 2, 7, 1, 3}
	expected := []int{1, 2, 3, 4, 7}

	result := SelectionSort(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
