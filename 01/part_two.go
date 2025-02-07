package main

import (
	"fmt"
)

func solvePartTwo(content string) {
	lhs, rhs := createLeftAndRight(content)

	countMap := createCountMap(rhs)

	sum := 0
	for _, val := range lhs {
		count, ok := countMap[val]
		if !ok {
			count = 0
		}
		sum += val * count
	}

	fmt.Println(sum)
}

func createCountMap(ints []int) map[int]int {
	countMap := map[int]int{}
	for _, val := range ints {
		current, ok := countMap[val]
		if !ok {
			countMap[val] = 1
		} else {
			countMap[val] = current + 1
		}
	}
	return countMap
}
