package main

import "strconv"

// HalfwayAroundCaptcha Now, instead of considering the next digit, it wants you to consider the digit halfway around the circular list.
// That is, if your list contains 10 items, only include a digit in your sum if the digit 10/2 = 5 steps forward matches it.
// Fortunately, your list has an even number of elements.
// https://adventofcode.com/2017/day/1
func HalfwayAroundCaptcha(numbers string) int {
	sum := 0
	stepsForward := len(numbers) / 2
	maxIndex := len(numbers) - 1

	for pos := range numbers {
		var next string
		var nextInt int
		actual := string(numbers[pos])
		actualInt, _ := strconv.Atoi(actual)

		if pos <= maxIndex-stepsForward {
			next = string(numbers[pos+stepsForward])
		} else {
			next = string(numbers[pos-stepsForward])
		}

		nextInt, _ = strconv.Atoi(next)

		if actualInt == nextInt {
			sum = sum + actualInt
		}
	}

	return sum
}
