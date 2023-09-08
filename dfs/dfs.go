package dfs

import "top25algorithms/graph"

func DFS(g *graph.Graph, currentVertex int) []int {
	// Create a map to track visited vertices
	visitedVertexes := make(map[int]bool)

	return dfs(g, currentVertex, visitedVertexes)
}

func dfs(g *graph.Graph, currentVertex int, visitedVertexes map[int]bool) []int {
	// Mark the starting vertex as visited
	visitedVertexes[currentVertex] = true

	// Create a result slice to store the BFS traversal result
	result := make([]int, 0)
	result = append(result, currentVertex)

	for _, neighbor := range g.Vertices[currentVertex] {
		// Check if the neighbor has been visited
		_, visited := visitedVertexes[neighbor]
		if !visited {
			// Mark the neighbor as visited
			visitedVertexes[neighbor] = true
			// Enqueue the neighbor for further exploration
			result = append(result, dfs(g, neighbor, visitedVertexes)...)
		}
	}

	return result
}
