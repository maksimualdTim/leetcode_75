package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
}
func reverseWords(s string) string {
	splited := strings.Fields(s)
	slices.Reverse(splited)

	return strings.Join(splited, " ")
}
