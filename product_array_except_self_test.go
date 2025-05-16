package main

import "testing"

func productExceptSelf(nums []int) []int {

	lenNums := len(nums)
	left := make([]int, lenNums)
	right := make([]int, lenNums)

	left[0] = 1
	right[lenNums-1] = 1
	for i := 1; i < lenNums; i++ {
		left[i] = left[i-1] * nums[i-1]
		right[lenNums-1-i] = right[lenNums-i] * nums[lenNums-i]
	}
	for i := range nums {
		nums[i] = left[i] * right[i]
	}
	return nums

}

func TestProductExceptSelf(t *testing.T) {

	type TestCase struct {
		nums         []int
		expectedNums []int
	}

	testcases := []TestCase{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, tc := range testcases {
		productExceptSelf(tc.nums)
		for i := range tc.nums {
			if tc.nums[i] != tc.expectedNums[i] {
				t.Errorf("productExceptSelf(%v) = %v; want %v", tc.nums, tc.nums, tc.expectedNums)
				return
			}
		}
	}
}
