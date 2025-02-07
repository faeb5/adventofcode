package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	inputFile := flag.String("i", "example.txt", "The input file.")
	part := flag.Int("p", 1, "The part to solve. Can be 1 or 2.")
	flag.Parse()

	rawContent, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	content := string(rawContent)

	switch *part {
	case 1:
		solvePartOne(content)
	case 2:
		solvePartTwo(content)
	default:
		log.Fatal("Unknown part number: ", *part)
	}
}
