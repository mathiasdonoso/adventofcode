package day1

import (
	"strconv"
	"strings"
)

// ChronalCalibration - Problem Day 1
// https://adventofcode.com/2018/day/1
func ChronalCalibration(frequencyChanges string) int {
	changes := strings.Split(frequencyChanges, "\n")

	frecuency := 0

	for _, change := range changes {
		changeInt, _ := strconv.Atoi(change[1:])
		if strings.HasPrefix(change, "+") {
			frecuency = frecuency + changeInt
		}

		if strings.HasPrefix(change, "-") {
			frecuency = frecuency - changeInt
		}
	}

	return frecuency
}
