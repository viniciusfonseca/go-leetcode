package main

import (
	"strings"
	"testing"
)

func intToRoman(num int) string {

	subtracts := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder
	for i, s := range subtracts {
		for num >= s {
			result.WriteString(symbols[i])
			num -= s
		}
	}

	return result.String()

}

func TestIntToRoman(t *testing.T) {
	num := 3749
	ret := intToRoman(num)
	want := "MMMDCCXLIX"
	if ret != want {
		t.Errorf("intToRoman(%d) = %s; want %s", num, ret, want)
	}

	num = 58
	ret = intToRoman(num)
	want = "LVIII"
	if ret != want {
		t.Errorf("intToRoman(%d) = %s; want %s", num, ret, want)
	}

	num = 1994
	ret = intToRoman(num)
	want = "MCMXCIV"
	if ret != want {
		t.Errorf("intToRoman(%d) = %s; want %s", num, ret, want)
	}
}
