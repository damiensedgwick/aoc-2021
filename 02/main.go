package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func MakeMap(s string) map[string]int {
	slice := strings.Fields(s)
	key := slice[0]
	value, _ := strconv.Atoi(slice[1])
	m := map[string]int{key: value}

	return m
}

func LoadInput(fileName string) []map[string]int {
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

	file.Close()

	var input []map[string]int

	for _, val := range lines {
		m := MakeMap(val)
		input = append(input, m)
	}

	return input
}

func CalculateInstructions(input []map[string]int) map[string]int {
	m := map[string]int{
		"horizontal": 0,
		"depth":      0,
	}

	for _, o := range input {
		for key, value := range o {
			switch key {
			case "forward":
				m["horizontal"] += value
			case "down":
				m["depth"] += value
			case "up":
				m["depth"] -= value
			}
		}
	}

	return m
}

func CalculateInstructionsWithAim(input []map[string]int) map[string]int {
	m := map[string]int{
		"horizontal": 0,
		"depth":      0,
	}

	aim := 0

	for _, o := range input {
		for key, value := range o {
			switch key {
			case "down":
				aim = aim + value
			case "up":
				aim = aim - value
			case "forward":
				m["horizontal"] += value

				if aim != 0 {
					m["depth"] = m["depth"] + aim*value
				}
			}
		}
	}

	return m
}

func MultiplyPosition(m map[string]int) int {
	return m["horizontal"] * m["depth"]
}

func PartOne() int {
	input := LoadInput("input.txt")
	sum := CalculateInstructions(input)
	result := MultiplyPosition(sum)

	fmt.Printf("Part One: %d \n", result)

	return result
}

func PartTwo() int {
	input := LoadInput("input.txt")
	sum := CalculateInstructionsWithAim(input)
	result := MultiplyPosition(sum)

	fmt.Printf("Part Two: %d \n", result)

	return result
}

func main() {
	PartOne()
	PartTwo()
}
