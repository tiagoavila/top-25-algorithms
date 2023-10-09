package kruskaltofindmst

import (
	"testing"
	"top25algorithms/weightedundirectedgraph"
)

func TestKruskalMinimumSpanningTree1(t *testing.T) {
	graph := weightedundirectedgraph.NewGraph(4)

	// Adding edges to the graph
	graph.AddEdge(0, 1, 2)
	graph.AddEdge(0, 2, 3)
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 4)
	graph.AddEdge(2, 3, 5)

	// Calculate the minimum spanning tree
	minimumSpanningTree := KruskalMST(graph)

	// Expected minimum spanning tree
	expectedEdges := []weightedundirectedgraph.Edge{
		{Src: 1, Dest: 2, Weight: 1},
		{Src: 0, Dest: 1, Weight: 2},
		{Src: 1, Dest: 3, Weight: 4},
	}

	// Check if the calculated minimum spanning tree has the same edges as expected
	if !compareEdges(minimumSpanningTree, expectedEdges) {
		t.Errorf("Kruskal algorithm did not produce the expected minimum spanning tree.")
	}
}

func TestKruskalMinimumSpanningTree2(t *testing.T) {
	graph := weightedundirectedgraph.NewGraph(5)

	// Adding edges to the graph
	graph.AddEdge(0, 1, 1)
	graph.AddEdge(0, 2, 7)
	graph.AddEdge(0, 3, 10)
	graph.AddEdge(0, 4, 5)
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(2, 3, 4)
	graph.AddEdge(3, 4, 2)

	// Calculate the minimum spanning tree
	minimumSpanningTree := KruskalMST(graph)

	// Expected minimum spanning tree
	expectedEdges := []weightedundirectedgraph.Edge{
		{Src: 0, Dest: 1, Weight: 1},
		{Src: 3, Dest: 4, Weight: 2},
		{Src: 1, Dest: 2, Weight: 3},
		{Src: 2, Dest: 3, Weight: 4},
	}

	// Check if the calculated minimum spanning tree has the same edges as expected
	if !compareEdges(minimumSpanningTree, expectedEdges) {
		t.Errorf("Kruskal algorithm did not produce the expected minimum spanning tree.")
	}
}

func compareEdges(edges1, edges2 []weightedundirectedgraph.Edge) bool {
	if len(edges1) != len(edges2) {
		return false
	}

	for i := 0; i < len(edges1); i++ {
		if edges1[i].Src != edges2[i].Src || edges1[i].Dest != edges2[i].Dest || edges1[i].Weight != edges2[i].Weight {
			return false
		}
	}

	return true
}
