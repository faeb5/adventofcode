package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func getInput() ([]string, error) {
	inputPath, err := getInputPath()
	if err != nil {
		return nil, err
	}

	input, err := readInput(inputPath)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func getInputPath() (string, error) {
	inputFlag := flag.String("i", "input.txt", `The input file.
Must be one of: input.txt|example_one.txt|example_two.txt
If a path is used, it must end in one of the above file names`)

	flag.Parse()

	r, err := regexp.Compile(".*(input|example_one|example_two)\\.txt$")
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
