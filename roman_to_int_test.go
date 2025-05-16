package main

import "testing"

func romanToInt(s string) int {

	answer := 0
	lenS := len(s)
	for i := lenS - 1; i >= 0; i-- {
		switch s[i] {
		case 'I':
			answer += 1
		case 'V':
			if i > 0 && s[i-1] == 'I' {
				answer += 3
			} else {
				answer += 5
			}
		case 'X':
			if i > 0 && s[i-1] == 'I' {
				answer += 8
			} else {
				answer += 10
			}
		case 'L':
			if i > 0 && s[i-1] == 'X' {
				answer += 30
			} else {
				answer += 50
			}
		case 'C':
			if i > 0 && s[i-1] == 'X' {
				answer += 80
			} else {
				answer += 100
			}
		case 'D':
			if i > 0 && s[i-1] == 'C' {
				answer += 300
			} else {
				answer += 500
			}
		case 'M':
			if i > 0 && s[i-1] == 'C' {
				answer += 800
			} else {
				answer += 1000
			}
		}
	}

	return answer
}

func TestRomanToInt(t *testing.T) {
	ret := romanToInt("III")
	if ret != 3 {
		t.Errorf("romanToInt(\"III\") = %d; want 3", ret)
	}

	ret = romanToInt("IV")
	if ret != 4 {
		t.Errorf("romanToInt(\"IV\") = %d; want 4", ret)
	}

	ret = romanToInt("IX")
	if ret != 9 {
		t.Errorf("romanToInt(\"IX\") = %d; want 9", ret)
	}

	ret = romanToInt("LVIII")
	if ret != 58 {
		t.Errorf("romanToInt(\"LVIII\") = %d; want 58", ret)
	}
}
