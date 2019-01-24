package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	// Day 1 - normal
	fmt.Println("Day 1 - normal:", InverseCaptcha(str))

	// Day 1 - extra
	fmt.Println("Day 1 - extra:", HalfwayAroundCaptcha(str))
}
