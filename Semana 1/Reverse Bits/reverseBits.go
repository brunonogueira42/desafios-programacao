package main

import "fmt"

// https://leetcode.com/problems/reverse-bits/description/

// Complexidade Temporal: O(1)
// Complexidade Espacial: O(1)

func reverseBits(num uint32) uint32 {
	var variavel uint32 = 0
	for i := 0; i < 32; i++ {
		variavel = (variavel << 1) | (num & 1) // Move os bits de variavel e adiciona o bit menos significativo de n
		num >>= 1                              // Desloca n para processar o próximo bit
	}
	return variavel
}

func main() {
	var num1 uint32 = 43261596
	var num2 uint32 = 4294967293

	fmt.Println(reverseBits(num1)) // Saída esperada: 964176192
	fmt.Println(reverseBits(num2)) // Saída esperada: 3221225471
}
