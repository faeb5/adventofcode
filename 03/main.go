package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	program := getProgram("input.txt")
	solvePartOne(program)
	solvePartTwo(program)
}

func solvePartTwo(program []string) {
	sum := 0
	enabled := true
	for _, line := range program {
		sanitizedLine, enabledNew := sanitizeLine(line, enabled)
		sum += getSum(sanitizedLine)
		enabled = enabledNew
	}
	fmt.Println("Solution to part two:", sum)
}

func sanitizeLine(line string, enabled bool) (string, bool) {
	var sb strings.Builder

	// valid symbols: `do()`, `don't()` and `mul(x,y)`
	grammar := newRegexp("don't\\(\\)|do\\(\\)|mul\\([0-9]+,[0-9]+\\)")
	symbols := grammar.FindAllString(line, -1)

	for _, symbol := range symbols {
		if symbol == "don't()" {
			enabled = false
		} else if symbol == "do()" {
			enabled = true
		} else if strings.HasPrefix(symbol, "mul") {
			if enabled {
				sb.WriteString(symbol)
			}
		} else {
			log.Fatal("Unknown symbol:", symbol)
		}
	}

	return sb.String(), enabled
}

func solvePartOne(program []string) {
	sum := 0

	for _, line := range program {
		sum += getSum(line)
	}

	fmt.Println("Solution to part one:", sum)
}

func getSum(line string) int {
	sum := 0

	mulRegexp := newRegexp("mul\\([0-9]+,[0-9]+\\)")
	operandRegexp := newRegexp("[0-9]+")

	for _, mul := range mulRegexp.FindAllString(line, -1) {
		product := 1
		for _, operand := range operandRegexp.FindAllString(mul, -1) {
			operandInt, err := strconv.Atoi(operand)
			if err != nil {
				log.Fatal(err)
			}
			product *= operandInt
		}

		sum += product
	}
	return sum
}

func getProgram(fileName string) []string {
	lines, err := readInput(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return lines
}

func newRegexp(pattern string) *regexp.Regexp {
	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
