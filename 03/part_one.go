package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func solvePartOne(content string) {
	content = strings.ReplaceAll(content, "\r", "")
	content = strings.ReplaceAll(content, "\n", "")

	sum := 0

	for _, match := range findAllSymbols(content) {
		x, y := findNumbers(match)
		sum += x * y
	}

	fmt.Println(sum)
}

func findNumbers(symbol string) (int, int) {
	numRegexp, err := regexp.Compile(`\d+`)
	if err != nil {
		log.Fatal(err)
	}
	nums := numRegexp.FindAllString(symbol, -1)

	x, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatal(err)
	}

	return x, y
}

func findAllSymbols(content string) []string {
	symbolRegexp, err := regexp.Compile(`mul\(\d+,\d+\)`)
	if err != nil {
		log.Fatal(err)
	}
	return symbolRegexp.FindAllString(content, -1)
}
