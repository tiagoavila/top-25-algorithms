package dijkstra

import (
	"top25algorithms/weighteddirectedgraph"

	queue "github.com/Jcowwell/go-algorithm-club/PriorityQueue"
)

type DijkstraNode struct {
	Vertex   string // Current vertex
	Previous string // Previous vertex where the current vertex was reached from
	Weight   int    // Value of the Total Weight from the source vertex to the current vertex
}

func NewDijkstraNode(values ...interface{}) *DijkstraNode {
	// Handle different cases based on the number and types of parameters
	if len(values) == 3 {
		vertex := values[0].(string)
		previous := values[1].(string)
		weight := values[2].(int)
		return &DijkstraNode{vertex, previous, weight}
	} else if len(values) == 2 {
		vertex := values[0].(string)
		weight := values[1].(int)
		return &DijkstraNode{vertex, "", weight}
	} else {
		// Handle error or default case
		return nil
	}
}

func lessThan(d1, d2 DijkstraNode) bool {
	return d1.Weight < d2.Weight
}

func Dijkstra(graph *weighteddirectedgraph.WeightedDirectedGraph, src, dest string) {
	priorityQueue := queue.PriorityQueueInit(lessThan)
	knownDistances := make(map[string]DijkstraNode)

	knownDistances[src] = *NewDijkstraNode(src, 0)

	// Add neighbors of the source vertex to the priority queue and knownDistances map
	for _, edge := range graph.Vertices[src] {
		updateQueueAndDistances(priorityQueue, knownDistances, src, 0, edge)
	}

	foundShortestPath := false

	for !priorityQueue.IsEmpty() {
		if foundShortestPath {
			break
		}

		currentNode, _ := priorityQueue.Dequeue()

		// Check edges of the current node
		for _, edge := range graph.Vertices[currentNode.Vertex] {
			neighborNode, neighborIsKnown := knownDistances[edge.Dest]

			if neighborIsKnown {
				// If neighbor is known and was already processed, skip it
				if priorityQueue.IndexOf(neighborNode) == -1 {
					continue
				}

				// If destination is known and new weight calculation is higher than weight to the destination, skip it
				destNode, destIsKnown := knownDistances[dest]
				if destIsKnown && destNode.Weight < currentNode.Weight+edge.Weight {
					foundShortestPath = true
					break
				}

				weightFromCurrentNodeToNeighbor := currentNode.Weight + edge.Weight
				if weightFromCurrentNodeToNeighbor < knownDistances[edge.Dest].Weight {
					updatedNeighbor := *NewDijkstraNode(edge.Dest, currentNode.Vertex, weightFromCurrentNodeToNeighbor)

					neighborIndex := priorityQueue.IndexOf(knownDistances[edge.Dest])

					knownDistances[edge.Dest] = updatedNeighbor
					priorityQueue.ChangePriority(neighborIndex, updatedNeighbor)
				}
			} else {
				updateQueueAndDistances(priorityQueue, knownDistances, currentNode.Vertex, currentNode.Weight, edge)
			}
		}
	}

	// Print the shortest path from the source vertex to the destination vertex
	node := knownDistances[dest]
	path := ""
	for node.Vertex != "" {
		if path == "" {
			path = node.Vertex
		} else {
			path = node.Vertex + " -> " + path
		}

		node = knownDistances[node.Previous]
	}
}

func updateQueueAndDistances(priorityQueue *queue.PriorityQueue[DijkstraNode], knownDistances map[string]DijkstraNode, srcVertice string, srcWeight int, edge weighteddirectedgraph.Edge) {
	node := *NewDijkstraNode(edge.Dest, srcVertice, edge.Weight+srcWeight)
	priorityQueue.Enqueue(node)
	knownDistances[edge.Dest] = node
}
