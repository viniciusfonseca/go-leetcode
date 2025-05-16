package main

import "testing"

func removeDuplicates2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	k := 0
	state := 0

	for _, element := range nums[1:] {
		switch state {
		case 0:
			if element == nums[k] {
				state = 1
			}
			k++
			nums[k] = element
		case 1:
			if element != nums[k] {
				k++
				nums[k] = element
				state = 0
			}
		}
	}

	return k + 1
}

func TestRemoveDuplicates2(t *testing.T) {
	ret := removeDuplicates2([]int{1, 1, 1, 2, 2, 3})
	if ret != 5 {
		t.Errorf("removeDuplicates2([]int{1, 1, 1, 2, 2, 3}) = %d; want 5", ret)
	}

	ret = removeDuplicates2([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
	if ret != 7 {
		t.Errorf("removeDuplicates2([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}) = %d; want 7", ret)
	}
}
