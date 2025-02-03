package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Map of Rule -> bool
type Rules map[string]bool

// List of updates
type Update []int
type Input struct {
	rules   Rules
	updates []Update
}

func getInput() (Input, error) {
	lines, err := getInputLines()
	if err != nil {
		return Input{}, err
	}
	return parseInput(lines)
}

func parseInput(lines []string) (Input, error) {
	rules := Rules{}
	updates := []Update{}

	rulesSectionDone := false

	for _, line := range lines {
		if len(strings.Trim(line, " ")) == 0 {
			rulesSectionDone = true
			continue
		}

		if !rulesSectionDone {
			rules[line] = true
		} else {
			splitVals := strings.Split(line, ",")

			update := make([]int, len(splitVals))

			for i, strVal := range splitVals {
				val, err := strconv.Atoi(strVal)
				if err != nil {
					return Input{}, err
				}
				update[i] = val
			}

			updates = append(updates, update)
		}
	}

	if len(rules) == 0 {
		return Input{}, errors.New("Missing rules in input")
	}

	if !rulesSectionDone {
		return Input{}, errors.New("Missing blank line in input")
	}

	if len(updates) == 0 {
		return Input{}, errors.New("Missing updates in put")
	}

	return Input{rules, updates}, nil
}

func getInputLines() ([]string, error) {
	inputPath, err := getInputFilePath()
	if err != nil {
		return nil, err
	}

	input, err := readInput(inputPath)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func getInputFilePath() (string, error) {
	inputFlag := flag.String("i", "input.txt", `The input file.
Must be one of: input.txt|example_one.txt|example_two.txt
If a path is used, it must end in one of the above file names`)

	flag.Parse()

	r, err := regexp.Compile(".*(input|example|example_one|example_two)\\.txt$")
	if err != nil {
		return "", err
	}

	if !r.MatchString(*inputFlag) {
		return "", fmt.Errorf("Invalid input file: %q", *inputFlag)
	}

	return *inputFlag, nil
}

func readInput(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	var result []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
