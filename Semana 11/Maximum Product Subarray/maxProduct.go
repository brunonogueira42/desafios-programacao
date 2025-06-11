package main

import "fmt"

// https://leetcode.com/problems/maximum-product-subarray/

// Complexidade Temporal: O(n)
// Complexidade Espacial: O(1)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProd := nums[0]
	minProd := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		current := nums[i]

		if current < 0 {
			maxProd, minProd = minProd, maxProd
		}

		maxProd = max(current, maxProd*current)
		minProd = min(current, minProd*current)

		result = max(result, maxProd)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums1)) // Esperado: 6

	nums2 := []int{-2, 0, -1}
	fmt.Println(maxProduct(nums2)) // Esperado: 0
}
