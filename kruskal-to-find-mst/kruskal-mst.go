package kruskaltofindmst

import "top25algorithms/weightedundirectedgraph"

func KruskalMST(graph *weightedundirectedgraph.WeightedUndirectedGraph) []weightedundirectedgraph.Edge {
	// Sort all the edges in non-decreasing order of their weight.
	uniqueEdges := graph.GetUniqueEdges()
	uniqueEdges = graph.SortEdgesByWeight(uniqueEdges)

}
