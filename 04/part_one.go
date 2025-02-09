package main

import (
	"fmt"
	"strings"
)

func solvePartOne(content string) {
	sum := 0

	runes := toRunes(content)

	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes[i]); j++ {
			if runes[i][j] == rune('X') {
				// search right
				if len(runes[i])-j > 3 && strings.HasPrefix(string(runes[i][j:]), "XMAS") {
					sum++
				}

				// search left
				if j > 2 && strings.HasSuffix(string(runes[i][:j+1]), "SAMX") {
					sum++
				}

				// search down
				if len(runes)-i > 3 &&
					runes[i+1][j] == rune('M') &&
					runes[i+2][j] == rune('A') &&
					runes[i+3][j] == rune('S') {
					sum++
				}

				// search up
				if i > 2 &&
					runes[i-1][j] == rune('M') &&
					runes[i-2][j] == rune('A') &&
					runes[i-3][j] == rune('S') {
					sum++
				}

				// search diagonal up-right
				if i > 2 && len(runes[i])-j > 3 &&
					runes[i-1][j+1] == rune('M') &&
					runes[i-2][j+2] == rune('A') &&
					runes[i-3][j+3] == rune('S') {
					sum++
				}

				// search diagonal down-right
				if len(runes)-i > 3 &&
					len(runes[i])-j > 3 &&
					runes[i+1][j+1] == rune('M') &&
					runes[i+2][j+2] == rune('A') &&
					runes[i+3][j+3] == rune('S') {
					sum++
				}

				// search diagnoal down-left
				if len(runes)-i > 3 && j > 2 &&
					runes[i+1][j-1] == rune('M') &&
					runes[i+2][j-2] == rune('A') &&
					runes[i+3][j-3] == rune('S') {
					sum++
				}

				// search diagonal up-left
				if i > 2 && j > 2 &&
					runes[i-1][j-1] == rune('M') &&
					runes[i-2][j-2] == rune('A') &&
					runes[i-3][j-3] == rune('S') {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func toRunes(content string) [][]rune {
	var runes [][]rune

	for _, line := range strings.Split(content, "\n") {
		if len(line) == 0 {
			continue
		}

		var runesInLine []rune

		for _, r := range line {
			runesInLine = append(runesInLine, r)
		}

		runes = append(runes, runesInLine)
	}

	return runes
}
