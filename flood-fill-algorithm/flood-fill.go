package floodfillalgorithm

const (
	White  = "W"
	Black  = "B"
	Red    = "R"
	Green  = "G"
	Blue   = "B"
	Yellow = "Y"
)

type Coord struct {
	Row int
	Col int
}

func FloodFill(image [][]string, sr int, sc int, newColor string) [][]string {
	srcColor := image[sr][sc]
	replaceCollor(image, Coord{sr, sc}, newColor)

	visitedNodes := make(map[Coord]bool)
	visitedNodes[Coord{sr, sc}] = true
	queue := []Coord{{sr, sc}}

	directions := []Coord{{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		for _, direction := range directions {
			neighbor := Coord{currentNode.Row + direction.Row, currentNode.Col + direction.Col}
			if isInsideBounds(image, neighbor) && isNotVisited(visitedNodes, neighbor) {
				visitedNodes[neighbor] = true

				if image[neighbor.Row][neighbor.Col] == srcColor {
					replaceCollor(image, neighbor, newColor)
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return image
}

func replaceCollor(image [][]string, coord Coord, newColor string) {
	image[coord.Row][coord.Col] = newColor
}

func isNotVisited(visitedVertexes map[Coord]bool, neighbor Coord) bool {
	return !visitedVertexes[neighbor]
}

func isInsideBounds(image [][]string, coord Coord) bool {
	return coord.Row >= 0 && coord.Row < len(image) && coord.Col >= 0 && coord.Col < len(image[0])
}
