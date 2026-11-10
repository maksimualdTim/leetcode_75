package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(gcdOfStrings("ABABABAB", "ABAB"))
}

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < 1 || len(str2) > 1000 {
		return ""
	}
	if !isOnlyLetters(str1) || !isOnlyLetters(str2) {
		return ""
	}

	str1 = strings.ToUpper(str1)
	str2 = strings.ToUpper(str2)

	str1Len := len(str1)
	str2Len := len(str2)
	result := ""

	for i := 0; i < str1Len; i++ {
		prefix := str1[:i+1]

		if str1Len%len(prefix) != 0 || str2Len%len(prefix) != 0 {
			continue
		}

		repeatedStr1 := prefix + strings.Repeat(prefix, (str1Len-len(prefix))/len(prefix))
		repeatedStr2 := prefix + strings.Repeat(prefix, (str2Len-len(prefix))/len(prefix))
		if repeatedStr1 == str1 && repeatedStr2 == str2 {
			result = prefix
		}
	}
	return result
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
