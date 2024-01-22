package kahnstopologicalsort

import "top25algorithms/unweighteddirectedgraph"

func TopologicalSort(graph *unweighteddirectedgraph.UnweightedDirectedGraph) []string {
	sorted := make([]string, 0, len(graph.Vertices))
	inDegreeZeroVertices := graph.GetInDegreeZeroVertices()

	for len(inDegreeZeroVertices) > 0 {
		currentVertex := inDegreeZeroVertices[0]
		sorted = append(sorted, currentVertex)
		inDegreeZeroVertices = inDegreeZeroVertices[1:]

		for _, edge := range graph.Vertices[currentVertex].Edges {
			graph.DecreaseInDegree(edge.Dest)

			if graph.Vertices[edge.Dest].InDegree == 0 {
				inDegreeZeroVertices = append(inDegreeZeroVertices, edge.Dest)
			}
		}
	}

	return sorted
}
