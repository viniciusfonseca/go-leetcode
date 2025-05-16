package main

import (
	"testing"
)

func majorityElement(nums []int) int {

	mid := len(nums) / 2
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] > mid {
			return num
		}
	}
	return -1
}

func TestMajorityElement(t *testing.T) {
	ret := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	if ret != 2 {
		t.Errorf("majorityElement([]int{2, 2, 1, 1, 1, 2, 2}) = %d; want 2", ret)
	}

	ret = majorityElement([]int{3, 2, 3})
	if ret != 3 {
		t.Errorf("majorityElement([]int{3, 2, 3}) = %d; want 3", ret)
	}
}
