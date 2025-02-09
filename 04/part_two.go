package main

import "fmt"

func solvePartTwo(content string) {
	sum := 0

	runes := toRunes(content)

	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes[i]); j++ {
			if runes[i][j] == rune('A') &&
				isEnoughRoom(runes, i, j) &&
				isXmasShapeAt(runes, i, j) {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func isXmasShapeAt(runes [][]rune, i int, j int) bool {
	if runes[i-1][j-1] == rune('M') &&
		runes[i-1][j+1] == rune('S') &&
		runes[i+1][j-1] == rune('M') &&
		runes[i+1][j+1] == rune('S') {
		return true
	}

	if runes[i-1][j-1] == rune('M') &&
		runes[i-1][j+1] == rune('M') &&
		runes[i+1][j-1] == rune('S') &&
		runes[i+1][j+1] == rune('S') {
		return true
	}

	if runes[i-1][j-1] == rune('S') &&
		runes[i-1][j+1] == rune('M') &&
		runes[i+1][j-1] == rune('S') &&
		runes[i+1][j+1] == rune('M') {
		return true
	}

	if runes[i-1][j-1] == rune('S') &&
		runes[i-1][j+1] == rune('S') &&
		runes[i+1][j-1] == rune('M') &&
		runes[i+1][j+1] == rune('M') {
		return true
	}

	return false
}
func isEnoughRoom(runes [][]rune, i int, j int) bool {
	return i > 0 && len(runes)-i > 1 &&
		j > 0 && len(runes[i])-j > 1
}
