package main

import (
	"strings"
	"testing"
)

func reverseWords(s string) string {

	var words []string
	for _, word := range strings.Split(s, " ") {
		if word != "" {
			words = append(words, word)
		}
	}

	var result strings.Builder
	for i := len(words) - 1; i >= 0; i-- {
		result.WriteString(words[i])
		result.WriteString(" ")
	}

	resultstr := result.String()
	if resultstr[len(resultstr)-1] == ' ' {
		return resultstr[:len(resultstr)-1]
	}

	return resultstr
}

func TestReverseWords(t *testing.T) {
	ret := reverseWords("the sky is blue")
	if ret != "blue is sky the" {
		t.Errorf("reverseWords(\"the sky is blue\") = %s; want \"blue is sky the\"", ret)
	}

	ret = reverseWords("  hello world  ")
	if ret != "world hello" {
		t.Errorf("reverseWords(\"  hello world  \") = %s; want \"world hello\"", ret)
	}

	ret = reverseWords("a good   example")
	if ret != "example good a" {
		t.Errorf("reverseWords(\"a good   example\") = %s; want \"example good a\"", ret)
	}
}
