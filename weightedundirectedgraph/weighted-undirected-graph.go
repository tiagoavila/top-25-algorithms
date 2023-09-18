package weightedundirectedgraph

import (
	"fmt"
	"sort"
)

type Edge struct {
	src    int     // Source vertex
	dest   int     // Destination vertex
	weight float64 // Weight of the edge
}

type WeightedUndirectedGraph struct {
	vertices  int            // Number of vertices in the graph
	adjacency map[int][]Edge // Adjacency list to store edges
}

func NewGraph(vertices int) *WeightedUndirectedGraph {
	return &WeightedUndirectedGraph{
		vertices:  vertices,
		adjacency: make(map[int][]Edge),
	}
}

func (g *WeightedUndirectedGraph) AddEdge(src, dest int, weight float64) {
	// Add an undirected edge between src and dest with the given weight
	edge := Edge{src, dest, weight}
	g.adjacency[src] = append(g.adjacency[src], edge)

	// Since it's an undirected graph, we add an edge in the opposite direction as well
	edge = Edge{dest, src, weight}
	g.adjacency[dest] = append(g.adjacency[dest], edge)
}

func (g *WeightedUndirectedGraph) PrintGraph() {
	for vertex, edges := range g.adjacency {
		fmt.Printf("Vertex %d -> ", vertex)
		for _, edge := range edges {
			fmt.Printf("(Dest: %d, Weight: %.2f) ", edge.dest, edge.weight)
		}
		fmt.Println()
	}
}

func (g *WeightedUndirectedGraph) GetUniqueEdges() []Edge {
	uniqueEdges := make(map[string]Edge)
	for _, edges := range g.adjacency {
		for _, edge := range edges {
			// Create a unique key for each edge based on the src and dest vertices
			key := fmt.Sprintf("%d-%d", edge.src, edge.dest)

			// Check if the reverse key exists (e.g., if there's an edge from A to B, check if there's an edge from B to A)
			reverseKey := fmt.Sprintf("%d-%d", edge.dest, edge.src)
			_, reverseExists := uniqueEdges[reverseKey]

			// If the reverse edge doesn't exist, add the current edge to the uniqueEdges map
			if !reverseExists {
				uniqueEdges[key] = edge
			}
		}
	}

	// Convert the uniqueEdges map to a slice
	var result []Edge
	for _, edge := range uniqueEdges {
		result = append(result, edge)
	}

	return result
}

func (g *WeightedUndirectedGraph) SortEdgesByWeight(edges []Edge) []Edge {
	// Use the sort.Slice function to sort the edges by weight in ascending order
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	return edges
}
