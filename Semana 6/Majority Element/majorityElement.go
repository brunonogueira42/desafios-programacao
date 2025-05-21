package main

import "fmt"

// https://leetcode.com/problems/majority-element/description/?envType=problem-list-v2&envId=hash-table
// Resolver com Hashtable

// Complexidade Temporal: O(n)
// Complexidade Espacial: O(n)

func majorityElement(nums []int) int {
	counts := make(map[int]int)
	n := len(nums)

	for _, num := range nums {
		counts[num]++
		if counts[num] > n/2 {
			return num
		}
	}
	return -1
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))             // Esperado: 3
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // OutpEsperadout: 2
}
