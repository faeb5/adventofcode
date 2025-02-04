package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
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

	for _, u := range input.updates {
		if isUpdateCorrect(u, input.rules) {
			sum += getMedianOfUpdate(u)
		}
	}

	fmt.Println("The solution to part one is:", sum)
}

func isUpdateCorrect(u Update, r Rules) bool {
	for i, val := range u {
		left := u[:i]
		right := u[i+1:]

		valStr := strconv.Itoa(val)

		for _, leftVal := range left {
			var sb strings.Builder
			sb.WriteString(strconv.Itoa(leftVal))
			sb.WriteString("|")
			sb.WriteString(valStr)
			if _, ok := r[sb.String()]; !ok {
				return false
			}
		}

		for _, rightVal := range right {
			var sb strings.Builder
			sb.WriteString(valStr)
			sb.WriteString("|")
			sb.WriteString(strconv.Itoa(rightVal))
			if _, ok := r[sb.String()]; !ok {
				return false
			}
		}
	}

	return true
}

func getMedianOfUpdate(u Update) int {
	return u[int(math.Floor(float64(len(u))/2))]
}
