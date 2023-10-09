package weightedundirectedgraph

import (
	"fmt"
	"sort"

	"golang.org/x/exp/maps"
)

type Edge struct {
	Src    int     // Source vertex
	Dest   int     // Destination vertex
	Weight float64 // Weight of the edge
}

type WeightedUndirectedGraph struct {
	Vertices  int            // Number of vertices in the graph
	Adjacency map[int][]Edge // Adjacency list to store edges
}

func NewGraph(vertices int) *WeightedUndirectedGraph {
	return &WeightedUndirectedGraph{
		Vertices:  vertices,
		Adjacency: make(map[int][]Edge),
	}
}

func (g *WeightedUndirectedGraph) AddEdge(src, dest int, weight float64) {
	// Add an undirected edge between src and dest with the given weight
	edge := Edge{src, dest, weight}
	g.Adjacency[src] = append(g.Adjacency[src], edge)

	// Since it's an undirected graph, we add an edge in the opposite direction as well
	edge = Edge{dest, src, weight}
	g.Adjacency[dest] = append(g.Adjacency[dest], edge)
}

func (g *WeightedUndirectedGraph) PrintGraph() {
	for vertex, edges := range g.Adjacency {
		fmt.Printf("Vertex %d -> ", vertex)
		for _, edge := range edges {
			fmt.Printf("(Dest: %d, Weight: %.2f) ", edge.Dest, edge.Weight)
		}
		fmt.Println()
	}
}

func (g *WeightedUndirectedGraph) GetUniqueEdges() []Edge {
	uniqueEdges := make(map[string]Edge)
	v := g.GetVertices()
	for _, vertice := range v {
		for _, edge := range g.Adjacency[vertice] {
			// Check if the reverse key exists (e.g., if there's an edge from A to B, check if there's an edge from B to A)
			reverseKey := fmt.Sprintf("%d-%d", edge.Dest, edge.Src)
			_, reverseExists := uniqueEdges[reverseKey]

			// If the reverse edge doesn't exist, add the current edge to the uniqueEdges map
			if !reverseExists {
				// Create a unique key for each edge based on the src and dest vertices
				key := fmt.Sprintf("%d-%d", edge.Src, edge.Dest)
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
		return edges[i].Weight < edges[j].Weight
	})

	return edges
}

func (g *WeightedUndirectedGraph) GetVertices() []int {
	keys := maps.Keys(g.Adjacency)

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}
