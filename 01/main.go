package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func FormatInput() []int {
	var s []int

	file, openErr := os.Open("input.txt")

	if openErr != nil {
		log.Fatalf("failed opening file: %s", openErr)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	closeErr := file.Close()

	if closeErr != nil {
		log.Fatalf("failed closing file: %s", closeErr)
	}

	for _, line := range txtlines {
		if r, err := strconv.Atoi(line); err == nil {
			s = append(s, r)
		}
	}

	return s
}

func Part1(input []int) int {
	var count, prev, current = 0, 0, 1

	for _ = range input {
		if input[prev] < input[current] {
			count++
		}

		if input[current] != input[len(input)-1] {
			prev++
			current++
		}
	}

	fmt.Println(count - 1)

	return count
}

func SumWindow(input []int) []int {
	var windows []int

	for i := 0; i < len(input)-2; i++ {
		sum := input[i] + input[i+1] + input[i+2]
		windows = append(windows, sum)
	}

	return windows
}

func Part2(input []int) int {
	data := SumWindow(input)

	return Part1(data)
}

func main() {
	input := FormatInput()
	Part1(input)
	Part2(input)
}
