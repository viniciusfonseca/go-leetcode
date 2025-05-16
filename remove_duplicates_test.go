package main

import "testing"

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	k := 0

	for _, element := range nums[1:] {
		if element != nums[k] {
			k++
			nums[k] = element
		}
	}

	return k + 1

}

func TestRemoveDuplicates(t *testing.T) {
	ret := removeDuplicates([]int{1, 1, 2})
	if ret != 2 {
		t.Errorf("removeDuplicates([]int{1, 1, 2}) = %d; want 2", ret)
	}

	ret = removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	if ret != 5 {
		t.Errorf("removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}) = %d; want 5", ret)
	}
}
