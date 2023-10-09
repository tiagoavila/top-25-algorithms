package disjointset

import "fmt"

type DisjointSet struct {
	parent map[int]int
	rank   map[int]int
}

func NewDisjointSet() *DisjointSet {
	ds := &DisjointSet{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
	return ds
}

func (ds *DisjointSet) MakeSet(x int) {
	ds.parent[x] = x
	ds.rank[x] = 0
}

func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] == x {
		return x
	}

	return ds.Find(ds.parent[x])
}

func (ds *DisjointSet) Union(x, y int) {
	rootX := ds.Find(x)
	rootY := ds.Find(y)

	if rootX != rootY {
		if ds.rank[rootX] < ds.rank[rootY] {
			ds.parent[rootX] = rootY
		} else if ds.rank[rootX] > ds.rank[rootY] {
			ds.parent[rootY] = rootX
		} else {
			ds.parent[rootY] = rootX
			ds.rank[rootX]++
		}
	}
}

func main() {
	ds := NewDisjointSet()

	// Create sets
	ds.MakeSet(1)
	ds.MakeSet(2)
	ds.MakeSet(3)

	// Perform union operations
	ds.Union(1, 2)
	ds.Union(2, 3)

	// Check if elements are in the same set
	fmt.Println(ds.Find(1) == ds.Find(3)) // Should print true
	fmt.Println(ds.Find(2) == ds.Find(3)) // Should print true
}
