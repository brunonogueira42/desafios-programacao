// Bruno Nogueira
// Leonardo Dorneles

package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/3sum/

// Complexidade Temporal: O(n³ * m)
// Complexidade Espacial: O(n³)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var triplets [][]int

	for i := 0; i < len(nums)-2; i++ {
		for j := i; j < len(nums)-1; j++ {
			for k := j; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 && i != j && i != k && j != k {
					t := []int{nums[i], nums[j], nums[k]}
					new := true

					for l := 0; l < len(triplets); l++ {
						if comparate(t, triplets[l]) {
							new = false
							break
						}
					}

					if new {
						triplets = append(triplets, t)
					}
				}
			}
		}
	}

	return triplets
}

func comparate(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // Esperado: [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // Esperado: []
	fmt.Println(threeSum([]int{0, 0, 0}))             // Esperado: [[0,0,0]]
}
