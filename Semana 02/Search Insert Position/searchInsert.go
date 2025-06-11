package main

import "fmt"

// https://leetcode.com/problems/search-insert-position/

// Complexidade Temporal: O(n)
// Complexidade Espacial: O(1)

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		} else {
			dif := target - nums[i]
			if dif < 0 {
				return i
			}
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 3, 5, 6}

	fmt.Println(searchInsert(nums, 5)) // Esperado: 2
	fmt.Println(searchInsert(nums, 2)) // Esperado: 1
	fmt.Println(searchInsert(nums, 7)) // Esperado: 4
}
