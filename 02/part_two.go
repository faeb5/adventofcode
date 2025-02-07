package main

import (
	"fmt"
	"slices"
)

func solvePartTwo(content string) {
	reports := createReports(content)

	sum := 0

	for _, report := range reports {
		if isReportSafe(report) || isDampenedReportSafe(report) {
			sum++
		}
	}

	fmt.Println(sum)
}

func isDampenedReportSafe(report []int) bool {
	for i := 0; i < len(report); i++ {
		dampenedReport := slices.Delete(copyReport(report), i, i+1)
		if isReportSafe(dampenedReport) {
			return true
		}
	}
	return false
}

func copyReport(report []int) []int {
	reportCopy := make([]int, len(report))
	for i := 0; i < len(report); i++ {
		reportCopy[i] = report[i]
	}
	return reportCopy
}
