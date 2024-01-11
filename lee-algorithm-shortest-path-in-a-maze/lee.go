package leealgorithmshortestpathinamaze

type Coord struct {
	Row int
	Col int
}

type Node struct {
	Row      int
	Col      int
	PathSize int
}

// The Lee algorithm is one possible solution for maze routing problems based on Breadth–first search.
// https://www.techiedelight.com/lee-algorithm-shortest-path-in-a-maze/
func FindShortedsPathInMaze(maze [][]int, source Coord, destination Coord) int {
	shortestPathSize := 0
	// Create a map to track visited vertices
	visitedVertexes := make(map[Coord]bool)
	// Mark the starting vertex as visited
	visitedVertexes[source] = true
	// Create a queue to store vertices to be processed
	queue := []Node{{source.Row, source.Col, 0}}

	directions := []Coord{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	// Loop till queue is empty.
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		if currentNode.Row == destination.Row && currentNode.Col == destination.Col {
			shortestPathSize = currentNode.PathSize
			break
		}

		// Explore neighbors of the current vertex
		for _, direction := range directions {
			neighbor := Coord{currentNode.Row + direction.Row, currentNode.Col + direction.Col}
			if isInsideBounds(maze, neighbor) && isAccessibleCoord(maze, neighbor) && isNotVisited(visitedVertexes, neighbor) {
				visitedVertexes[neighbor] = true
				queue = append(queue, Node{neighbor.Row, neighbor.Col, currentNode.PathSize + 1})
			}
		}
	}

	return shortestPathSize
}

func isInsideBounds(maze [][]int, coord Coord) bool {
	return coord.Row >= 0 && coord.Row < len(maze) && coord.Col >= 0 && coord.Col < len(maze[0])
}

func isAccessibleCoord(maze [][]int, coord Coord) bool {
	return maze[coord.Row][coord.Col] == 1
}

func isNotVisited(visitedVertexes map[Coord]bool, neighbor Coord) bool {
	return !visitedVertexes[neighbor]
}
