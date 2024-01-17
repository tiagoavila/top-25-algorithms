package floodfillalgorithm

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	// Define the test matrix
	testMatrix := [][]string{
		{"Y", "Y", "Y", "G", "G", "G", "G", "G", "G", "G"},
		{"Y", "Y", "Y", "Y", "Y", "Y", "G", "X", "X", "X"},
		{"G", "G", "G", "G", "G", "G", "G", "X", "X", "X"},
		{"W", "W", "W", "W", "W", "G", "G", "G", "G", "X"},
		{"W", "R", "R", "R", "R", "R", "G", "X", "X", "X"},
		{"W", "W", "W", "R", "R", "G", "G", "X", "X", "X"},
		{"W", "B", "X", "R", "R", "R", "R", "R", "R", "X"},
		{"W", "B", "B", "B", "X", "R", "R", "X", "X", "X"},
		{"W", "B", "B", "X", "B", "B", "B", "B", "X", "X"},
		{"W", "B", "B", "X", "X", "X", "X", "X", "X", "X"},
	}

	// Define the expected result after applying FloodFill
	expectedResult := [][]string{
		{"Y", "Y", "Y", "G", "G", "G", "G", "G", "G", "G"},
		{"Y", "Y", "Y", "Y", "Y", "Y", "G", "C", "C", "C"},
		{"G", "G", "G", "G", "G", "G", "G", "C", "C", "C"},
		{"W", "W", "W", "W", "W", "G", "G", "G", "G", "C"},
		{"W", "R", "R", "R", "R", "R", "G", "C", "C", "C"},
		{"W", "W", "W", "R", "R", "G", "G", "C", "C", "C"},
		{"W", "B", "X", "R", "R", "R", "R", "R", "R", "C"},
		{"W", "B", "B", "B", "C", "R", "R", "C", "C", "C"},
		{"W", "B", "B", "C", "B", "B", "B", "B", "C", "C"},
		{"W", "B", "B", "C", "C", "C", "C", "C", "C", "C"},
	}

	// Call the FloodFill function
	result := FloodFill(testMatrix, 3, 9, "C")

	// Check if the result matches the expected outcome
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("FloodFill result is incorrect, got: %v, want: %v", result, expectedResult)
	}
}
