package main

import (
	"testing"
)

func canCompleteCircuit(gas []int, cost []int) int {

	totalGas := 0
	totalCost := 0
	currentGas := 0
	start := 0

	for i := range gas {
		totalGas += gas[i]
		totalCost += cost[i]

		currentGas += gas[i] - cost[i]
		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
	}

	if totalGas >= totalCost {
		return start
	}

	return -1

}

func TestCanCompleteCircuit(t *testing.T) {
	ret := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	if ret != 3 {
		t.Errorf("canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}) = %d; want 3", ret)
	}

	ret = canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})
	if ret != -1 {
		t.Errorf("canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}) = %d; want -1", ret)
	}
}
