package main

import (
	"fmt"
	"strings"
)

func solvePartTwo(content string) {
	content = strings.ReplaceAll(content, "\r", "")
	content = strings.ReplaceAll(content, "\n", "")

	content = strings.ReplaceAll(content, "do()", "\n")
	content = strings.ReplaceAll(content, "don't()", "\nSKIP")

	sum := 0

	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "SKIP") {
			continue
		}
		for _, symbol := range findAllSymbols(line) {
			x, y := findNumbers(symbol)
			sum += x * y
		}
	}

	fmt.Println(sum)
}
