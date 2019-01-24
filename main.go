package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mathiasdonoso/adventofcode/day1"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	data, err := ioutil.ReadFile("day1/data.txt")
	handleError(err)
	str := string(data)

	fmt.Println("Day 1 - normal:", day1.InverseCaptcha(str))
	fmt.Println("Day 1 - extra:", day1.HalfwayAroundCaptcha(str))
}
