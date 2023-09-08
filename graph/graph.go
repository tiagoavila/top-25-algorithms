package graph

import "fmt"

type Graph struct {
	Vertices map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.Vertices[v1] = append(g.Vertices[v1], v2)
	g.Vertices[v2] = append(g.Vertices[v2], v1)
}

func (g *Graph) PrintGraph() {
	for vertex, neighbors := range g.Vertices {
		fmt.Printf("Vertex %d: ", vertex)
		for _, neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}
