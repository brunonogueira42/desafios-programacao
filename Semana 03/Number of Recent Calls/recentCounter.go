package main

import (
	"fmt"
)

// Definindo a estrutura RecentCounter
type RecentCounter struct {
	requests []int
}

// Construtor
func Constructor() RecentCounter {
	return RecentCounter{
		requests: []int{},
	}
}

// Método ping(t)
func (rc *RecentCounter) Ping(t int) int {
	// Adiciona o novo tempo à fila
	rc.requests = append(rc.requests, t)

	// Remove todos os tempos fora da janela de 3000ms
	for len(rc.requests) > 0 && rc.requests[0] < t-3000 {
		rc.requests = rc.requests[1:] // Remove o primeiro elemento
	}

	// O tamanho da fila é a quantidade de pings na janela
	return len(rc.requests)
}

// Função principal para testar
func main() {
	rc := Constructor()
	fmt.Println(rc.Ping(1))    // 1
	fmt.Println(rc.Ping(100))  // 2
	fmt.Println(rc.Ping(3001)) // 3
	fmt.Println(rc.Ping(3002)) // 3
}
