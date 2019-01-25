package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mathiasdonoso/adventofcode/2017/day1"
	"github.com/mathiasdonoso/adventofcode/2017/day2"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var data []byte
	var str string
	var err error

	data, err = ioutil.ReadFile("day1/data.txt")
	handleError(err)
	str = string(data)
	fmt.Println("Day 1 - normal:", day1.InverseCaptcha(str))
	fmt.Println("Day 1 - extra:", day1.HalfwayAroundCaptcha(str))

	data, err = ioutil.ReadFile("day2/data.txt")
	handleError(err)
	str = string(data)
	fmt.Println("Day 2 - normal:", day2.CorruptionChecksum(str))
	fmt.Println("Day 2 - normal:", day2.DivisibleChecksum(str))
}
