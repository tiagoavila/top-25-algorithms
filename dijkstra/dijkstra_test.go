package dijkstra

import (
	"reflect"
	"testing"
	"top25algorithms/weighteddirectedgraph"
)

func TestFindShortestPathWithDijkstra(t *testing.T) {
	// Define test cases using a table
	tests := []struct {
		description  string
		startNode    string
		endNode      string
		expectedPath []string
	}{
		{"Test Case 1", "A", "D", []string{"A", "E", "D"}},
		{"Test Case 2", "A", "B", []string{"A", "E", "B"}},
		{"Test Case 3", "A", "C", []string{"A", "E", "B", "C"}},
		{"Test Case 4", "A", "E", []string{"A", "E"}},
	}

	// Create a sample graph for each test case
	graph := createSampleGraph()

	// Iterate through the test cases
	for _, test := range tests {
		// Call the function to find the shortest path
		resultPath := FindShortestPathWithDijkstra(graph, test.startNode, test.endNode)

		// Compare the result with the expected path
		if !reflect.DeepEqual(resultPath, test.expectedPath) {
			t.Errorf("%s: Expected path %v, but got %v", test.description, test.expectedPath, resultPath)
		}
	}
}

func TestFindShortestPathWithDijkstraElemarLessonExample(t *testing.T) {
	graph := createSampleGraphFromElemarLessonExample()

	// Call the function to find the shortest path
	resultPath := FindShortestPathWithDijkstra(graph, "S", "E")

	expectedPath := []string{"S", "B", "H", "G", "E"}

	// Compare the result with the expected path
	if !reflect.DeepEqual(resultPath, expectedPath) {
		t.Errorf("Expected path %v, but got %v", expectedPath, resultPath)
	}
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

func createSampleGraphFromElemarLessonExample() *weighteddirectedgraph.WeightedDirectedGraph {
	// Create a weighted directed graph
	graph := weighteddirectedgraph.NewGraph()

	// Add edges to the graph
	graph.AddEdge("S", "A", 7)
	graph.AddEdge("S", "B", 2)
	graph.AddEdge("S", "C", 3)
	graph.AddEdge("A", "B", 3)
	graph.AddEdge("A", "D", 4)
	graph.AddEdge("B", "D", 4)
	graph.AddEdge("B", "H", 1)
	graph.AddEdge("C", "L", 2)
	graph.AddEdge("D", "F", 5)
	graph.AddEdge("E", "K", 5)
	graph.AddEdge("F", "H", 3)
	graph.AddEdge("G", "E", 2)
	graph.AddEdge("H", "G", 2)
	graph.AddEdge("L", "I", 4)
	graph.AddEdge("L", "J", 4)
	graph.AddEdge("I", "J", 6)
	graph.AddEdge("I", "K", 4)
	graph.AddEdge("J", "K", 4)

	return graph
}
