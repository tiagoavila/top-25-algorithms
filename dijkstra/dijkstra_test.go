package dijkstra

import (
	"testing"
	"top25algorithms/weighteddirectedgraph"
)

func TestDijkstra(t *testing.T) {
	// Create a sample graph for testing
	graph := createSampleGraph()

	// Call the Dijkstra function with source and destination vertices
	Dijkstra(graph, "A", "B")
}

// Helper function to create a sample graph for testing
func createSampleGraph() *weighteddirectedgraph.WeightedDirectedGraph {
	// Create a weighted directed graph
	graph := weighteddirectedgraph.NewGraph()

	// Add edges to the graph
	graph.AddEdge("A", "B", 10)
	graph.AddEdge("A", "E", 3)
	graph.AddEdge("B", "C", 2)
	graph.AddEdge("B", "E", 4)
	graph.AddEdge("C", "D", 9)
	graph.AddEdge("D", "C", 7)
	graph.AddEdge("E", "B", 1)
	graph.AddEdge("E", "C", 8)
	graph.AddEdge("E", "D", 2)

	return graph
}
