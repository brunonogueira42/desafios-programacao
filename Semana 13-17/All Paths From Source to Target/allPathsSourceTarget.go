package main

import "fmt"

// https://leetcode.com/problems/all-paths-from-source-to-target/

// Complexidade Temporal: O(2ⁿ)
// Complexidade Espacial: O(n × m)

func allPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	var path []int

	var dfs func(node int)
	dfs = func(node int) {
		path = append(path, node)

		if node == len(graph)-1 {
			completePath := make([]int, len(path))
			copy(completePath, path)
			result = append(result, completePath)
		} else {
			for _, neighbor := range graph[node] {
				dfs(neighbor)
			}
		}

		path = path[:len(path)-1]
	}

	dfs(0)
	return result
}

func main() {
	graph := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}

	result := allPathsSourceTarget(graph)

	fmt.Println("Caminhos do nó 0 ao", len(graph)-1, ":")
	for _, path := range result {
		fmt.Println(path)
	}
}
