package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}
func reverseWords(s string) string {
	splited := strings.Fields(s)
	slices.Reverse(splited)

	result := ""
	for _, word := range splited {
		result += word + " "
	}
	return strings.TrimSpace(result)
}
