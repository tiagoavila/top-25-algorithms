package bfs

import "top25algorithms/graph"

func BFS(g *graph.Graph, startingVertex int) []int {
	// Create a map to track visited vertices
	visitedVertexes := make(map[int]bool)
	// Mark the starting vertex as visited
	visitedVertexes[startingVertex] = true
	// Create a queue to store vertices to be processed
	queue := []int{startingVertex}
	// Create a result slice to store the BFS traversal result
	result := make([]int, 0)

	// Continue the BFS traversal while the queue is not empty
	for len(queue) > 0 {
		// Dequeue the first vertex in the queue
		var vertex int = queue[0]
		queue = queue[1:]
		// Append the dequeued vertex to the result
		result = append(result, vertex)

		// Explore neighbors of the current vertex
		for _, neighbor := range g.Vertices[vertex] {
			// Check if the neighbor has been visited
			_, visited := visitedVertexes[neighbor]
			if !visited {
				// Mark the neighbor as visited
				visitedVertexes[neighbor] = true
				// Enqueue the neighbor for further exploration
				queue = append(queue, neighbor)
			}
		}
	}

	// Return the BFS traversal result
	return result
}
