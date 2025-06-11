package main

import "fmt"

// https://leetcode.com/problems/subsets/

// Complexidade Temporal: O(n * 2^n)
// Complexidade Espacial: O(n * 2^n)

func subsets(nums []int) [][]int {
	var result [][]int
	generateSubsets(nums, 0, []int{}, &result)
	return result
}

func generateSubsets(nums []int, index int, current []int, result *[][]int) {
	copySubset := make([]int, len(current))
	copy(copySubset, current)
	*result = append(*result, copySubset)

	for i := index; i < len(nums); i++ {
		current = append(current, nums[i])
		generateSubsets(nums, i+1, current, result)
		current = current[:len(current)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	result := subsets(nums)

	for _, subset := range result {
		fmt.Println(subset)
	}
}
