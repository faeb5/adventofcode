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

type RuleLines []string
type UpdateLines []string
type Rule [2]int
type Update []int
type Input struct {
	rules   []Rule
	updates []Update
}

func parseInput() (Input, error) {
	lines, err := getInputLines()
	if err != nil {
		return Input{}, err
	}

	ruleLines, updateLines, err := splitInputLines(lines)
	if err != nil {
		return Input{}, err
	}

	rules, err := parseRules(ruleLines)
	if err != nil {
		return Input{}, err
	}

	updates, err := parseUpdates(updateLines)
	if err != nil {
		return Input{}, err
	}

	return Input{rules, updates}, nil
}

func splitInputLines(lines []string) (RuleLines, UpdateLines, error) {
	splitLine := -1

	// find split line index
	for i, line := range lines {
		if len(strings.Trim(line, " ")) == 0 {
			splitLine = i
			break
		}
	}

	if splitLine == -1 {
		return []string{}, []string{}, errors.New("No blank line found in input")
	}

	return lines[:splitLine], lines[splitLine+1:], nil
}

func parseRules(lines RuleLines) ([]Rule, error) {
	rules := make([]Rule, len(lines))

	for i, line := range lines {
		ruleParts := strings.Split(line, "|")
		if len(ruleParts) != 2 {
			return []Rule{}, fmt.Errorf("Unable to parse rule: %q", line)
		}

		before, err := strconv.Atoi(ruleParts[0])
		if err != nil {
			return []Rule{}, err
		}

		after, err := strconv.Atoi(ruleParts[1])
		if err != nil {
			return []Rule{}, err
		}

		rules[i] = Rule{before, after}
	}

	return rules, nil
}

func parseUpdates(lines UpdateLines) ([]Update, error) {
	updates := make([]Update, len(lines))

	for i, line := range lines {
		updateParts := strings.Split(line, ",")
		if len(updateParts) == 0 {
			return []Update{}, fmt.Errorf("Unable to parse updates: %q", line)
		}

		update := make(Update, len(updateParts))
		for j, val := range updateParts {
			num, err := strconv.Atoi(val)
			if err != nil {
				return []Update{}, err
			}
			update[j] = num
		}

		updates[i] = update
	}

	return updates, nil
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
