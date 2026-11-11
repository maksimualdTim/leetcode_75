package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
}

func reverseVowels(s string) string {
	start := 0
	end := len(s) - 1
	arr := []byte(s)

	for start < end {
		if isVowel(arr[start]) && isVowel(arr[end]) {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
		if !isVowel(arr[start]) {
			start++
		}
		if !isVowel(arr[end]) {
			end--
		}
	}
	return string(arr)
}

func isVowel(r byte) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
