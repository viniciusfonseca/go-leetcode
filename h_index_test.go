package main

import (
	"sort"
	"testing"
)

func hIndex(citations []int) int {

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for i, v := range citations {
		if i >= v {
			return i
		}
	}
	return len(citations)
}

func TestHIndex(t *testing.T) {
	ret := hIndex([]int{3, 0, 6, 1, 5})
	if ret != 3 {
		t.Errorf("hIndex([]int{3, 0, 6, 1, 5}) = %d; want 3", ret)
	}

	ret = hIndex([]int{1, 3, 1})
	if ret != 1 {
		t.Errorf("hIndex([]int{1, 3, 1}) = %d; want 1", ret)
	}

	ret = hIndex([]int{1})
	if ret != 1 {
		t.Errorf("hIndex([]int{1}) = %d; want 1", ret)
	}

	ret = hIndex([]int{11, 15})
	if ret != 2 {
		t.Errorf("hIndex([]int{11, 15}) = %d; want 2", ret)
	}
}
