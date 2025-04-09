package main

import "fmt"

// https://leetcode.com/problems/jump-game/description/

// Complexidade Temporal: O(n)
// Complexidade Espacial: 0(1)

func canJump(nums []int) bool {
	maxJump := 0

	for i := 0; i < len(nums); i++ {
		if i > maxJump {
			return false
		}

		if i+nums[i] > maxJump {
			maxJump = i + nums[i]
		}

		if maxJump >= len(nums)-1 {
			return true
		}
	}

	return false
}

func main() {
	nums1 := []int{2, 3, 1, 1, 4}
	nums2 := []int{3, 2, 1, 0, 4}

	fmt.Println(canJump(nums1)) // Esperado: true
	fmt.Println(canJump(nums2)) // Esperado: false
}
