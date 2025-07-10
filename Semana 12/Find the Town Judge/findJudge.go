package main

import "fmt"

// https://leetcode.com/problems/find-the-town-judge/description/

// Complexidade Temporal: O(n + m)
// Complexidade Espacial: O(n)

func findJudge(n int, trust [][]int) int {
	inDegree := make([]int, n+1)
	outDegree := make([]int, n+1)

	for _, t := range trust {
		truster := t[0]
		trustee := t[1]
		outDegree[truster]++
		inDegree[trustee]++
	}

	for i := 1; i <= n; i++ {
		if inDegree[i] == n-1 && outDegree[i] == 0 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))                 // Esperado: 2
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))         // Esperado: 3
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}})) // Esperado: -1
}
