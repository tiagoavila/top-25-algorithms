package bellmanford

import (
	"fmt"
	"testing"
	"top25algorithms/weighteddirectedgraph"
)

func TestPrintGraph(t *testing.T) {
	graph := createSampleGraph()

	fmt.Println("Graph", graph)

	BellmanFord(graph, "A", 5)
}

func createSampleGraph() *weighteddirectedgraph.WeightedDirectedGraph {
	// Create a weighted directed graph
	graph := weighteddirectedgraph.NewGraph()

	// Add edges to the graph
	graph.AddEdge("A", "B", -1)
	graph.AddEdge("A", "C", 4)
	graph.AddEdge("B", "C", 3)
	graph.AddEdge("B", "D", 2)
	graph.AddEdge("B", "E", 2)
	graph.AddEdge("D", "C", 5)
	graph.AddEdge("D", "B", 1)
	graph.AddEdge("E", "D", -3)

	return graph
}
