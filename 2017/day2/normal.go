package day2

import (
	"strconv"
	"strings"
)

// CorruptionChecksum - Problem Day 2
// The spreadsheet consists of rows of apparently-random numbers.
// To make sure the recovery process is on the right track, they need you to calculate the spreadsheet's checksum.
// For each row, determine the difference between the largest value and the smallest value; the checksum is the sum of all of these differences.
// https://adventofcode.com/2017/day/2
func CorruptionChecksum(values string) int {
	data := strings.Split(values, "\n")
	arr := make([][]string, len(data))
	sum := 0

	for pos, row := range data {
		arr[pos] = strings.Split(row, "\t")
	}

	for _, r := range arr {
		min, _ := strconv.Atoi(r[0])
		max, _ := strconv.Atoi(r[0])
		for _, value := range r {
			intVal, _ := strconv.Atoi(value)

			if intVal < min {
				min = intVal
			}

			if intVal > max {
				max = intVal
			}
		}

		sum = sum + (max - min)
	}

	return sum
}
