package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mathiasdonoso/adventofcode/2018/day1"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	input, err := ioutil.ReadFile("day1/input.txt")
	handleError(err)

	data := string(input)

	fmt.Println("Day 1 - normal:", day1.ChronalCalibration(data))
	fmt.Println("Day 1 - normal:", day1.TwiceCalibration(data))
}
