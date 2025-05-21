package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapST := make(map[byte]byte)
	mapTS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]

		if val, ok := mapST[c1]; ok {
			if val != c2 {
				return false
			}
		} else {
			if _, exists := mapTS[c2]; exists {
				return false // c2 já está mapeado para outro caractere
			}
			mapST[c1] = c2
			mapTS[c2] = c1
		}
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))     // true
	fmt.Println(isIsomorphic("foo", "bar"))     // false
	fmt.Println(isIsomorphic("paper", "title")) // true
}
