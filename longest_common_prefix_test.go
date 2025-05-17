package main

import "testing"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := range len(strs[0]) {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func TestLongestCommonPrefix(t *testing.T) {
	ret := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if ret != "fl" {
		t.Errorf("longestCommonPrefix([]string{\"flower\", \"flow\", \"flight\"}) = %s; want \"fl\"", ret)
	}

	ret = longestCommonPrefix([]string{"dog", "racecar", "car"})
	if ret != "" {
		t.Errorf("longestCommonPrefix([]string{\"dog\", \"racecar\", \"car\"}) = %s; want \"\"", ret)
	}
}
