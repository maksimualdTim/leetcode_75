package main

import "testing"

func testMergeAlternately_BothLengthEqual(t *testing.T) {
	word1 := "abc"
	word2 := "pqr"
	want := "apbqcr"

	result := mergeAlternately(word1, word2)

	if result != want {
		t.Errorf("mergeAlternately(%s, %s) = %s; want %s", word1, word2, result, want)
	}
}

func testMergeAlternately_Word1Longer(t *testing.T) {
	word1 := "ab"
	word2 := "pqrs"
	want := "apbqrs"

	result := mergeAlternately(word1, word2)

	if result != want {
		t.Errorf("mergeAlternately(%s, %s) = %s; want %s", word1, word2, result, want)
	}
}

func testMergeAlternately_Word2Longer(t *testing.T) {
	word1 := "abcd"
	word2 := "pq"
	want := "apbqcd"

	result := mergeAlternately(word1, word2)

	if result != want {
		t.Errorf("mergeAlternately(%s, %s) = %s; want %s", word1, word2, result, want)
	}
}

func testMergeAlternately_Lowercase(t *testing.T) {
	word1 := "aBc"
	word2 := "PqR"
	want := "apbqcr"

	result := mergeAlternately(word1, word2)

	if result != want {
		t.Errorf("mergeAlternately(%s, %s) = %s; want %s", word1, word2, result, want)
	}
}

func testMergeAlternately_Letter_ValidationError(t *testing.T) {
	word1 := "1Bc"
	word2 := "PqR"
	want := ""

	result := mergeAlternately(word1, word2)

	if result != want {
		t.Errorf("mergeAlternately(%s, %s) = %s; want %s", word1, word2, result, want)
	}
}

func testMergeAlternately_Length_ValidationError(t *testing.T) {
	word1 := ""
	word2 := "PqR"
	want := ""

	result := mergeAlternately(word1, word2)

	if result != want {
		t.Errorf("mergeAlternately(%s, %s) = %s; want %s", word1, word2, result, want)
	}

	word1 = "abc"
	word2 = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" //101 symbol
	want = ""

	result = mergeAlternately(word1, word2)

	if result != want {
		t.Errorf("mergeAlternately(%s, %s) = %s; want %s", word1, word2, result, want)
	}
}
