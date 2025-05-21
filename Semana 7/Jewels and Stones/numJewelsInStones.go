package main

import "fmt"

// https://leetcode.com/problems/jewels-and-stones/

// Complexidade Temporal: O(n + m)
// Complexidade Espacial: O(n)

func numJewelsInStones(jewels string, stones string) int {
	jewelSet := make(map[rune]bool)
	for _, j := range jewels {
		jewelSet[j] = true
	}

	count := 0
	for _, s := range stones {
		if jewelSet[s] {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) // Esperado: 3
	fmt.Println(numJewelsInStones("z", "ZZ"))       // Esperado: 0
}
