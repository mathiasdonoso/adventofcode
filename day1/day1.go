package day1

import (
	"strconv"
)

// InverseCaptcha Problem Day 1
// The captcha requires you to review a sequence of digits (your puzzle input)
// and find the sum of all digits that match the next digit in the list.
// The list is circular, so the digit after the last digit is the first digit in the list.
// https://adventofcode.com/2017/day/1
func InverseCaptcha(numbers string) int {
	sum := 0

	for pos := range numbers {
		var next string
		var nextInt int
		actual := string(numbers[pos])
		actualInt, _ := strconv.Atoi(actual)
		if pos < len(numbers)-1 {
			next = string(numbers[pos+1])
		} else {
			next = string(numbers[0])
		}

		nextInt, _ = strconv.Atoi(next)

		if actualInt == nextInt {
			sum = sum + actualInt
		}
	}

	return sum
}
