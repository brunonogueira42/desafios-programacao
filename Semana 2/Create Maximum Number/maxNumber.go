package main

import "fmt"

// https://leetcode.com/problems/create-maximum-number/

// Complexidade Temporal:
// Complexidade Espacial:

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	mat1 := make([][]int, k+1)
	mat2 := make([][]int, k+1)
	mat3 := make([][]int, 0)

	for i := 0; i <= k; i++ {
		mat1[i] = maxArray(nums1, i)
		mat2[i] = maxArray(nums2, i)
	}

	for i := 0; i <= k; i++ {
		a := mat1[i]
		b := mat2[k-i]

		if len(a)+len(b) == k {
			merged := mergeMax(a, b)
			mat3 = append(mat3, merged)
		}
	}

	return maxMatrix(mat3)
}

func maxArray(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	if k > len(nums) {
		return nums
	}

	maxAux := 0
	for i := 1; i <= len(nums)-k; i++ {
		if nums[i] > nums[maxAux] {
			maxAux = i
		}
	}

	return append([]int{nums[maxAux]}, maxArray(nums[maxAux+1:], k-1)...)
}

func mergeMax(a, b []int) []int {
	max := []int{}
	for len(a) > 0 || len(b) > 0 {
		if compareMax(a, b) {
			max = append(max, a[0])
			a = a[1:]
		} else {
			max = append(max, b[0])
			b = b[1:]
		}
	}
	return max
}

func maxMatrix(mat [][]int) []int {
	if len(mat) == 0 {
		return []int{}
	}

	max := mat[0]
	for i := 1; i < len(mat); i++ {
		maxAux := mat[i]
		if compareMax(maxAux, max) {
			max = maxAux
		}
	}
	return max
}

func compareMax(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}
	return len(a) > len(b)
}

func main() {
	fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5)) // Esperado: [9 8 6 5 3]
	fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))                // Esperado: [6 7 6 0 4]
	fmt.Println(maxNumber([]int{3, 9}, []int{8, 9}, 3))                   // Esperado: [9 8 9]
	fmt.Println(maxNumber([]int{8, 1, 8, 8, 6}, []int{4}, 2))             // Esperado: [8, 8]

}
