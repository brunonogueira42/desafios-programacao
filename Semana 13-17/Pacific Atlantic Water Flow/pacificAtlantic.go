package main

import "fmt"

// https://leetcode.com/problems/pacific-atlantic-water-flow/

// Complexidade Temporal: O(m * n)
// Complexidade Espacial: O(m * n)

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)

	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func(r, c int, visited [][]bool, prevHeight int)
	dfs = func(r, c int, visited [][]bool, prevHeight int) {
		if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] || heights[r][c] < prevHeight {
			return
		}

		visited[r][c] = true

		dfs(r+1, c, visited, heights[r][c])
		dfs(r-1, c, visited, heights[r][c])
		dfs(r, c+1, visited, heights[r][c])
		dfs(r, c-1, visited, heights[r][c])
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, pacific, heights[i][0])
	}
	for j := 0; j < n; j++ {
		dfs(0, j, pacific, heights[0][j])
	}

	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic, heights[i][n-1])
	}
	for j := 0; j < n; j++ {
		dfs(m-1, j, atlantic, heights[m-1][j])
	}

	var result [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func main() {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}

	result := pacificAtlantic(heights)

	fmt.Println("Células que podem atingir os dois oceanos:")
	for _, cell := range result {
		fmt.Printf("[%d, %d]\n", cell[0], cell[1])
	}
}
