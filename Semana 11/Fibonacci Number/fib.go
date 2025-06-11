package main

import "fmt"

// https://leetcode.com/problems/fibonacci-number/

// Complexidade Temporal:
// Complexidade Espacial:

func fibRecursive(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibRecursive(n-1) + fibRecursive(n-2)
	}
}

func fibDynamic(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

func main() {
	fmt.Println(fibRecursive(2)) // Esperado: 1
	fmt.Println(fibRecursive(3)) // Esperado: 2
	fmt.Println(fibRecursive(4)) // Esperado: 3
	fmt.Println(fibRecursive(5)) // Esperado: 5
	fmt.Println(fibRecursive(6)) // Esperado: 8
	fmt.Println(fibRecursive(7)) // Esperado: 13
	fmt.Println(fibRecursive(8)) // Esperado: 21

	fmt.Println(fibDynamic(2)) // Esperado: 1
	fmt.Println(fibDynamic(3)) // Esperado: 2
	fmt.Println(fibDynamic(4)) // Esperado: 3
	fmt.Println(fibDynamic(5)) // Esperado: 5
	fmt.Println(fibDynamic(6)) // Esperado: 8
	fmt.Println(fibDynamic(7)) // Esperado: 13
	fmt.Println(fibDynamic(8)) // Esperado: 21
}
