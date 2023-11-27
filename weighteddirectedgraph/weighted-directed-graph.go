package weighteddirectedgraph

type Edge struct {
	Dest   string // Destination vertex
	Weight int    // Weight of the edge
}

type WeightedDirectedGraph struct {
	Vertices map[string][]Edge // Adjacency list to store edges
}

func NewGraph() *WeightedDirectedGraph {
	return &WeightedDirectedGraph{
		Vertices: make(map[string][]Edge),
	}
}

func (g *WeightedDirectedGraph) AddEdge(src, dest string, weight int) {
	// Add a directed edge between src and dest with the given weight
	edge := Edge{dest, weight}
	g.Vertices[src] = append(g.Vertices[src], edge)
}
