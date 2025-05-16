package main

import (
	"testing"
)

func maxProfit2(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	if len(prices) == 1 {
		return 0
	}

	maxV := 0
	minV := prices[0]
	state := 0
	total := 0

	if prices[1] < prices[0] {
		state = 0
	} else {
		state = 1
	}
	for _, value := range prices[1:] {
		switch state {
		case 0:
			if value <= minV {
				minV = value
			} else {
				state = 1
				maxV = value
			}
		case 1:
			if value >= maxV {
				maxV = value
			} else {
				state = 0
				total += maxV - minV
				minV = value
				maxV = 0
			}
		}
	}

	return total + max(0, maxV-minV)

}

func TestMaxProfit2(t *testing.T) {
	ret := maxProfit2([]int{7, 1, 5, 3, 6, 4})
	if ret != 7 {
		t.Errorf("maxProfit2([]int{7, 1, 5, 3, 6, 4}) = %d; want 7", ret)
	}

	ret = maxProfit2([]int{1, 2, 3, 4, 5})
	if ret != 4 {
		t.Errorf("maxProfit2([]int{7, 6, 4, 3, 1}) = %d; want 4", ret)
	}

	ret = maxProfit2([]int{7, 6, 4, 3, 1})
	if ret != 0 {
		t.Errorf("maxProfit2([]int{7, 6, 4, 3, 1}) = %d; want 0", ret)
	}

	ret = maxProfit2([]int{2, 4, 1})
	if ret != 2 {
		t.Errorf("maxProfit2([]int{2, 4, 1}) = %d; want 2", ret)
	}

	ret = maxProfit2([]int{1, 7, 4, 2})
	if ret != 6 {
		t.Errorf("maxProfit2([]int{1,7,4,2}) = %d; want 6", ret)
	}
}
