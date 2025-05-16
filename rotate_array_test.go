package main

import "testing"

func rotate(nums []int, k int) {

	lenNums := len(nums)
	k = k % lenNums
	mem := make([]int, lenNums)
	copy(mem, nums)
	for i := range lenNums {
		nums[(i+k)%lenNums] = mem[i]
	}
}

func TestRotate(t *testing.T) {

	type TestCase struct {
		nums         []int
		k            int
		expectedNums []int
	}

	testcases := []TestCase{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{0, 1, 2}, 4, []int{2, 0, 1}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}

	for _, tc := range testcases {
		rotate(tc.nums, tc.k)
		for i := range tc.nums {
			if tc.nums[i] != tc.expectedNums[i] {
				t.Errorf("rotate(%v, %d) = %v; want %v", tc.nums, tc.k, tc.nums, tc.expectedNums)
				return
			}
		}
	}

}
