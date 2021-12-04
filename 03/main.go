package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func LoadInput(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Error opening file %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	_ = file.Close()

	var input []string

	for _, val := range lines {
		input = append(input, val)
	}

	return input
}

func PartOne() {
	input := LoadInput("input_test.txt")
	fmt.Println(input)
}

func main() {
	PartOne()
}
