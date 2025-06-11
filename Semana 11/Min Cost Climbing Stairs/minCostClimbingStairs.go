package main

import "fmt"

// https://leetcode.com/problems/min-cost-climbing-stairs/

// Complexidade Temporal: O(n)
// Complexidade Espacial: O(1)

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}

	first := cost[0]
	second := cost[1]

	for i := 2; i < len(cost); i++ {
		current := min(first, second) + cost[i]
		first = second
		second = current
	}

	return min(first, second)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cost1 := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost1)) // Esperado: 15

	cost2 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost2)) // Esperado: 6
}
