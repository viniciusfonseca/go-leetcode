package main

import "testing"

func trap(height []int) int {

	totalVolume := 0

	for i := range height {
		leftMax := 0
		rightMax := 0
		for j := i - 1; j >= 0; j-- {
			leftMax = max(leftMax, height[j])
		}
		for j := i + 1; j < len(height); j++ {
			rightMax = max(rightMax, height[j])
		}
		volume := min(leftMax, rightMax) - height[i]
		if volume > 0 {
			totalVolume += volume
		}
	}

	return totalVolume
}

func TestTrap(t *testing.T) {
	ret := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	if ret != 6 {
		t.Errorf("trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}) = %d; want 6", ret)
	}

	ret = trap([]int{4, 2, 0, 3, 2, 5})
	if ret != 9 {
		t.Errorf("trap([]int{4, 2, 0, 3, 2, 5}) = %d; want 9", ret)
	}
}
