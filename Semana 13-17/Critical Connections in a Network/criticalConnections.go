package main

import "fmt"

// https://leetcode.com/problems/critical-connections-in-a-network/

// Complexidade Temporal: O(n + m)
// Complexidade Espacial: O(n + m)

func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	for _, conn := range connections {
		u, v := conn[0], conn[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	disc := make([]int, n)
	low := make([]int, n)
	time := 1
	var result [][]int

	var dfs func(node, parent int)
	dfs = func(node, parent int) {
		disc[node] = time
		low[node] = time
		time++

		for _, neighbor := range graph[node] {
			if neighbor == parent {
				continue
			}
			if disc[neighbor] == 0 {
				dfs(neighbor, node)
				low[node] = min(low[node], low[neighbor])
				if low[neighbor] > disc[node] {
					result = append(result, []int{node, neighbor})
				}
			} else {

				low[node] = min(low[node], disc[neighbor])
			}
		}
	}

	dfs(0, -1)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 4
	connections := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}

	result := criticalConnections(n, connections)

	fmt.Println("Arestas críticas:", result)
}
