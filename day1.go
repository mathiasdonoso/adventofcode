package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	data, err := ioutil.ReadFile("day1.txt")
	handleError(err)
	str := string(data)
	fmt.Println(InverseCaptcha(str))
}

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
