package dijkstra

import (
	"top25algorithms/weighteddirectedgraph"

	queue "github.com/Jcowwell/go-algorithm-club/PriorityQueue"
)

type DijkstraNode struct {
	vertex   int     // Current vertex
	previous int     // Previous vertex where the current vertex was reached from
	weigth   float64 // Value of the Total Weight from the source vertex to the current vertex
}

func lessThan(d1, d2 DijkstraNode) bool {
	return d1.weigth < d2.weigth
}

func Dijkstra(graph *weighteddirectedgraph.WeightedDirectedGraph, src, dest int) {
	priorityQueue := queue.PriorityQueueInit(lessThan)
	priorityQueue.IsEmpty()
}
