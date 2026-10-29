package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(mergeAlternately("asd", "gdf"))
}

func mergeAlternately(word1 string, word2 string) string {

	if len(word2) > 100 || len(word1) < 1 {
		return ""
	}

	if !isOnlyLetters(word1) || !isOnlyLetters(word2) {
		return ""
	}

	result := []byte{}
	word1Length := len(word1)
	word2Length := len(word2)

	for i := range max(word1Length, word2Length) {
		if i < word1Length {
			result = append(result, word1[i])
		}
		if i < word2Length {
			result = append(result, word2[i])
		}
	}
	return string(result)
}

func isOnlyLetters(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
