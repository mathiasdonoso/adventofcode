package day2

import (
	"strconv"
	"strings"
)

// DivisibleChecksum - Problem Day 2
// It sounds like the goal is to find the only two numbers in each row where one evenly divides the other - that is,
// where the result of the division operation is a whole number. They would like you to find those numbers on each line,
// divide them, and add up each line's result.
func DivisibleChecksum(values string) int {
	data := strings.Split(values, "\n")
	arr := make([][]string, len(data))
	sum := 0

	for pos, row := range data {
		arr[pos] = strings.Split(row, "\t")
	}

	for _, r := range arr {
		var a int
		var b int
		for index := range r {
			a, _ = strconv.Atoi(string(r[index]))
			for j := range r {
				b, _ = strconv.Atoi(string(r[j]))

				if a != b {
					if a%b == 0 {
						sum = sum + (a / b)
					}
				}
			}

		}
	}

	return sum
}
