package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		sorted := sortString(str)
		anagramMap[sorted] = append(anagramMap[sorted], str)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}
