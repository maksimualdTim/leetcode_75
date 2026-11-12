package main

import "testing"

func TestReverseWords_Success(t *testing.T) {
	input := "the sky is blue"
	want := "blue is sky the"

	result := reverseWords(input)

	if result != want {
		t.Errorf("simple case. input: %s, want: %s, result: %s", input, want, result)
	}

	input = "  hello world  "
	want = "world hello"

	result = reverseWords(input)

	if result != want {
		t.Errorf("begin-end spaces case. input: %s, want: %s, result: %s", input, want, result)
	}

	input = "a good   example"
	want = "example good a"

	result = reverseWords(input)

	if result != want {
		t.Errorf("middle spaces case. input: %s, want: %s, result: %s", input, want, result)
	}
}
