package main

import "fmt"

// Complexidade Temporal: O(n²)
// Complexidade Espacial: O(1)

func hIndex(citations []int) int {
	for i := 0; i < len(citations)-1; i++ {
		for j := 0; j < len(citations)-i-1; j++ {
			if citations[j] < citations[j+1] {
				citations[j], citations[j+1] = citations[j+1], citations[j]
			}
		}
	}

	for i := 0; i < len(citations); i++ {
		if citations[i] < i+1 {
			return i
		}
	}

	return len(citations)
}

func main() {
	citations1 := []int{3, 0, 6, 1, 5}
	citations2 := []int{1, 3, 1}

	fmt.Println(hIndex(citations1)) // Esperado: 3
	fmt.Println(hIndex(citations2)) // Esperado: 1
}
