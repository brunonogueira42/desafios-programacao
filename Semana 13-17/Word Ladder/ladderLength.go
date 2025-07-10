package main

import "fmt"

// https://leetcode.com/problems/word-ladder/

// Complexidade Temporal: O(n * m²)
// Complexidade Espacial: O(n * m)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	type Pair struct {
		word  string
		steps int
	}

	queue := []Pair{{beginWord, 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.word == endWord {
			return current.steps
		}

		for i := 0; i < len(current.word); i++ {
			for ch := 'a'; ch <= 'z'; ch++ {
				newWord := current.word[:i] + string(ch) + current.word[i+1:]
				if wordSet[newWord] {
					queue = append(queue, Pair{newWord, current.steps + 1})
					delete(wordSet, newWord)
				}
			}
		}
	}

	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Println("Comprimento da menor transformação:", ladderLength(beginWord, endWord, wordList))
}
