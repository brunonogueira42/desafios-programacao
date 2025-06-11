package main

import "fmt"

// https://leetcode.com/problems/daily-temperatures/

// Complexidade Temporal: O(n)
// Complexidade Espacial: O(n)

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))

	for i := len(temperatures) - 2; i >= 0; i-- {
		j := i + 1
		for j < len(temperatures) {
			if temperatures[j] > temperatures[i] {
				answer[i] = j - i
				break
			} else if answer[j] == 0 {
				break
			} else {
				j += answer[j]
			}
		}
	}

	return answer
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})) // Esperado: [1 1 4 2 1 1 0 0]
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))                 // Esperado: [1 1 1 0]
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))                     // Esperado: [1 1 0]
}
