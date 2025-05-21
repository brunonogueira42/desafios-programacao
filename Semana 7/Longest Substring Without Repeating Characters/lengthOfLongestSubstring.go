package main

import "fmt"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Complexidade Temporal: O(n)
// Complexidade Espacial: O(m)

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		char := s[i]

		lastIndex, found := charIndex[char]
		if found && lastIndex >= start {
			start = lastIndex + 1
		}

		charIndex[char] = i

		length := i - start + 1
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Esperado: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Esperado: 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Esperado: 3
}
