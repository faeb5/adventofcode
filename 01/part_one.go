package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func solvePartOne(content string) {
	lhs, rhs := createLeftAndRight(content)

	sort.Ints(lhs)
	sort.Ints(rhs)

	sum := addDistances(lhs, rhs)
	fmt.Println(sum)
}

func addDistances(lhs []int, rhs []int) int {
	sum := 0

	for i := 0; i < len(lhs); i++ {
		dist := lhs[i] - rhs[i]
		if dist < 0 {
			dist = dist * -1
		}
		sum += dist
	}
	return sum
}

func createLeftAndRight(content string) ([]int, []int) {
	lines := strings.Split(content, "\n")
	lhs := make([]int, len(lines)-1)
	rhs := make([]int, len(lines)-1)

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		r, err := regexp.Compile("\\d+")
		if err != nil {
			log.Fatal(err)
		}

		nums := r.FindAllString(line, -1)

		num, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}
		lhs[i] = num

		num, err = strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}
		rhs[i] = num
	}
	return lhs, rhs
}
