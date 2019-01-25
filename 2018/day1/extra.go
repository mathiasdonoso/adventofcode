package day1

import (
	"strconv"
	"strings"
)

// TwiceCalibration - Problem Day 1
// https://adventofcode.com/2018/day/1#part2
func TwiceCalibration(frequencyChanges string) int {
	changes := strings.Split(frequencyChanges, "\n")

	store := make(map[int]int)

	frecuency, times := 0, 0

	for {
		for _, change := range changes {
			changeInt, _ := strconv.Atoi(change[1:])
			if strings.HasPrefix(change, "+") {
				frecuency = frecuency + changeInt
			}

			if strings.HasPrefix(change, "-") {
				frecuency = frecuency - changeInt
			}

			store[frecuency]++

			if times > 0 && store[frecuency] == 2 {
				return frecuency
			}
		}

		times++
	}

}
