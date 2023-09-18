package kruskaltofindmst

import "top25algorithms/weightedundirectedgraph"

func KruskalMST(graph *weightedundirectedgraph.WeightedUndirectedGraph) []weightedundirectedgraph.Edge {
	mst := make([]weightedundirectedgraph.Edge, 0)

	// Sort all the edges in non-decreasing order of their weight.
	uniqueEdges := graph.GetUniqueEdges()
	uniqueEdges = graph.SortEdgesByWeight(uniqueEdges)

	return mst
}
