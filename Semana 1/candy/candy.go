package main

import "fmt"

// https://leetcode.com/problems/candy/description/

// Complexidade Temporal: O(n)
// Complexidade Espacial: O(n)

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)

	// Cada criança recebe pelo menos um doce
	for i := range candies {
		candies[i] = 1
	}

	// Passagem da esquerda para a direita
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Passagem da direita para a esquerda
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// Soma total de doces
	total := 0
	for _, c := range candies {
		total += c
	}
	return total
}

// Função auxiliar para obter o máximo entre dois números
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Testes
	fmt.Println(candy([]int{1, 0, 2})) // Saída: 5
	fmt.Println(candy([]int{1, 2, 2})) // Saída: 4
}
