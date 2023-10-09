package weightedundirectedgraph

import (
	"reflect"
	"sort"
	"testing"
)

func TestWeightedUndirectedGraph(t *testing.T) {
	g := NewGraph(5)

	// Add edges to the graph
	g.AddEdge(0, 1, 2.5)
	g.AddEdge(1, 2, 1.8)
	g.AddEdge(2, 3, 3.0)
	g.AddEdge(3, 4, 2.1)

	// Check the number of vertices
	if g.Vertices != 5 {
		t.Errorf("Expected 5 vertices, got %d", g.Vertices)
	}

	// Check the edges and their weights
	expectedEdges := []Edge{
		{0, 1, 2.5},
		{1, 0, 2.5},
		{1, 2, 1.8},
		{2, 1, 1.8},
		{2, 3, 3.0},
		{3, 2, 3.0},
		{3, 4, 2.1},
		{4, 3, 2.1},
	}

	for _, edge := range expectedEdges {
		src := edge.Src
		dest := edge.Dest
		weight := edge.Weight

		found := false
		for _, e := range g.Adjacency[src] {
			if e.Dest == dest && e.Weight == weight {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Edge (%d, %d, %.2f) not found in the graph", src, dest, weight)
		}
	}
}

func TestGetVertices(t *testing.T) {
	g := NewGraph(5)

	// Add edges to the graph
	g.AddEdge(0, 1, 2.5)
	g.AddEdge(1, 2, 1.8)
	g.AddEdge(2, 3, 3.0)
	g.AddEdge(3, 4, 2.1)

	// Get the vertices
	vertices := g.GetVertices()

	// Sort the vertices for comparison
	sort.Ints(vertices)

	// Expected vertices
	expectedVertices := []int{0, 1, 2, 3, 4}

	// Check if the vertices match the expected ones
	if !reflect.DeepEqual(vertices, expectedVertices) {
		t.Errorf("Expected vertices: %v, got: %v", expectedVertices, vertices)
	}
}
