package main

import "fmt"

// https://leetcode.com/problems/minimum-height-trees/

// Complexidade Temporal: O(n)
// Complexidade Espacial: O(n)

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make([][]int, n)
	degree := make([]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
		degree[u]++
		degree[v]++
	}

	leaves := []int{}
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			leaves = append(leaves, i)
		}
	}

	remainingNodes := n
	for remainingNodes > 2 {
		remainingNodes -= len(leaves)
		newLeaves := []int{}

		for _, leaf := range leaves {
			for _, neighbor := range graph[leaf] {
				degree[neighbor]--
				if degree[neighbor] == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}

		leaves = newLeaves
	}

	return leaves
}

func main() {
	n := 6
	edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}

	result := findMinHeightTrees(n, edges)

	fmt.Println("MHT roots:", result)
}
