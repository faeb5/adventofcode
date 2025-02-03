package main

import (
	"fmt"
	"log"
)

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	solvePartOne(input)
}

func solvePartOne(input Input) {
	sum := 0

	// TODO solve the problem
	fmt.Println("The solution to part one is:", sum)
}
