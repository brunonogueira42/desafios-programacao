package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	prev1 := 0 // dp[i-1]
	prev2 := 0 // dp[i-2]

	for _, num := range nums {
		current := max(prev1, num+prev2)
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))    // Output: 4
	fmt.Println(rob([]int{2, 7, 9, 3, 1})) // Output: 12
}
