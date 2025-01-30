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
	solvePartTwo(input)
}

func getInputFlag() any {
	panic("unimplemented")
}

func solvePartOne(input []string) {
	sum := 0
	runes := toRuneArray(input)
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes[i]); j++ {
			if runes[i][j] == rune('X') {
				sum += searchForXmasWord(runes, i, j)
			}
		}
	}
	fmt.Println("Solution to part one:", sum)
}

func solvePartTwo(input []string) {
	sum := 0
	runes := toRuneArray(input)
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes[i]); j++ {
			if runes[i][j] == 'A' {
				sum += searchForXmasShape(runes, i, j)
			}
		}
	}
	fmt.Println("Solution to part two:", sum)
}

func searchForXmasShape(runes [][]rune, i, j int) int {
	// M.M
	// .A.
	// S.S
	if searchTopLeft(runes, []rune("AM"), i, j) &&
		searchTopRight(runes, []rune("AM"), i, j) &&
		searchDownRight(runes, []rune("AS"), i, j) &&
		searchDownLeft(runes, []rune("AS"), i, j) {
		return 1
	}

	// M.S
	// .A.
	// M.S
	if searchTopLeft(runes, []rune("AM"), i, j) &&
		searchTopRight(runes, []rune("AS"), i, j) &&
		searchDownRight(runes, []rune("AS"), i, j) &&
		searchDownLeft(runes, []rune("AM"), i, j) {
		return 1
	}

	// S.S.
	// .A.
	// M.M.
	if searchTopLeft(runes, []rune("AS"), i, j) &&
		searchTopRight(runes, []rune("AS"), i, j) &&
		searchDownRight(runes, []rune("AM"), i, j) &&
		searchDownLeft(runes, []rune("AM"), i, j) {
		return 1
	}

	// S.M
	// .A.
	// S.M
	if searchTopLeft(runes, []rune("AS"), i, j) &&
		searchTopRight(runes, []rune("AM"), i, j) &&
		searchDownRight(runes, []rune("AM"), i, j) &&
		searchDownLeft(runes, []rune("AS"), i, j) {
		return 1
	}

	return 0
}

func searchForXmasWord(runes [][]rune, i int, j int) int {
	hits := 0

	pattern := []rune("XMAS")

	if searchTop(runes, pattern, i, j) {
		hits++
	}
	if searchTopRight(runes, pattern, i, j) {
		hits++
	}
	if searchRight(runes, pattern, i, j) {
		hits++
	}
	if searchDownRight(runes, pattern, i, j) {
		hits++
	}
	if searchDown(runes, pattern, i, j) {
		hits++
	}
	if searchDownLeft(runes, pattern, i, j) {
		hits++
	}
	if searchLeft(runes, pattern, i, j) {
		hits++
	}
	if searchTopLeft(runes, pattern, i, j) {
		hits++
	}

	return hits
}

func searchRight(runes [][]rune, pattern []rune, i int, j int) bool {
	if (len(runes[i]) - j) < len(pattern) {
		return false
	}

	// go right for each rune in the pattern
	for x, r := range pattern {
		if runes[i][j+x] != r {
			return false
		}
	}

	return true
}

func searchLeft(runes [][]rune, pattern []rune, i int, j int) bool {
	// is enough space to the left?
	if (j + 1) < len(pattern) {
		return false
	}

	// go left for each rune in the pattern
	for x, r := range pattern {
		if runes[i][j-x] != r {
			return false
		}
	}

	return true
}

func searchDown(runes [][]rune, pattern []rune, i int, j int) bool {
	// is enough space at the bottom?
	if (len(runes) - i) < len(pattern) {
		return false
	}

	// go down for each rune in the pattern
	for x, r := range pattern {
		if runes[i+x][j] != r {
			return false
		}
	}

	return true
}

func searchTop(runes [][]rune, pattern []rune, i int, j int) bool {
	// is enough space at the top?
	if (i + 1) < len(pattern) {
		return false
	}

	// go up for each rune in the pattern
	for x, r := range pattern {
		if runes[i-x][j] != r {
			return false
		}
	}

	return true
}

func searchDownRight(runes [][]rune, pattern []rune, i int, j int) bool {
	if (len(runes)-i) < len(pattern) ||
		(len(runes[i])-j) < len(pattern) {
		return false
	}

	// go one down and one right for each rune in the pattern
	for x, r := range pattern {
		if runes[i+x][j+x] != r {
			return false
		}
	}

	return true
}

func searchTopRight(runes [][]rune, pattern []rune, i int, j int) bool {
	// is there enough space at the top and right?
	if (i+1) < len(pattern) ||
		(len(runes[i])-j) < len(pattern) {
		return false
	}

	// go one up and one right for each rune in the pattern
	for x, r := range pattern {
		if runes[i-x][j+x] != r {
			return false
		}
	}

	return true
}

func searchDownLeft(runes [][]rune, pattern []rune, i int, j int) bool {
	// is there enough space at the bottom and left?
	if (len(runes)-i) < len(pattern) ||
		(j+1) < len(pattern) {
		return false
	}

	// go one down and one left for each rune in the pattern
	for x, r := range pattern {
		if runes[i+x][j-x] != r {
			return false
		}
	}

	return true
}

func searchTopLeft(runes [][]rune, pattern []rune, i int, j int) bool {
	if (i+1) < len(pattern) ||
		(j+1) < len(pattern) {
		return false
	}

	// go one up and one left for each rune in the pattern
	for x, r := range pattern {
		if runes[i-x][j-x] != r {
			return false
		}
	}

	return true
}

func toRuneArray(input []string) [][]rune {
	runes := make([][]rune, len(input))
	for i, line := range input {
		runes[i] = []rune(line)
	}
	return runes
}
