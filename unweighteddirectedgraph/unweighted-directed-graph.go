package unweighteddirectedgraph

type Edge struct {
	Dest string // Destination vertex
}

type VerticeData struct {
	Edges    []Edge
	InDegree int
}

type UnweightedDirectedGraph struct {
	Vertices map[string]VerticeData // Adjacency list to store edges
}

func NewGraph() *UnweightedDirectedGraph {
	return &UnweightedDirectedGraph{
		Vertices: make(map[string]VerticeData),
	}
}

func (g *UnweightedDirectedGraph) AddEdge(src, dest string) {
	// Add a directed edge between src
	edge := Edge{dest}
	verticeData := g.Vertices[src]
	verticeData.Edges = append(verticeData.Edges, edge)
	g.Vertices[src] = verticeData
	g.increaseInDegree(dest)
}

func (g *UnweightedDirectedGraph) increaseInDegree(dest string) {
	verticeData := g.Vertices[dest]
	verticeData.InDegree++
	g.Vertices[dest] = verticeData
}

func (g *UnweightedDirectedGraph) DecreaseInDegree(dest string) {
	verticeData := g.Vertices[dest]
	verticeData.InDegree--
	g.Vertices[dest] = verticeData
}

func (g *UnweightedDirectedGraph) GetInDegreeZeroVertices() []string {
	inDegreeZeroVertices := make([]string, 0, len(g.Vertices))

	for vertex, verticeData := range g.Vertices {
		if verticeData.InDegree == 0 {
			inDegreeZeroVertices = append(inDegreeZeroVertices, vertex)
		}
	}

	return inDegreeZeroVertices
}
