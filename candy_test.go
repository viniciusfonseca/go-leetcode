package main

import "testing"

func candy(ratings []int) int {

	total := 0
	candies := make([]int, len(ratings))
	candies[0] = 1

	for i := 1; i < len(ratings); i++ {
		candies[i] = 1
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	for _, v := range candies {
		total += v
	}

	return total
}

func TestCandy(t *testing.T) {
	ret := candy([]int{1, 0, 2})
	if ret != 5 {
		t.Errorf("candy([]int{1, 0, 2}) = %d; want 5", ret)
	}

	ret = candy([]int{1, 2, 2})
	if ret != 4 {
		t.Errorf("candy([]int{1, 2, 2}) = %d; want 4", ret)
	}
}
