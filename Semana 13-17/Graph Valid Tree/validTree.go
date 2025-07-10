package main

import "fmt"

// https://neetcode.io/solutions/graph-valid-tree

// Complexidade Temporal: O(n × m(n))
// Complexidade Espacial: O(n)

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		rootX := find(x)
		rootY := find(y)
		if rootX == rootY {
			return false
		}
		parent[rootY] = rootX
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return false
		}
	}

	return true
}

func main() {
	n := 5
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}

	fmt.Println("É árvore válida?", validTree(n, edges))

	edges2 := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}
	fmt.Println("É árvore válida?", validTree(n, edges2))
}
