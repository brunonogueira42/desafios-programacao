package main

import "fmt"

// https://leetcode.com/problems/redundant-connection/

// Complexidade Temporal: O(n × m(n))
// Complexidade Espacial: O(n)

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)

	for i := 1; i <= n; i++ {
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
		px := find(x)
		py := find(y)
		if px == py {
			return false
		}
		parent[py] = px
		return true
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if !union(u, v) {
			return edge
		}
	}

	return nil
}

func main() {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}

	result := findRedundantConnection(edges)

	fmt.Println("Aresta redundante:", result)
}
