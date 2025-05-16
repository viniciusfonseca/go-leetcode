package main

import (
	"testing"
)

func jumpGame2(nums []int) int {

	k := 0
	jumps := 0
	maxJump := 0

	for i := range len(nums) - 1 {
		maxJump = max(maxJump, i+nums[i])
		if i == k {
			k = maxJump
			jumps++
		}
	}
	return jumps

}

func TestJumpGame2(t *testing.T) {

	ret := jumpGame2([]int{2, 3, 1, 1, 4})
	if ret != 2 {
		t.Errorf("jumpGame2([]int{2, 3, 1, 1, 4}) = %d; want 2", ret)
	}

	ret = jumpGame2([]int{2, 3, 0, 1, 4})
	if ret != 2 {
		t.Errorf("jumpGame2([]int{2,3,0,1,4}) = %d; want 2", ret)
	}
}
