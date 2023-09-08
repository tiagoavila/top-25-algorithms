package dfs

import (
	"reflect"
	"testing"
	"top25algorithms/graph"
)

func TestBFSWithExampleFromTop25Algorithms(t *testing.T) {
	graph := graph.NewGraph()

	// Adding edges to the graph
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(2, 6)
	graph.AddEdge(5, 9)
	graph.AddEdge(5, 10)
	graph.AddEdge(4, 7)
	graph.AddEdge(4, 8)
	graph.AddEdge(7, 11)
	graph.AddEdge(7, 12)

	// Perform BFS starting from vertex 1.
	result := DFS(graph, 1)

	// Define the expected result.
	expected := []int{1, 2, 5, 9, 10, 6, 3, 4, 7, 11, 12, 8}

	// Check if the result matches the expected result.
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("DFS result is incorrect. Got: %v, Expected: %v", result, expected)
	}
}
