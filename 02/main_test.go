package main

import (
	"math/rand/v2"
	"strings"
	"testing"
)

func testGcdOfStrings_constraints(t *testing.T) {
	randomStr := generateRandomString(1001)
	result := gcdOfStrings("asd", randomStr)
	want := ""

	if want != result {
		t.Errorf("gcdOfStrings(%s, %s) = %s; want %s", "asd", randomStr, result, want)
	}

	result = gcdOfStrings("", "asd")
	want = ""

	if want != result {
		t.Errorf("gcdOfStrings(%s, %s) = %s; want %s", "", "asd", result, want)
	}
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sb strings.Builder
	sb.Grow(length) // Optimize memory allocation upfront

	for i := 0; i < length; i++ {
		// rand.IntN is safe and concurrently secure in math/rand/v2
		randomIndex := rand.IntN(len(charset))
		sb.WriteByte(charset[randomIndex])
	}

	return sb.String()
}
