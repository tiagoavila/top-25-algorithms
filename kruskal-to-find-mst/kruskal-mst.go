package kruskaltofindmst

import (
	"top25algorithms/disjointset"
	"top25algorithms/weightedundirectedgraph"
)

func KruskalMST(graph *weightedundirectedgraph.WeightedUndirectedGraph) []weightedundirectedgraph.Edge {
	minimumSpanningTree := make([]weightedundirectedgraph.Edge, 0, graph.Vertices-1)

	// #1 Sort all the edges in non-decreasing order of their weight.
	uniqueEdges := graph.GetUniqueEdges()
	uniqueEdges = graph.SortEdgesByWeight(uniqueEdges)

	// #2 Pick the smallest edge. Check if it forms a cycle with the spanning tree formed so far. If the cycle is not formed, include this edge. Else, discard it.

	// #2a Create a disjoint set for all Vertices
	ds := disjointset.NewDisjointSet()
	for _, vertex := range graph.GetVertices() {
		ds.MakeSet(vertex)
	}

	for _, edge := range uniqueEdges {
		rootSrc := ds.Find(edge.Src)
		rootDest := ds.Find(edge.Dest)

		if rootSrc != rootDest {
			minimumSpanningTree = append(minimumSpanningTree, edge)
			ds.Union(edge.Src, edge.Dest)

			// Repeat step #2 until there are (V-1) edges in the spanning tree.
			if len(minimumSpanningTree) == graph.Vertices-1 {
				break
			}
		}
	}

	return minimumSpanningTree
}
