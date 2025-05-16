package main

import "testing"

func jumpGame(nums []int) bool {

	maxV := 0
	for i, v := range nums {
		if i > maxV {
			return false
		}
		maxV = max(maxV, i+v)
	}
	return true
}

func TestJumpGame(t *testing.T) {
	ret := jumpGame([]int{2, 3, 1, 1, 4})
	if !ret {
		t.Errorf("jumpGame([]int{2, 3, 1, 1, 4}) = %t; want true", ret)
	}

	ret = jumpGame([]int{3, 2, 1, 0, 4})
	if ret {
		t.Errorf("jumpGame([]int{3, 2, 1, 0, 4}) = %t; want false", ret)
	}
}
