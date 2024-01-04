package bellmanford

import (
	"fmt"
	"math"
	"top25algorithms/weighteddirectedgraph"
)

type Edge struct {
	Src    string
	Dest   string
	Weight int
}

func BellmanFord(graph *weighteddirectedgraph.WeightedDirectedGraph, startNode string, verticesCount int) map[string]int {
	distances := make(map[string]int)
	parent := make(map[string]string)
	edges := make([]Edge, 0)

	// Step 1 – initialize the graph. In the beginning, all vertices weight of
	// INFINITY and a null parent, except for the source, where the weight is 0
	for k, vertice_edges := range graph.Vertices {
		distances[k] = math.MaxInt32
		parent[k] = ""

		//Make list of all edges in graph
		for _, v := range vertice_edges {
			edges = append(edges, Edge{k, v.Dest, v.Weight})

			_, present := distances[v.Dest]
			if !present {
				distances[v.Dest] = math.MaxInt32
				parent[v.Dest] = ""
			}
		}

	}
	distances[startNode] = 0

	// Step 2 – relax edges repeatedly
	// relaxation step (run V-1 times)
	for i := 0; i < verticesCount-1; i++ {
		for _, edge := range edges {
			// edge from `u` to `v` having weight `w`
			srcVertice := edge.Src
			destVertice := edge.Dest
			weight := edge.Weight

			// if the distance to destination `v` (destVertice) can be shortened by taking edge (srcVertice, destVertice)
			if distances[srcVertice] != math.MaxInt32 && distances[srcVertice]+weight < distances[destVertice] {
				distances[destVertice] = distances[srcVertice] + weight
				parent[destVertice] = srcVertice
			}
		}
	}

	// Step 3 – check for negative-weight cycles
	// run relaxation step once more for n'th time to check for negative-weight cycles
	for _, edge := range edges {
		// edge from `u` to `v` having weight `w`
		srcVertice := edge.Src
		destVertice := edge.Dest
		weight := edge.Weight

		// if the distance to destination `v` (destVertice) can be shortened by taking edge (srcVertice, destVertice)
		if distances[srcVertice] != math.MaxInt32 && distances[srcVertice]+weight < distances[destVertice] {
			fmt.Println("Graph contains negative-weight cycle")
			return nil
		}
	}

	return distances
}
