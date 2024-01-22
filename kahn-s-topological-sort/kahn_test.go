package kahnstopologicalsort

import (
	"reflect"
	"testing"
	"top25algorithms/unweighteddirectedgraph"
)

func TestTopologicalSort(t *testing.T) {
	graph := unweighteddirectedgraph.NewGraph()

	// Add edges to create a directed acyclic graph (DAG)
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")

	expectedOrder := []string{"A", "B", "C", "D"}
	sorted := TopologicalSort(graph)

	if !reflect.DeepEqual(sorted, expectedOrder) {
		t.Errorf("Topological sort result is incorrect. Expected: %v, Got: %v", expectedOrder, sorted)
	}
}

// https://www.techiedelight.com/kahn-topological-sort-algorithm/
func TestTopologicalSortExampleFromArticle(t *testing.T) {
	graph := unweighteddirectedgraph.NewGraph()

	// Add edges to create a directed acyclic graph (DAG)
	graph.AddEdge("0", "6")
	graph.AddEdge("1", "2")
	graph.AddEdge("1", "4")
	graph.AddEdge("1", "6")
	graph.AddEdge("3", "0")
	graph.AddEdge("3", "4")
	graph.AddEdge("5", "1")
	graph.AddEdge("7", "0")
	graph.AddEdge("7", "1")

	//3   5   7   0   1   2   6   4
	expectedOrder := []string{"3", "5", "7", "0", "1", "2", "4", "6"}
	sorted := TopologicalSort(graph)

	if !reflect.DeepEqual(sorted, expectedOrder) {
		t.Errorf("Topological sort result is incorrect. Expected: %v, Got: %v", expectedOrder, sorted)
	}
}

func TestTopologicalSortExampleFromGeeksForGeeks(t *testing.T) {
	graph := unweighteddirectedgraph.NewGraph()

	// Add edges to create a directed acyclic graph (DAG)
	graph.AddEdge("2", "3")
	graph.AddEdge("3", "1")
	graph.AddEdge("4", "0")
	graph.AddEdge("4", "1")
	graph.AddEdge("5", "0")
	graph.AddEdge("5", "2")

	expectedOrder := []string{"4", "5", "0", "2", "3", "1"}
	sorted := TopologicalSort(graph)

	if !reflect.DeepEqual(sorted, expectedOrder) {
		t.Errorf("Topological sort result is incorrect. Expected: %v, Got: %v", expectedOrder, sorted)
	}
}
