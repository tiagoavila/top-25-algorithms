package leealgorithmshortestpathinamaze

import "testing"

func TestFindShortedsPathInMaze(t *testing.T) {
	data := [][]int{
		{1, 1, 1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 0, 1, 0, 1},
		{0, 0, 1, 0, 1, 1, 1, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 0, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 1, 0, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 0, 0, 1, 1, 1},
		{0, 0, 1, 0, 0, 1, 1, 0, 0, 1},
	}
	source := Coord{0, 0}
	destination := Coord{7, 5}
	expectedPathSize := 12

	shortestPathSize := FindShortedsPathInMaze(data, source, destination)
	if expectedPathSize != shortestPathSize {
		t.Errorf("Expected path size %d, but got %d", expectedPathSize, shortestPathSize)
	}
}
