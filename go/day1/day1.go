package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	count := 0
	prev := 0
	current := 1

	for _, line := range txtlines {
		if _, err := strconv.Atoi(line); err == nil {
			if txtlines[prev] < txtlines[current] {
				count++
			}
		}

		if txtlines[current] != txtlines[len(txtlines)-1] {
			prev++
			current++
		}
	}

	fmt.Println(count)
}
