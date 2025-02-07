package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func solvePartOne(content string) {
	reports := createReports(content)

	sum := 0

	for _, report := range reports {
		if isReportSafe(report) {
			sum++
		}
	}

	fmt.Println(sum)
}

func isReportSafe(report []int) bool {
	// direction < 0 = negative, direction > 0 = positive
	dir := 0

	for i := 0; i < len(report)-1; i++ {
		dist := report[i] - report[i+1]
		if (dist > 0 && dir < 0) || (dist < 0 && dir > 0) {
			return false
		}

		if dist < 0 {
			dir = -1
			dist = dist * -1
		} else {
			dir = 1
		}

		if dist < 1 || dist > 3 {
			return false
		}

	}

	return true
}

func createReports(content string) [][]int {
	lines := strings.Split(content, "\n")
	reports := make([][]int, len(lines)-1)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		r, err := regexp.Compile("\\d+")
		if err != nil {
			log.Fatal(err)
		}

		matches := r.FindAllString(line, -1)
		report := make([]int, len(matches))
		for j, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				log.Fatal(err)
			}
			report[j] = num
		}
		reports[i] = report
	}

	return reports
}
