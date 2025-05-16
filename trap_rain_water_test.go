package main

import "testing"

func trap(height []int) int {

	lenHeight := len(height)

	if lenHeight < 3 {
		return 0
	}
	left, right := 0, lenHeight-1 // indexes
	maxLeft, maxRight := 0, 0
	totalVolume := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				totalVolume += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				totalVolume += maxRight - height[right]
			}
			right--
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
