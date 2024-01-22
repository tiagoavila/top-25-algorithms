package unweighteddirectedgraph

import (
	"testing"
)

func TestAddEdgeToNewVertex(t *testing.T) {
	graph := NewGraph()

	src := "A"
	dest := "B"

	graph.AddEdge(src, dest)

	// Check if the destination vertex is added to the source vertex's edges
	srcVerticeData, ok := graph.Vertices[src]
	if !ok {
		t.Errorf("Source vertex not found in the graph")
	}

	if len(srcVerticeData.Edges) != 1 || srcVerticeData.Edges[0].Dest != dest {
		t.Errorf("Edge not added to the new vertex correctly")
	}

	// Check InDegree of the destination vertex
	destVerticeData, ok := graph.Vertices[dest]
	if !ok {
		t.Errorf("Destination vertex not found in the graph")
	}

	if destVerticeData.InDegree != 1 {
		t.Errorf("InDegree of the destination vertex is not correct")
	}
}

func TestAddEdgeToExistingVertex(t *testing.T) {
	graph := NewGraph()

	src := "A"
	dest1 := "B"
	dest2 := "C"

	// Add an initial edge
	graph.AddEdge(src, dest1)

	// Add another edge to the existing vertex
	graph.AddEdge(src, dest2)

	// Check if the second destination vertex is added to the source vertex's edges
	srcVerticeData, ok := graph.Vertices[src]
	if !ok {
		t.Errorf("Source vertex not found in the graph")
	}

	if len(srcVerticeData.Edges) != 2 || srcVerticeData.Edges[1].Dest != dest2 {
		t.Errorf("Edge not added to the existing vertex correctly")
	}
}

func TestInDegreeOfNewVertex(t *testing.T) {
	graph := NewGraph()

	src := "A"
	src2 := "B"
	dest := "C"

	graph.AddEdge(src, dest)
	graph.AddEdge(src2, dest)

	// Check InDegree of the destination vertex
	destVerticeData, ok := graph.Vertices[dest]
	if !ok {
		t.Errorf("Destination vertex not found in the graph")
	}

	if destVerticeData.InDegree != 2 {
		t.Errorf("InDegree of the destination vertex is not correct")
	}
}
