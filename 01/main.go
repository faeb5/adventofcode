package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left, right := readInput("input.txt")
	solvePartOne(left, right)
	solvePartTwo(left, right)
	solvePartTwoWithMap(left, right)
}

func solvePartOne(left, right []int) {
	distances := getDistances(left, right)
	sum := sumOfInts(distances)
	fmt.Println("Solution for part one:", sum)
}

// O(2n)
func solvePartTwoWithMap(left, right []int) {
	sum := 0

	occurences := make(map[int]int)

	for _, vright := range right {
		if _, ok := occurences[vright]; ok {
			occurences[vright] += 1
		} else {
			occurences[vright] = 1
		}
	}

	for _, vleft := range left {
		occurence, ok := occurences[vleft]
		if !ok {
			occurence = 0
		}
		sum += vleft * occurence
	}

	fmt.Println("Solution for part two (with map):", sum)
}

// O(n^2)
func solvePartTwo(left, right []int) {
	sum := 0
	for _, vleft := range left {
		occurences := 0
		for _, vright := range right {
			if vright == vleft {
				occurences += 1
			}
		}
		sum += vleft * occurences
	}
	fmt.Println("Solution for part two:", sum)
}

func sumOfInts(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func getDistances(left, right []int) []int {
	sort.Ints(left)
	sort.Ints(right)

	distances := make([]int, len(left))

	for i := 0; i < len(left); i++ {
		distances[i] = int(math.Abs(float64(left[i]) - float64(right[i])))
	}

	return distances
}

func readInput(fileName string) ([]int, []int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var listOne, listTwo []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			log.Fatalf("Corrupt line: %q", scanner.Text())
		}

		valueOne, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}

		valueTwo, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}

		listOne = append(listOne, valueOne)
		listTwo = append(listTwo, valueTwo)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return listOne, listTwo
}
