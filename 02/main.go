package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	reports := buildReports("input.txt")
	solvePartOne(reports)
	solvePartTwo(reports)
}

func solvePartOne(reports [][]int) {
	validReports := 0
	for _, report := range reports {
		if isReportValid(report) {
			validReports++
		}
	}
	fmt.Println("Solution for part one:", validReports)
}

func solvePartTwo(reports [][]int) {
	validReports := 0
	for _, report := range reports {
		if isReportValid(report) || isDampenedReportValid(report) {
			validReports++
		}
	}
	fmt.Println("Solution for part one:", validReports)
}

func isDampenedReportValid(report []int) bool {
	for i := 0; i < len(report); i++ {
		modifiedReport := removeLevelAt(report, i)
		if isReportValid(modifiedReport) {
			return true
		}
	}
	return false
}

func removeLevelAt(report []int, index int) []int {
	newReport := make([]int, len(report))
	copy(newReport, report)
	return append(newReport[:index], newReport[index+1:]...)
}

func isReportValid(report []int) bool {
	direction := 0 // positive = increasing, negative = decreasing
	for i := 1; i < len(report); i++ {
		difference := report[i] - report[i-1]

		// check if difference d is: 0 < d < 4
		diffAbs := math.Abs(float64(difference))
		if diffAbs < 1 || diffAbs > 3 {
			return false
		}

		// check for change in direction
		if direction == 0 {
			direction = difference
		} else {
			if (direction > 0 && difference < 0) || (direction < 0 && difference > 0) {
				return false
			}
		}
	}
	return true
}

func buildReports(fileName string) [][]int {
	lines, err := readInput(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var result [][]int

	for _, line := range lines {
		var ints []int
		fields := strings.Fields(line)
		for _, strVal := range fields {
			intVal, err := strconv.Atoi(strVal)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, intVal)
		}
		result = append(result, ints)
	}

	return result
}
