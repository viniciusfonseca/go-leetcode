package main

import "testing"

func lengthOfLastWord(s string) int {

	if len(s) == 0 {
		return 0
	}

	state := 0
	lenW := 0
	if s[len(s)-1] != ' ' {
		state = 1
		lenW = 1
	}

	for i := len(s) - 2; i >= 0; i-- {
		switch state {
		case 0:
			if s[i] != ' ' {
				state = 1
				lenW = 1
			}
		case 1:
			if s[i] == ' ' {
				return lenW
			} else {
				lenW++
			}
		}
	}

	return lenW
}

func TestLengthOfLastWord(t *testing.T) {
	ret := lengthOfLastWord("Hello World")
	if ret != 5 {
		t.Errorf("lengthOfLastWord(\"Hello World\") = %d; want 5", ret)
	}

	ret = lengthOfLastWord("   fly me   to   the moon  ")
	if ret != 4 {
		t.Errorf("lengthOfLastWord(\"   fly me   to   the moon  \") = %d; want 4", ret)
	}
}
