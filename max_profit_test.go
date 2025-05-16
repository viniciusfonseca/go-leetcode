package main

import "testing"

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	if len(prices) == 1 {
		return 0
	}

	max := 0
	min := prices[0]
	for _, v := range prices[1:] {
		if v < min {
			min = v
		} else if v-min > max {
			max = v - min
		}
	}
	return max
}

func TestMaxProfit(t *testing.T) {
	ret := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if ret != 5 {
		t.Errorf("maxProfit([]int{7, 1, 5, 3, 6, 4}) = %d; want 5", ret)
	}

	ret = maxProfit([]int{7, 6, 4, 3, 1})
	if ret != 0 {
		t.Errorf("maxProfit([]int{7, 6, 4, 3, 1}) = %d; want 0", ret)
	}
}
