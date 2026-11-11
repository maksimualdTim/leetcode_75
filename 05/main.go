package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	vowels := []byte{}
	result := []byte{}

	for i := range s {
		if isVowel(s[i]) {
			vowels = append(vowels, s[i])
		}
	}

	currentVowelIndex := len(vowels) - 1
	for i := range s {
		if isVowel(s[i]) {
			result = append(result, vowels[currentVowelIndex])
			currentVowelIndex--
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}

func isVowel(r byte) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
