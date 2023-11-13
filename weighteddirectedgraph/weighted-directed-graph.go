package weighteddirectedgraph

type Edge struct {
	Dest   int     // Destination vertex
	Weight float64 // Weight of the edge
}

type WeightedDirectedGraph struct {
	Vertices map[int][]Edge // Adjacency list to store edges
}

func NewGraph() *WeightedDirectedGraph {
	return &WeightedDirectedGraph{
		Vertices: make(map[int][]Edge),
	}
}

func AddEdge(g *WeightedDirectedGraph, src, dest int, weight float64) {
	// Add a directed edge between src and dest with the given weight
	edge := Edge{dest, weight}
	g.Vertices[src] = append(g.Vertices[src], edge)
}
