package main

import "fmt"

// https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/

// Complexidade Temporal: O(n * m)
// Complexidade Espacial: O(m + n)

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	length := len(words[0])
	count := len(words)
	totalLength := length * count

	wordMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		word := words[i]
		wordMap[word]++
	}

	substrings := []int{}

	for i := 0; i <= len(s)-totalLength; i++ {
		wordsFound := make(map[string]int)
		j := 0
		for j < count {
			start := i + j*length
			word := s[start : start+length]

			expectedCount, ok := wordMap[word]
			if ok {
				wordsFound[word]++
				if wordsFound[word] > expectedCount {
					break
				}
			} else {
				break
			}
			j++
		}

		if j == count {
			substrings = append(substrings, i)
		}
	}

	return substrings
}

func main() {
	// Teste 1
	s1 := "barfoofoobarthefoobarman"
	words1 := []string{"bar", "foo", "the"}
	fmt.Println(findSubstring(s1, words1)) // Esperado: [6 9 12]

	// Teste 2
	s2 := "barfoothefoobarman"
	words2 := []string{"foo", "bar"}
	fmt.Println(findSubstring(s2, words2)) // Saída esperada: [0 9]

	// Teste 3
	s3 := "wordgoodgoodgoodbestword"
	words3 := []string{"word", "good", "best", "word"}
	fmt.Println(findSubstring(s3, words3)) // Saída esperada: []
}
